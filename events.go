package iticapital

import (
	"github.com/devbackend/iticapital/pkg/datetime"
	"github.com/devbackend/iticapital/pkg/money"
)

type OrderState uint8

const (
	Unknown OrderState = iota
	Pending
	Open
	Expired
	Cancel
	Filled
	Partial
	ContractorCancel
	SystemReject
	SystemCancel
)

type AddBar struct {
	Row        uint64                  `json:"row"`
	Symbol     string                  `json:"symbol"`
	Interval   GetBarsRequest_Interval `json:"interval"`
	DateTime   datetime.DateTime       `json:"date_time"`
	OpenPrice  money.Money             `json:"open_price"`
	HighPrice  money.Money             `json:"high_price"`
	LowPrice   money.Money             `json:"low_price"`
	ClosePrice money.Money             `json:"close_price"`
	Volume     float64                 `json:"volume"`
	Opened     float64                 `json:"opened"`
}

type AddPortfolio struct {
	Row       int             `json:"row"`
	RowsCount int             `json:"rows_count"`
	Name      string          `json:"name"`
	Exchange  string          `json:"exchange"`
	Status    PortfolioStatus `json:"status"`
}

type AddSymbol struct {
	Row              int               `json:"row"`
	RowsCount        int               `json:"rows_count"`
	ISIN             string            `json:"isin"`
	Symbol           string            `json:"symbol"`
	ShortName        string            `json:"short_name"`
	LongName         string            `json:"long_name"`
	Type             string            `json:"type"`
	Decimals         uint8             `json:"decimals"`
	LotSize          int64             `json:"lot_size"`
	PriceStep        float64           `json:"price_step"`
	ExchangeName     string            `json:"exchange_name"`
	ExpiryDate       datetime.DateTime `json:"expiry_date"`
	DaysBeforeExpiry float64           `json:"days_before_expiry"`
	PriceStepPoint   float64           `json:"price_step_point"`
	Strike           float64           `json:"strike"`
}

type AddTick struct {
	Symbol   string                   `json:"symbol"`
	DateTime datetime.DateTime        `json:"date_time"`
	Price    float64                  `json:"price"`
	Volume   float64                  `json:"volume"`
	TradeID  string                   `json:"trade_id"`
	Action   PlaceOrderRequest_Action `json:"action"`
}

type AddTickHistory struct {
	Row       int                      `json:"row"`
	RowsCount int                      `json:"rows_count"`
	Symbol    string                   `json:"symbol"`
	DateTime  datetime.DateTime        `json:"date_time"`
	Price     float64                  `json:"price"`
	Volume    float64                  `json:"volume"`
	TradeID   string                   `json:"trade_id"`
	Action    PlaceOrderRequest_Action `json:"action"`
}

type AddTrade struct {
	Portfolio     string            `json:"portfolio"`
	Symbol        string            `json:"symbol"`
	OrderID       string            `json:"order_id"`
	Price         money.Money       `json:"price"`
	Amount        float64           `json:"amount"`
	DateTime      datetime.DateTime `json:"date_time"`
	TradeID       string            `json:"trade_id"`
	Value         money.Money       `json:"value"`
	AccruedCoupon money.Money       `json:"accrued_coupon"`
}

type OrderCancelFailed struct {
	OrderID string
}

type OrderCancelSucceeded struct {
	OrderID string
}

type OrderMoveFailed struct {
	OrderID string
}

type OrderMoveSucceeded struct {
	OrderID string
}

type OrderFailed struct {
	OrderID  string `json:"order_id"`
	UniqueID int    `json:"unique_id"`
	Reason   string `json:"reason"`
}

type OrderSucceeded struct {
	OrderID  string `json:"order_id"`
	UniqueID int    `json:"unique_id"`
}

type SetMyClosePosition struct {
	Row          int               `json:"row"`
	RowsCount    int               `json:"rows_count"`
	Portfolio    string            `json:"portfolio"`
	Symbol       string            `json:"symbol"`
	Amount       float64           `json:"amount"`
	PriceBuy     money.Money       `json:"price_buy"`
	PriceSell    money.Money       `json:"price_sell"`
	DateTime     datetime.DateTime `json:"date_time"`
	OpenOrderID  string            `json:"open_order_id"`
	CloseOrderID string            `json:"close_order_id"`
}

