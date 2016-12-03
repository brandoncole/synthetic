package simulator

import (
	"runtime"
	"time"
    "fmt"
)

type Simulator struct {
	fn func()
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func NewSimulator(simulation func()) *Simulator {
	return &Simulator{
		fn: simulation,
	}
}

func (s *Simulator) runner(quit chan bool, throughput chan uint64) {

	var qty uint64

	for {
		select {
		case <-quit:
			throughput <- qty
			return
		default:
			s.fn()
			qty++
		}
	}

}

func (s *Simulator) calibrate() {

	goroutines := runtime.NumCPU() * 8
	quit := make(chan bool)
	results := make(chan uint64)

	for i := 0; i < goroutines; i++ {
		go s.runner(quit, results)
	}

	<-time.After(30 * time.Second)

	for i := 0; i < goroutines; i++ {
		quit <- true
	}

    var total uint64
    for i:= 0 ; i < goroutines; i++ {
        total += <-results
    }

    fmt.Printf("Calibration throughput: %d\n", total)

}

func (s *Simulator) Run() {

	fmt.Println("Simulation beginning...")
	s.calibrate()
	fmt.Println("Simulation completed.")

}
