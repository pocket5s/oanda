package oanda

import . "github.com/shopspring/decimal"

type Account struct {
    Id                          string      `json:"ID"`
    Alias                       string      `json:"alias"`
    Currency                    string      `json:"currency"`
    CreatedByUserId             int32       `json:"createdByUserID"`
    CreatedTime                 string      `json:"createdTime"`
    GuaranteedStopLossOrderMode string      `json:"guaranteedStopLossOrderMode"`
    ResettablePLTime            string      `json:"resettablePLTime"`
    MarginRate                  Decimal     `json:"marginRate"`
    OpenTradeCount              int         `json:"openTradeCount"`
    OpenPositionCount           int         `json:"openPositionCount"`
    PendingOderCount            int         `json:"pendingOrderCount"`
    HedgingEnabled              bool        `json:"hedgingEnabled"`
    UnrealizedPL                Decimal     `json:"unrealizedPL"`
    NAV                         string      `json:"NAV"`
    MarginUsed                  Decimal     `json:"marginUsed"`
    PositionValue               Decimal     `json:"positionValue"`
    MarginCloseoutUnrealizedPL  Decimal     `json:"marginCloseoutUnrealizedPL"`
    MarginCloseoutNAV           Decimal     `json:"marginCloseoutNAV"`
    MarginCloseoutMarginUsed    Decimal     `json:"marginCloseoutMarginused"`
    MarginCloseoutPercent       Decimal     `json:"marginCloseoutPercent"`
    MarginCloseoutPositionValue Decimal     `json:"marginCloseoutPositionValue"`
    WithdrawalLimit             Decimal     `json:"withdrawalLimit"`
    MarginCallMarginUsed        Decimal     `json:"marginCallMarginUsed"`
    MarginCallPercent           Decimal     `json:"marginCallPercent"`
    Balance                     Decimal     `json:"balance"`
    Pl                          Decimal     `json:"PL"`
    ResettablePL                Decimal     `json:"resettablePL"`
    Financing                   Decimal     `json:"financing"`
    Commission                  Decimal     `json:"commission"`
    DividendAdjustment          Decimal     `json:"dividendAdjustment"`
    GuaranteedExecutionFees     Decimal     `json:"guaranteedExecutionFees"`
    MarginCallEnterTime         string      `json:"marginCallEnterTime"`
    MarginCallExtensionCount    int         `json:"marginCallExtensionCount"`
    LastMarginCallExtensionTime string      `json:"lastMarginCallExtensionTime"`
    LastTransactionID           string      `json:"lastTransactionID"`
    Trades                      []Trade     `json:"trades"`
    Positions                   []Position  `json:"positions"`
    Orders                      []Order     `json:"orders"`
}

type AccountSummary struct {
    Id                          string  `json:"ID"`
    Alias                       string  `json:"alias"`
    Currency                    string  `json:"currency"`
    CreatedByUserId             string  `json:"createdByUserId"`
    CreatedTime                 string  `json:"createdTime"`
    GuaranteedStopLossOrderMode string  `json:"guaranteedStopLossOrderMode"`
    ResettablePLTime            string  `json:"resettablePLTime"`
    MarginRate                  Decimal `json:"marginRate"`
    OpenTradeCount              int     `json:"openTradeCount"`
    OpenPositionCount           int     `json:"openPositionCount"`
    PendingOrderCount           int     `json:"pendingOrderCount"`
    HedgingEnabled              bool    `json:"hedgingEnabled"`
    UnrealizedPl                Decimal `json:"unrealizedPL"`
    NAV                         Decimal `json:"NAV"`
    MarginUsed                  Decimal `json:"marginUsed"`
    MarginAvailable             Decimal `json:"marginAvailable"`
    PositionValue               Decimal `json:"positionValue"`
    MarginCloseoutUnrealizedPl  Decimal `json:"marginCloseoutUnrealizedPL"`
    MarginClouseoutNAV          Decimal `json:"marginCloseoutNAV"`
    MarginCloseoutMarginUsed    Decimal `json:"marginCloseoutMarginUsed"`
    MarginCloseoutPercent       Decimal `json:"marginCloseoutPercent"`
    MarginCloseoutPositionValue Decimal `json:"marginCloseoutPositionValue"`
    WithdrawalLimit             Decimal `json:"withdrawalLimit"`
    MarginCallMarginUsed        Decimal `json:"marginCallMarginUsed"`
    MarginCallPercent           Decimal `json:"marginCallPercent"`
    Balance                     Decimal `json:"balance"`
    Pl                          Decimal `json:"pl"`
    ResettablePL                Decimal `json:"resettablePL"`
    Financing                   Decimal `json:"financing"`
    Commission                  Decimal `json:"commission"`
    DividendAdjustment          Decimal `json:"dividendAdjustment"`
    GuaranteedExecutionFees     Decimal `json:"guaranteedExecutionFees"`
    MarginCallEnterTime         string  `json:"marginCallEnterTime"`
    MarginCallExtensionCount    int     `json:"marginCallExtensionCount"`
    LastMarginCallExtensionTime string  `json:"lastMarginCallExtensionTime"`
    LastTransactionId           string  `json:"lastTransactionID"`
}

