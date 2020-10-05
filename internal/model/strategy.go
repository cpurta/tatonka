package model

// Strategy represents a market strategy to determine a signal in the market.
type Strategy interface {
	Name() string
	Description() string
	Options() []Option
	Calculate(periods []*Period) (float64, error)
	Signal(float64) Signal
}
