package chart

import (
	"strconv"
)

// OHLC is a Candlestick having buy and sell volume too.
type OHLC struct {
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
	Vbuy      float64 `json:"vbuy"`
	Vsell     float64 `json:"vsell"`
	Range     uint16  `json:"range"`
}

// OHLCs is collection of OHLC
type OHLCs []OHLC

func (ohlc *OHLC) Update(price float64, quantity float64, side string) {
	ohlc.Close = price
	if ohlc.Open == 0 {
		ohlc.Open = price
	}
	if ohlc.High < price {
		ohlc.High = price
	}
	if ohlc.Low == 0 || ohlc.Low > price {
		ohlc.Low = price
	}
	if side == "buy" {
		ohlc.Vbuy += quantity
	} else if side == "sell" {
		ohlc.Vsell += quantity
	}
}

func (ohlc *OHLC) ToArray() []string {
	open := strconv.FormatFloat(ohlc.Open, 'f', 4, 64)
	high := strconv.FormatFloat(ohlc.High, 'f', 4, 64)
	low := strconv.FormatFloat(ohlc.Low, 'f', 4, 64)
	close := strconv.FormatFloat(ohlc.Close, 'f', 4, 64)
	timestamp := strconv.FormatInt(ohlc.Timestamp, 10)
	return []string{
		open,
		high,
		low,
		close,
		timestamp,
	}
}