type SetMyOrder struct {
	Row          int                        `json:"row"`
	RowsCount    int                        `json:"rows_count"`
	Portfolio    string                     `json:"portfolio"`
	Symbol       string                     `json:"symbol"`
	Amount       float64                    `json:"amount"`
	Price        money.Money                `json:"price"`
	PriceStop    money.Money                `json:"price_stop"`
	FilledVolume float64                    `json:"filled_volume"`
	DateTime     datetime.DateTime          `json:"date_time"`
	OrderID      string                     `json:"order_id"`
	StockOrderID string                     `json:"stock_order_number"`
	UniqueID     int                        `json:"unique_id"`
	State        OrderState                 `json:"state"`
	Action       PlaceOrderRequest_Action   `json:"action"`
	Type         PlaceOrderRequest_Type     `json:"type"`
	Validity     PlaceOrderRequest_Validity `json:"validity"`
}

type SetMyTrade struct {
	Row           int                      `json:"row"`
	RowsCount     int                      `json:"rows_count"`
	Portfolio     string                   `json:"portfolio"`
	Symbol        string                   `json:"symbol"`
	DateTime      datetime.DateTime        `json:"date_time"`
	Price         money.Money              `json:"price"`
	Volume        float64                  `json:"volume"`
	OrderID       string                   `json:"order_id"`
	StockOrderID  string                   `json:"stock_order_id"`
	Value         money.Money              `json:"value"`
	AccruedCoupon money.Money              `json:"accrued_coupon"`
	Action        PlaceOrderRequest_Action `json:"action"`
}

type SetPortfolio struct {
	Portfolio        string      `json:"portfolio"`
	Leverage         float64     `json:"leverage"`
	Commission       float64     `json:"commission"`
	Money            money.Money `json:"money"`
	Profit           money.Money `json:"profit"`
	LiquidationPrice money.Money `json:"liquidation_price"`
	InitialMargin    money.Money `json:"initial_margin"`
	TotalAssets      money.Money `json:"total_assets"`
}

type UpdateOrder struct {
	Portfolio    string                     `json:"portfolio"`
	Symbol       Symbol                     `json:"symbol"`
	State        OrderState                 `json:"state"`
	Action       PlaceOrderRequest_Action   `json:"action"`
	Type         PlaceOrderRequest_Type     `json:"type"`
	Validity     PlaceOrderRequest_Validity `json:"validity"`
	Price        money.Money                `json:"price"`
	Amount       float64                    `json:"amount"`
	PriceStop    money.Money                `json:"price_stop"`
	FilledVolume float64                    `json:"filled_volume"`
	DateTime     datetime.DateTime          `json:"date_time"`
	OrderID      string                     `json:"order_id"`
	StockOrderID string                     `json:"stock_order_number"`
	StatusMask   int                        `json:"status_mask"`
	UniqueID     int                        `json:"unique_id"`
	Error        string                     `json:"error"`
}

type UpdatePosition struct {
	Portfolio     string      `json:"portfolio"`
	Symbol        string      `json:"symbol"`
	AveragePrice  money.Money `json:"average_price"`
	Amount        float64     `json:"amount"`
	PlannedAmount money.Money `json:"planned_amount"`
}

type UpdateBidAsks struct {
	Row      uint64      `json:"row"`
	Rows     int64       `json:"rows"`
	Symbol   string      `json:"symbol"`
	BidPrice money.Money `json:"bid_price"`
	BidSize  float64     `json:"bid_size"`
	AskPrice money.Money `json:"ask_price"`
	AskSize  float64     `json:"ask_size"`
}

type UpdateQuote struct {
	Symbol           string            `json:"symbol"`
	DateTime         datetime.DateTime `json:"date_time"`
	PriceOpen        money.Money       `json:"price_open"`
	PriceHigh        money.Money       `json:"price_high"`
	PriceLow         money.Money       `json:"price_low"`
	PriceClose       money.Money       `json:"price_close"`
	PriceLast        money.Money       `json:"price_last"`
	LastSize         float64           `json:"last_size"`
	SessionVolume    float64           `json:"session_volume"`
	PriceBid         money.Money       `json:"price_bid"`
	PriceAsk         money.Money       `json:"price_ask"`
	SizeBid          float64           `json:"size_bid"`
	SizeAsk          float64           `json:"size_ask"`
	OpenedPrice      money.Money       `json:"opened_price"`
	GuaranteeBuy     money.Money       `json:"guarantee_buy"`
	GuaranteeSell    money.Money       `json:"guarantee_sell"`
	PriceHighLimit   money.Money       `json:"price_high_limit"`
	PriceLowLimit    money.Money       `json:"price_low_limit"`
	IsTrading        bool              `json:"is_trading"`
	Volatility       float64           `json:"volatility"`
	TheoreticalPrice money.Money       `json:"theoretical_price"`
	StepPrice        money.Money       `json:"step_price"`
}
