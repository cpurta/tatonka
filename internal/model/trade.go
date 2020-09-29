package model

import "time"

type Trade struct {
	TradeID string
	Price   string
	Size    string
	Time    time.Time
	Side    string
}

type Trades []*Trade

func (t Trades) Len() int           { return len(t) }
func (t Trades) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t Trades) Less(i, j int) bool { return t[i].Time.Before(t[j].Time) }
