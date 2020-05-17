package oanda

import . "github.com/shopspring/decimal"

type Position struct {
    Instrument              string          `json:"instrument"`
    Pl                      Decimal         `json:"pl"`
    UnrealizedPl            Decimal         `json:"unrealizedPL"`
    MarginUsed              Decimal         `json:"marginUsed"`
    ResettablePl            Decimal         `json:"resettablePL"`
    Financing               Decimal         `json:"financing"`
    Commission              Decimal         `json:"commission"`
    DividendAdjustment      Decimal         `json:"dividendAdjustment"`
    GuaranteedExecutionFees Decimal         `json:"guaranteedExecutionFees"`
    Long                    PositionSide    `json:"long"`
    Short                   PositionSide    `json:"short"`
}

type PositionSide struct {
    Units                   Decimal     `json:"units"`
    AveragePrice            Decimal     `json:"averagePrice"`
    TradeIds                []string    `json:"tradeIDs"`
    Pl                      Decimal     `json:"pl"`
    UnrealizedPl            Decimal     `json:"unrealizedPL"`
    ResettablePl            Decimal     `json:"resettablePL"`
    Financing               Decimal     `json:"financing"`
    DividenAdjustment       Decimal     `json:"dividendAdjustment"`
    GuaranteedExecutionFees Decimal     `json:"guaranteedExecutionFees"`
}

type PositionsResponse struct {
    Positions           []Position  `json:"positions"`
    LastTransactionId   string      `json:"lastTransactionID"`
}

type InstrumentClose struct {
    LongUnits               string              `json:"longUnits"`
    LongClientExtensions    []ClientExtensions  `json:"longClientExtensions"`
    ShortUnits              string              `json:"shortUnits"`
    ShortClientExtensions   []ClientExtensions  `json:"shortClientExtensions"`
}

type CalculatedPositionState struct {
}

type PositionsResponse struct {
    Positions           []Position  `json:"positions"`
    LastTransactionId   string      `json:"lastTransactionID"`
}

type PositionResponse struct {
    Position            Position  `json:"positions"`
    LastTransactionId   string    `json:"lastTransactionID"`
}

type CloseoutPositionRequest struct {
    LongUnits               string              `json:"longUnits"`
    LongClientExtensions    ClientExentsions    `json:"longClientExtensions"`
    ShortUnits              string              `json:"shortUnits"`
    ShortClientExtensions   ClientExtensions    `json:"shortClientExtensions"`
}

type CloseoutPositionResponse struct {
    LongOrderCreateTransaction  MarketOrderTransaction  `json:"longOrderCreateTransaction"`
    LongOrderFillTransaction    OrderFillTransaction    `json:"longOrderFillTransaction"`
    LongOrderCancelTransaction  OrderCancelTransaction  `json:"longOrderCancelTransaction"`
    ShortOrderCreateTransaction MarketOrderTransaction  `json:"shortOrderCreateTransaction"`
    ShortOrderFillTransaction   OrderFillTransaction    `json:"shortOrderFillTransaction"`
    ShortOrderCancelTransaction OrderCancelTransaction  `json:"shortOrderCancelTransaction"`
    RelatedTransactionIDs       []string                `json:"relatedTransactionIDs"`
    LastTransactionID           string                  `json:"lastTransactionID"`
}

// Received from a 400 or 404
type CloseoutPositionInvalidResponse struct {
   LongOrderRejectTransaction   MarketOrderRejectTransaction    `json:"longOrderRejectTransaction"`
   ShortOrderRejectTransaction  MarketOrderRejectTransaction    `json:"shortOrderRejectTransaction"`
   RelatedTransactionIDs        []string                        `json:"relatedTransactionIDs"`
   LastTransactionId            string                          `json:"lastTransactionID"`
   ErrorCode                    string                          `json:"errorCode"`
   ErrorMessage                 string                          `json:"errorMessage"`
}
