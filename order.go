package oanda

import (
    "encoding/json"
    . "github.com/shopspring/decimal"
)

type Order struct {
    Id               string          `json:"ID"`
    CreateTime       string          `json:"createTime"`
    State            string          `json:"state"`
    ClientExtensions ClientExtension `json:"clientExtensions"`
}

type MarketOrder struct {
    Order
    Type                    string                          `json:"type"`
    Instrument              string                          `json:"instrument"`
    Units                   Decimal                         `json:"units"`
    TimeInForce             string                          `json:"timeInForce"`
    PriceBound              Decimal                         `json:"priceBound"`
    PositionFill            string                          `json:"positionFill"`
    TradeClose              MarketOrderTradeClose           `json:"tradeClose"`
    LongPositionCloseout    MarketOrderPositionCloseout     `json:"longPositionCloseout"`
    ShortPositionCloseout   MarketOrderPositionCloseout     `json:"shortPositionCloseout"`
    MarginCloseout          MarketOrderMarginCloseout       `json:"maringClouseout"`
    DelayedTradeClose       MarketORderDelayedTradeCloseout `json:"delayedTradeClose"`
    TakeProfitOnFill        TakeProfitDetails               `json:"takeProfitOnFill"`
    StopLossOnFill          StopLossDetails                 `json:"stopLossOnFill"`
    TrailingStopLossOnFill  TrailingStopLossDetails         `json:"trailingStopLossOnFill"`
    TradeClientExtensions   ClientExtensions                `json:"tradeClientExtensions"`
    FillingTransactionId    string                          `json:"fillingTransactinoID"`
    FilledTime              string                          `json:"filledTime"`
    TradeOpenedId           string                          `json:"tradeOpenedID"`
    TradeReducedId          string                          `json:"tradeReducedID"`
    TradeClosedIds          []string                        `json:"tradeClosedIDs"`
    CancellingTransactionId string                          `json:"cancellingTransactionID"`
    CancelledTime           string                          `json:"cancelledTime"`
}

type MarketOrderTradeClose struct {
    TradeId       string `json:"tradeID"`
    ClientTradeId string `json:"clientTradeID"`
    Units         string `json:"units"`
}

type MarketOrderPositionCloseout struct {
    Instrument string `json:"instrument"`
    Units      string `json:"units"`
}

type MarketOrderMarginCloseout struct {
    Reason string `json:"reason"`
}

type MarketOrderDelayedTradeClose struct {
    TradeId       string `json:"tradeID"`
    ClientTradeId string `json:"clientTradeID"`
    SourceTradeId string `json:"sourceTradeID"`
}

type FixedPriceOrder struct {
    Order
    Type                    string                  `json:"type"`
    Instrument              string                  `json:"instrument"`
    Units                   Decimal                 `json:"units"`
    Price                   Decimal                 `json:"price"`
    PositionFill            string                  `json:"positionFill"`
    TradeState              string                  `json:"tradeState"`
    TakeProfitOnFil         TakeProfitDetails       `json:"takeProfitOnFill"`
    StopLossOnFill          StopLossDetails         `json:"stopLossOnFill"`
    TrailingStopLossOnFill  TrailingStopLossDetails `json:"trailingStopLossOnFill"`
    TradeClientExtensions   ClientExtensions        `json:"tradeClientExtensions"`
    FillingTransactionId    string                  `json:"fillingTransactionId"`
    FilledTime              string                  `json:"filledtime"`
    TradeOpenedId           string                  `json:"tradeOpenID"`
    TradeReducedId          string                  `json:"tradeReducedID"`
    TradeClosedIds          []string                `json:"tradeClosedIDs"`
    CancellingTransactionId string                  `json:"cancellingTransactionID"`
    CancelledTime           string                  `json:"cancelledTime"`
}

type DynamicOrderState struct {
    Id                      string  `json:"id"`
    TrailingStopValue       Decimal `json:"trailingStopValue"`
    TriggerDistance         Decimal `json:"triggerDistance"`
    IsTriggerDistanceExact  bool    `json:"isTriggerDistanceExact"`
}

