package coingate

import (
	"time"
)

const (
	// live url
	ApiLiveUrl = "https://api.coingate.com/v2"
	// sand box url
	ApiSandBoxUrl = "https://api-sandbox.coingate.com/v2"
)

// https://developer.coingate.com/reference/order-statuses
const (
	StatusNew               = "new"
	StatusPending           = "pending"
	StatusConfirming        = "confirming"
	StatusPaid              = "paid"
	StatusInvalid           = "invalid"
	StatusExpired           = "expired"
	StatusCanceled          = "canceled"
	StatusRefunded          = "refunded"
	StatusPartiallyRefunded = "partially_refunded"
)

type (
	Client struct {
		AppId   int    // mecharnt app id
		Token   string // app token
		BaseUrl string // api url
		Timeout time.Duration
	}

	CreateOrderRequest struct {
		OrderID         string  `json:"order_id"`
		PriceAmount     float64 `json:"price_amount"`
		PriceCurrency   string  `json:"price_currency"`
		ReceiveCurrency string  `json:"receive_currency"`
		Title           string  `json:"title"`
		Description     string  `json:"description"`
		CallbackURL     string  `json:"callback_url"`
		CancelURL       string  `json:"cancel_url"`
		SuccessURL      string  `json:"success_url"`
		Token           string  `json:"token"`
		PurchaserEmail  string  `json:"purchaser_email"`
	}

	Order struct {
		ID               int           `json:"id"`
		Status           string        `json:"status"`
		Title            string        `json:"title"`
		DoNotConvert     bool          `json:"do_not_convert"`
		OrderableType    string        `json:"orderable_type"`
		OrderableID      int           `json:"orderable_id"`
		PriceCurrency    string        `json:"price_currency"`
		PriceAmount      string        `json:"price_amount"`
		PayCurrency      string        `json:"pay_currency"`
		PayAmount        string        `json:"pay_amount"`
		LightningNetwork bool          `json:"lightning_network"`
		ReceiveCurrency  string        `json:"receive_currency"`
		ReceiveAmount    string        `json:"receive_amount"`
		CreatedAt        time.Time     `json:"created_at"`
		ExpireAt         time.Time     `json:"expire_at"`
		PaidAt           time.Time     `json:"paid_at"`
		PaymentAddress   string        `json:"payment_address"`
		OrderID          string        `json:"order_id"`
		PaymentURL       string        `json:"payment_url"`
		UnderpaidAmount  string        `json:"underpaid_amount"`
		OverpaidAmount   string        `json:"overpaid_amount"`
		IsRefundable     bool          `json:"is_refundable"`
		ConversionRate   string        `json:"conversion_rate"`
		Refunds          []Refund      `json:"refunds"`
		Voids            []interface{} `json:"voids"`
		Fees             []Fee         `json:"fees"`
	}

	CheckoutRequest struct {
		PayCurrency      string `json:"pay_currency"`
		LightningNetwork bool   `json:"lightning_network,omitempty"`
		PurchaserEmail   string `json:"purchaser_email,omitempty"`
		PlatformId       int32  `json:"platform_id,omitempty"`
	}

	Platform struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		IDName string `json:"id_name"`
	}

	CheckoutResponse struct {
		ID               int       `json:"id"`
		Status           string    `json:"status"`
		DoNotConvert     bool      `json:"do_not_convert"`
		PriceCurrency    string    `json:"price_currency"`
		PriceAmount      string    `json:"price_amount"`
		PayCurrency      string    `json:"pay_currency"`
		PayAmount        string    `json:"pay_amount"`
		LightningNetwork bool      `json:"lightning_network"`
		ReceiveCurrency  string    `json:"receive_currency"`
		ReceiveAmount    string    `json:"receive_amount"`
		CreatedAt        time.Time `json:"created_at"`
		ExpireAt         time.Time `json:"expire_at"`
		PaymentAddress   string    `json:"payment_address"`
		OrderID          string    `json:"order_id"`
		PaymentURL       string    `json:"payment_url"`
		UnderpaidAmount  string    `json:"underpaid_amount"`
		OverpaidAmount   string    `json:"overpaid_amount"`
		IsRefundable     bool      `json:"is_refundable"`
		Platform         Platform  `json:"platform"`
	}

	Refund struct {
		ID            int         `json:"id"`
		RequestAmount string      `json:"request_amount"`
		RefundAmount  string      `json:"refund_amount"`
		Address       string      `json:"address"`
		Status        string      `json:"status"`
		Memo          interface{} `json:"memo"`
		CreatedAt     time.Time   `json:"created_at"`
		Order         struct {
			ID int `json:"id"`
		} `json:"order"`
		RefundCurrency RefundCurrency `json:"refund_currency"`
		Transactions   []interface{}  `json:"transactions"`
		LedgerAccount  LedgerAccount  `json:"ledger_account"`
	}

	RefundCurrency struct {
		ID       int      `json:"id"`
		Title    string   `json:"title"`
		Symbol   string   `json:"symbol"`
		Platform Platform `json:"platform"`
	}

	LedgerAccount struct {
		ID       string   `json:"id"`
		Currency Currency `json:"currency"`
	}

	Currency struct {
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Symbol string `json:"symbol"`
	}

	Fee struct {
		Type     string   `json:"type"`
		Amount   string   `json:"amount"`
		Currency Currency `json:"currency"`
	}

	ListOrdersRequest struct {
		PerPage  int32  `json:"per_page"`
		Page     int32  `json:"page"`
		Sort     string `json:"sort"`
		CreateAt struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"create_at"`
	}

	Orders struct {
		CurrentPage int     `json:"current_page"`
		PerPage     int     `json:"per_page"`
		TotalOrders int     `json:"total_orders"`
		TotalPages  int     `json:"total_pages"`
		Orders      []Order `json:"orders"`
	}

	CallbackData struct {
		ID              int       `json:"id"`
		OrderID         string    `json:"order_id"`
		Status          string    `json:"status"`
		PriceAmount     string    `json:"price_amount"`
		PriceCurrency   string    `json:"price_currency"`
		ReceiveCurrency string    `json:"receive_currency"`
		ReceiveAmount   string    `json:"receive_amount"`
		PayAmount       string    `json:"pay_amount"`
		PayCurrency     string    `json:"pay_currency"`
		UnderpaidAmount string    `json:"underpaid_amount"`
		OverpaidAmount  string    `json:"overpaid_amount"`
		IsRefundable    bool      `json:"is_refundable"`
		CreatedAt       time.Time `json:"created_at"`
		Token           string    `json:"token"`
	}

	ErrorResponse struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Reason  string `json:"reason"`
	}
)
