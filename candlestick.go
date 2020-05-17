
package oanda

import . "github.com/shopspring/decimal"

type CandleStick struct {
    Time string
    Bid CandleStickData
    Ask CandleStickData
    Mid CandleStickData
    Volume int32
    Complete bool
}

type CandleStickData struct {
    O Decimal
    H Decimal
    L Decimal
    C Decimal
}

type InstrumentCandlesResponse struct {
    Instrument string
    Granularity string
    Candles []CandleStick
}
