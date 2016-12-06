// Package Simulator is where the implementations of the simulators are located.
package simulator

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// Simulator is the base type that all other simulators are derived from.
type Simulator struct {
	// Duration specifies the amount of time that the simulation will run.  0s is interpreted as infinite.
	Duration time.Duration
}

// ThroughputSimulator is a specific type of simulator that requires throughput calibration to determine the extrema.
type ThroughputSimulator struct {
	Simulator
	CalibrationDuration  time.Duration
	PeriodDuration       time.Duration
	limiter              IThroughputLimiter
	simulation           func()
	simulationsCompleted uint64
	simulationsPerPeriod float64
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// NewThroughputSimulator creates a new throughput simulator
func NewThroughputSimulator(limiter IThroughputLimiter, simulation func()) *ThroughputSimulator {
	return &ThroughputSimulator{
		limiter:    limiter,
		simulation: simulation,
	}
}

func (s *ThroughputSimulator) runner(limiter IThroughputLimiter, quit chan bool) {

	startNano := time.Now().UnixNano()

	state := NewThroughputLimiterState()
	state.CyclesPerPeriod = s.simulationsPerPeriod

	for {

		elapsed := time.Now().UnixNano() - startNano
		state.PeriodsCompleted = float64(elapsed) / float64(s.PeriodDuration.Nanoseconds())
		state.CyclesCompleted = atomic.LoadUint64(&s.simulationsCompleted)

		select {
		case <-quit:
			return
		default:
			if limiter.Throttled(state) {
				time.Sleep(10 * time.Millisecond)
			} else {
				s.simulation()
				atomic.AddUint64(&s.simulationsCompleted, 1)
			}
		}
	}

}

func (s *ThroughputSimulator) calibrate(duration time.Duration) uint64 {

	fmt.Println("Calibration begun...")
	defer fmt.Println("Calibration completed.")

	s.simulationsCompleted = 0

	goroutines := runtime.NumCPU()
	quit := make(chan bool)

	unlimited := NewThroughputLimiterUnlimited()
	for i := 0; i < goroutines; i++ {
		go s.runner(unlimited, quit)
	}

	<-time.After(duration)

	for i := 0; i < goroutines; i++ {
		quit <- true
	}

	return s.simulationsCompleted

}

func (s *ThroughputSimulator) Run() {

	cycles := s.calibrate(s.CalibrationDuration)
	qpp := float64(cycles) * float64(s.PeriodDuration/s.CalibrationDuration)
	s.simulationsPerPeriod = qpp
	fmt.Printf("Calibration Cycles: %d, Cycles Per Period: %f\n", cycles, qpp)

	goroutines := runtime.NumCPU()
	quit := make(chan bool)
	s.simulationsCompleted = 0

	for i := 0; i < goroutines; i++ {
		go s.runner(s.limiter, quit)
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
