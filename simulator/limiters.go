package simulator

import (
	"math"
)

// ThroughputLimiterState defines an instantaneous state on which to base throughput decision on
type ThroughputLimiterState struct {
	// Set during the calibration period, indicates the maximum throughput in Cycles / Period
	CyclesPerPeriod float64
	// Number of cycles that have been completed in the current simulation
	CyclesCompleted uint64
	// Amount of time that has passed in the current simulation
	PeriodsCompleted float64
}
func NewThroughputLimiterState() *ThroughputLimiterState{
    return &ThroughputLimiterState{}
}

// IThroughputLimiter is an interface implemented by limiters to throttle throughput
type IThroughputLimiter interface {
	// Receives information about how many cycles have completed and returns whether the caller should throttle itself. Returns false if the caller has not been throttled and can continue. Returns true if the caller is proceeding too fast and should temporarily pause.
	Throttled(state *ThroughputLimiterState) bool
}

// ThroughputLimiterUnlimited
type ThroughputLimiterUnlimited struct{}

func NewThroughputLimiterUnlimited() *ThroughputLimiterUnlimited {
	return &ThroughputLimiterUnlimited{}
}
func (l *ThroughputLimiterUnlimited) Throttled(state *ThroughputLimiterState) bool {
	return false
}

// ThroughputLimiterFlat
type ThroughputLimiterFlat struct {
	percentage float64
}

func NewThroughputLimiterFlat(percentage float64) *ThroughputLimiterFlat {
	return &ThroughputLimiterFlat{
		percentage: percentage,
	}
}
func (l *ThroughputLimiterFlat) Throttled(state *ThroughputLimiterState) bool {
	limit := state.PeriodsCompleted * state.CyclesPerPeriod * l.percentage
	return float64(state.CyclesCompleted) > limit
}

// ThroughputLimiterSine
type ThroughputLimiterSine struct {
	min float64
	max float64
}

func NewThroughputLimiterSine(min, max float64) *ThroughputLimiterSine {
	return &ThroughputLimiterSine{
		min: min,
		max: max,
	}
}
func (l *ThroughputLimiterSine) Throttled(state *ThroughputLimiterState) bool {

	// Integrals calculated from http://www.wolframalpha.com
	magnitude := (l.max - l.min) * state.CyclesPerPeriod
	magnitudeOverTwo := magnitude / 2.0
	magnitudeOverFour := magnitudeOverTwo / 2.0
	minimum := l.min * state.CyclesPerPeriod

	limita := (magnitudeOverTwo + minimum) * state.PeriodsCompleted
	limitb := (magnitudeOverFour * math.Sin(2*math.Pi*state.PeriodsCompleted)) / math.Pi
	limit := limita - limitb

	return float64(state.CyclesCompleted) > limit

}