// Different order requests
type OrderRequest struct {
    Type                    string                  `json:"type"`
    Instrument              string                  `json:"instrument"`
    Units                   Decimal                 `json:"units"`
    TimeInForce             string                  `json:"timeInForce"`
    PositionFill            string                  `json:"positionFill"`
    TakeProfitOnFill        TakeProfitDetails       `json:"takeProfitOnFill"`
    StopLossOnFill          StopLossDetails         `json:"stopLossOnFill"`
    TrailingStopLossOnFill  TrailingStopLossDetails `json:"trailingStopLossOnFill"`
    TradeClientExtensions   ClientExtension         `json:"tradeClientExtensions"`
    ClientExtensions        ClientExtension         `json:"clientExtensions"`
}

type MarketOrderRequest struct {
    OrderRequest
    PriceBound Decimal `json:"priceBound"`
}

type LimitOrderRequest struct {
    OrderRequest
    Price            Decimal `json:"price"`
    GtdTime          string  `json:"gtdTime"`
    TriggerCondition string  `json:"triggerCondition"`
}

type StopOrderRequest struct {
    OrderRequest
    Price            Decimal `json:"price"`
    PriceBound       Decimal `json:"priceBound"`
    TriggerCondition string  `json:"triggerCondition"`
}

type MarketIfTouchedOrderRequest struct {
    OrderRequest
    Price       Decimal `json:"price"`
    PriceBound  Decimal `json:"priceBound"`
    GtdTime     string  `json:"gtdTime"`
}

type TakeProfitOrderRequest struct {
    Type             string          `json:"type"`
    TradeId          string          `json:"tradeID"`
    ClientTradeId    string          `json:"clientTradeID"`
    Price            Decimal         `json:"price"`
    GtdTime          string          `json:"gtdTime"`
    TriggerCondition string          `json:"triggerCondition"`
    ClientExtensions ClientExtension `json:"clientExtensions"`
}

type StopLossOrderRequest struct {
    TakeProfitOrderRequest
    Distance    Decimal `json:"distance"`
    TimeInForce string  `json:"timeInForce"`
}

type TrailingStopLossOrderRequest struct {
    Type             string          `json:"type"`
    TradeId          string          `json:"tradeID"`
    ClientTradeId    string          `json:"clientTradeID"`
    Distance         Decimal         `json:"distance"`
    GtdTime          string          `json:"gtdTime"`
    TriggerCondition string          `json:"triggerCondition"`
    ClientExtensions ClientExtension `json:"clientExtensions"`
}

type ProfitDetails struct {
    TimeInForce      string          `json:"timeInForce"`
    GtdTime          string          `json:"gtdTime"`
    ClientExtensions ClientExtension `json:"clientExtensions"`
}

type TakeProfitDetails struct {
    ProfitDetails
    Price Decimal `json:"price"`
}

type TrailingStopLossDetails struct {
    ProfitDetails
    Distance Decimal `json:"distance"`
}

type StopLossDetails struct {
    TrailingStopLossDetails
    Price Decimal `json:"price"`
}

// Order responses
type OrdersResponse struct {
    Orders            []Order `json:"orders"`
    LastTransactionId string  `json:"lastTransactionID"`
}

// helper functions
func NewMarketOrder() MarketOrderRequest {
    o := MarketOrderRequest {
        Type: "MARKET"
    }
    return o
}

func NewLimitOrder() MarketOrderRequest {
    o := LimitOrderRequest {
        Type: "LIMIT"
    }
    return o
}

func NewStopOrder() MarketOrderRequest {
    o := StopOrderRequest {
        Type: "STOP"
    }
    return o
}

func NewMarketIfTouchedOrder() MarketOrderRequest {
    o := MarketIfTouchedOrderRequest {
        Type: "MARKET_IF_TOUCHED"
    }
    return o
}

func NewTakeProfitOrder() MarketOrderRequest {
    o := TakeProfitOrderRequest {
        Type: "TAKE_PROFIT"
    }
    return o
}

func NewStopLossOrder() MarketOrderRequest {
    o := StopLossOrderRequest {
        Type: "STOP_LOSS"
    }
    return o
}

func NewTrailingStopLossOrder() MarketOrderRequest {
    o := TrailingStopLossOrderRequest {
        Type: "TRAILING_STOP_LOSS"
    }
    return o
}
