package model

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

type Trade struct {
	TradeID string
	Price   float64
	Size    float64
	Time    time.Time
	Side    string
}

func (t *Trade) ID() string {
	hash := md5.New()
	io.WriteString(hash, t.TradeID)
	io.WriteString(hash, strconv.FormatFloat(t.Price, 'f', 5, 64))
	io.WriteString(hash, strconv.FormatFloat(t.Size, 'f', 5, 64))
	io.WriteString(hash, t.Time.Format(time.RFC3339))
	io.WriteString(hash, t.Side)
	return fmt.Sprintf("%x", hash.Sum(nil))[:16]
}

type Trades []*Trade

func (t Trades) Len() int           { return len(t) }
func (t Trades) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t Trades) Less(i, j int) bool { return t[i].Time.Before(t[j].Time) }
