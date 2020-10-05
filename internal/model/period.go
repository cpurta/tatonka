package model

// Period holds high level information about exchange trades with a certain time
// "period". Periods will be used to feed into strategies to determine market
// signals
type Period struct {
	High     float64
	Low      float64
	Close    float64
	Selector string
}