type ClientConfigureTransaction struct {
    Id                  string  `json:"id"`
    Time                string  `json:"time"`
    UserId              int32   `json:"userID"`
    AccountId           string  `json:"accountID"`
    BatchId             string  `json:"batchID"`
    RequestID           string  `json:"requestID"`
    Type                string  `json:"type"`
    Alias               string  `json:"alias"`
    MarginRate          Decimal `json:"marginRate"`
}

type AccountChanges struct {
    OrdersCreated   []Order         `json:"ordersCreated"`
    OrdersCancelled []Order         `json:"ordersCancelled"`
    OrdersFilled    []Order         `json:"ordersFilled"`
    OrdersTriggered []Order         `json:"ordersTriggered"`
    TradesOpened    []TradeSummary  `json:"tradesOpened"`
    TradesReduced   []TradeSummary  `json:"tradesReduced"`
    TradesClosed    []TradeSummary  `json:"tradesClosed"`
    Positions       []Position      `json:"positions"`
    Transactions    []Transaction   `json:"transactions"`
}

type AccountChangesState struct {
    UnrealizedPl                Decimal                     `json:"unrealizedPL"`
    NAV                         Decimal                     `json:"NAV"`
    MarginUsed                  Decimal                     `json:"marginUsed"`
    MarginAvailable             Decimal                     `json:"marginAvailable"`
    PositionValue               Decimal                     `json:"positionValue"`
    MarginCloseoutUnrealizedPl  Decimal                     `json:"marginCloseoutUnrealizedPL"`
    MarginCLoseoutNAV           Decimal                     `json:"marginCloseoutNAV"`
    MarginCloseoutMarginused    Decimal                     `json:"marginCloseoutMarginUsed"`
    MarginCloseoutPercent       Decimal                     `json:"marginCloseoutPercent"`
    MarginCloseoutPositionValue Decimal                     `json:"marginCloseoutPositionValue"`
    WithdrawalLimit             Decimal                     `json:"withdrawalLimit"`
    MarginCallMarginUsed        Decimal                     `json:"marginCallMarginUsed"`
    MarginCallPercent           Decimal                     `json:"marginCallPercent"`
    Balance                     Decimal                     `json:"balance"`
    Pl                          Decimal                     `json:"pl"`
    ResettablePl                Decimal                     `json:"resettablePL"`
    Financing                   Decimal                     `json:"financing"`
    Commission                  Decimal                     `json:"commission"`
    DividendAdjustment          Decimal                     `json:"dividendAdjustment"`
    GuaranteedExecutionFees     Decimal                     `json:"guaranteedExecutionFees"`
    MarginCallEnterTime         string                      `json:"marginCallEnterTime"`
    MarginCallExtensionCount    int                         `json:"marginCallExtensionCount"`
    LastMarginCallExtensionTime string                      `json:"lastMarginCallExtensionTime"`
    Orders                      []DynamicOrderState         `json:"orders"`
    Trades                      []CalculatedTradeState      `json:"trades"`
    Positions                   []CalculatedPositionState   `json:"positions"`
}

// Responses
type AccountResponse struct {
    Account             Account `json:"account"`
    LastTransactionId   string  `json:"lastTransactionID"`
}

type AccountSummaryResponse struct {
    Account             AccountSummary  `json:"account"`
    LastTransactionId   string          `json:"lastTransactionID"`
}

type InstrumentsResponse struct {
    Instruments         []string    `json:"instruments"`
    LastTransactionId   string      `json:"lastTransactionID"`
}

type AccountPatchRequest struct {
    Alias       string  `json:"alias"`
    MarginRate  Decimal `json:"marginRate"`
}

type AccountPatchResponse struct {
    ClientConfigureTransaction  ClientConfigureTransaction  `json:"clientConfigureTransaction"`
    LastTransactionId           string                      `json:"lastTransactionID"`
}

type AccountChangesResponse struct {
    Changes             AccountChanges      `json:"changes"`
    State               AccountChangesState `json:"accountChangesState"`
    LastTransactionId   string              `json:"lastTransactionID"`
}   
