package simulator

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

type Simulator struct {
	fn        func()
	completed uint64
}

type RateLimiter func(completed uint64) bool

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func NewSimulator(simulation func()) *Simulator {
	return &Simulator{
		fn: simulation,
	}
}

func (s *Simulator) runner(rateLimiter RateLimiter, quit chan bool) {

	for {
		select {
		case <-quit:
			return
		default:
			if rateLimiter(atomic.LoadUint64(&s.completed)) {
				time.Sleep(10 * time.Millisecond)
			} else {
				s.fn()
				atomic.AddUint64(&s.completed, 1)
			}
		}
	}

}

func infiniteLimiter() RateLimiter {
	return func(completed uint64) bool {
		return false
	}
}

func constantLimiter(startNano int64, qpm uint64) RateLimiter {
	return func(completed uint64) bool {
		elapsedNano := time.Now().UnixNano() - startNano
		minutes := float64(elapsedNano) / float64(time.Minute)
		limit := minutes * float64(qpm)
		return float64(completed) > limit
	}
}

func (s *Simulator) calibrate(duration time.Duration) uint64 {

	fmt.Println("Calibration begun...")
	defer fmt.Println("Calibration completed.")

	s.completed = 0

	goroutines := runtime.NumCPU()
	quit := make(chan bool)

	for i := 0; i < goroutines; i++ {
		go s.runner(infiniteLimiter(), quit)
	}

	<-time.After(duration)

	for i := 0; i < goroutines; i++ {
		quit <- true
	}

	return s.completed

}

func (s *Simulator) Run(duration time.Duration) {

	cd := 15 * time.Second
	cycles := s.calibrate(cd)

	fmt.Printf("Calibration Cycles: %d\n", cycles)

	goroutines := runtime.NumCPU()
	quit := make(chan bool)
	qpm := float64(cycles) * float64(time.Minute/cd) * 0.6
	fmt.Printf("Target QPM: %v\n", qpm)
	s.completed = 0

	limiter := constantLimiter(time.Now().UnixNano(), uint64(qpm))
	for i := 0; i < goroutines; i++ {
		go s.runner(limiter, quit)
	}

	if 0 == duration.Nanoseconds() {
		// Run indefinitely
		for {
			time.Sleep(1 * time.Minute)
		}
	} else {
		<-time.After(duration)
		for i := 0; i < goroutines; i++ {
			quit <- true
		}
	}

}
