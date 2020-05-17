package oanda

import . "github.com/shopspring/decimal"

type TradeSummary struct {
    Id                      string              `json:"id"`
    Instrument              string              `json:"instrument"`
    Price                   Decimal             `json:"price"`
    OpenTime                string              `json:"openTime"`
    State                   string              `json:"state"`
    InitialUnits            Decimal             `json:"initialUnits"`
    InitialMarginRequired   Decimal             `json:"initialMarginRequired"`
    CurrentUnits            Decimal             `json:"currentUnits"`
    RealizedPL              Decimal             `json:"realizedPL"`
    UnrealizedPL            Decimal             `json:"unrealizedPL"`
    MarginUsed              Decimal             `json:"marginUsed"`
    AverageClosePrice       Decimal             `json:"averageClosePrice"`
    ClosingTransactionIds   []string            `json:"closingTransactionIDs"`
    Financing               Decimal             `json:"financing"`
    DividendAdjustment      Decimal             `json:"dividendAdjustment"`
    CloseTime               string              `json:"closeTime"`
    ClientExtensions        []ClientExtension   `json:"clientExtensions"`
    TakeProfitOrderId       string              `json:"takeProfitOrderID"`
    StopLossOrderId         string              `json:"stopLossOrderID"`
    TrailingStorOrderId     string              `json:"trailingStopOrderID"`
}

type ClientExtension struct {
    Id      string `json:"id"`
    Tag     string `json:"tag"`
    Comment string `json:"comment"`
}

type CalculatedTradeState struct {
}


// Requests and responses
type TradesResponse struct {
    Trades              []Trade `json:"trades"`
    LastTransactionId   string  `json:"lastTransactionID"`
}

type TradesSpecifierResponse struct {
    Trade               Trade   `json:"trade"`
    LastTransactionId   string  `json:"lastTransactionID"`
}

type CloseTradeRequest struct {
    Units   string `json:"units"`
}

type CloseTradeResponse struct {
    OrderCreateTransaction  MarketOrderTransaction  `json:"orderCreateTransaction"`
    OrderFillTransaction    OrderFillTransaction    `json:"orderFillTransaction"`
    OrderCancelTransaction  OrderCancelTransaction  `json:"orderCancelTransaction"`
    RelatedTransactionIDs   []string                `json:"relatedTransactionIDs"`
    LastTransactionId       string                  `json:"lastTransactionID"`
}

// Received on a 400
type CloseTradeRejectedResponse struct {
    OrderRejectTransaction  MarketOrderTransactino  `json:"orderRejectTransaction"`
    ErrorCode               string                  `json:"errorCode"`
    ErrorMessage            string                  `json:"errorMessage"`
}

// Received on a 404
type CloseTradeNotFoundResponse struct {
    CloseTradeRejectedResponse
    RelatedTransactionIDs   []string    `json:"relatedTransactionIDs"`
}

type ClientExtensionsRequest struct {
    ClientExtensions    ClientExtension `json:"clientExtensions"`
}

type UpdateOrdersThroughTradeRequest struct {
    TakeProfit          TakeProfitDetails       `json:"takeProfit"`
    StopLoss            StopLossDetails         `json:"stopLoss"`
    TrailingStopLoss    TrailingStopLossDetails `json:"trailingStopLoss"`
}

