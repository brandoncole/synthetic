package simulator

import (
	"fmt"
	"math"
	"runtime"
	"sync/atomic"
	"time"
)

type Simulator struct {
	fn        func()
	completed uint64
	Duration  time.Duration
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

func sineLimiter(startNano int64, qpm uint64, min, max float64) RateLimiter {

	// Integrals calculated from http://www.wolframalpha.com
	magnitude := (max - min) * float64(qpm)
	magnitudeOverTwo := magnitude / 2.0
	magnitudeOverFour := magnitudeOverTwo / 2.0
	minimum := min * float64(qpm)

	return func(completed uint64) bool {

		elapsedNano := time.Now().UnixNano() - startNano
		minutes := float64(elapsedNano) / float64(time.Minute)
		limita := (magnitudeOverTwo + minimum) * minutes
		limitb := (magnitudeOverFour * math.Sin(2*math.Pi*minutes)) / math.Pi
		limit := limita - limitb

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

func (s *Simulator) Run() {

	cd := 10 * time.Second
	cycles := s.calibrate(cd)

	fmt.Printf("Calibration Cycles: %d\n", cycles)

	goroutines := runtime.NumCPU()
	quit := make(chan bool)
	qpp := float64(cycles) * float64(time.Minute/cd)
	fmt.Printf("Target QPP: %v\n", qpp)
	s.completed = 0

	//limiter := constantLimiter(time.Now().UnixNano(), uint64(qpm))
	limiter := sineLimiter(time.Now().UnixNano(), uint64(qpp), 0.1, 0.9)
	for i := 0; i < goroutines; i++ {
		go s.runner(limiter, quit)
	}

	if 0 == s.Duration.Nanoseconds() {
		var forever chan bool
		<-forever
	} else {
		<-time.After(s.Duration)
		for i := 0; i < goroutines; i++ {
			quit <- true
		}
	}

}
