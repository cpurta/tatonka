package model

type Strategy interface {
	Name() string
	Description() string
	Options() []Option
	Calculate(periods []*Period) (float64, error)
	Signal(float64) Signal
}
