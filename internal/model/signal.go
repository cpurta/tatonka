package model

// Signal represents a market signal that indicates whether a market is currently
// bullish (going up) or bearish (going down)
type Signal int

const (
	BuySignal Signal = iota
	SellSignal
	NeutralSignal
)
