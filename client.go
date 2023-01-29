package coingate

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

type ClientOption func(*Client)

func WithAppId(appId int) ClientOption {
	return func(c *Client) {
		c.AppId = appId
	}
}

func WithToken(token string) ClientOption {
	return func(c *Client) {
		c.Token = token
	}
}

func WithSandBox(isSandbox bool) ClientOption {
	return func(c *Client) {
		url := ApiLiveUrl
		if isSandbox {
			url = ApiSandBoxUrl
		}
		c.BaseUrl = url
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.Timeout = timeout
	}
}

// Create a coingate api client
func NewClient(appId int, token string, options ...ClientOption) (*Client, error) {
	if appId == 0 || token == "" {
		return nil, errors.New("appid and token are required to create a client")
	}

	client := &Client{
		AppId:   appId,
		Token:   token,
		BaseUrl: ApiLiveUrl,
		Timeout: time.Second * 15,
	}

	for _, option := range options {
		option(client)
	}

	return client, nil
}

func (c *Client) CreateOrder(d CreateOrderRequest) (*Order, error) {
	response, err := c.request(gorequest.POST, "/orders", d)
	if err != nil {
		return &Order{}, err
	}

	res := new(Order)
	err = json.Unmarshal([]byte(response), res)

	return res, err
}

func (c *Client) Checkout(orderId int, d CheckoutRequest) (*CheckoutResponse, error) {
	response, err := c.request(gorequest.POST, fmt.Sprintf("/orders/%d/checkout", orderId), d)
	if err != nil {
		return &CheckoutResponse{}, err
	}

	res := new(CheckoutResponse)
	err = json.Unmarshal([]byte(response), res)

	return res, err
}

func (c *Client) GetOrder(orderId int) (*Order, error) {
	response, err := c.request(gorequest.GET, fmt.Sprintf("/orders/%d", orderId), nil)
	if err != nil {
		return &Order{}, err
	}

	res := new(Order)
	err = json.Unmarshal([]byte(response), res)

	return res, err
}

func (c *Client) ListOrders(d ListOrdersRequest) (*Orders, error) {
	if d.PerPage <= 0 {
		d.PerPage = 10
	}

	if d.Page <= 1 {
		d.Page = 1
	}

	if len(d.Sort) == 0 {
		d.Sort = "created_at_desc"
	}

	response, err := c.request(gorequest.GET, "/orders", d)
	if err != nil {
		return &Orders{}, err
	}

	res := new(Orders)
	err = json.Unmarshal([]byte(response), res)

	return res, err
}

func (c *Client) ProcessCallback(r *http.Request) (*CallbackData, error) {
	err := r.ParseForm()
	if err != nil {
		return &CallbackData{}, err
	}
	data := new(CallbackData)

	data.ID, _ = strconv.Atoi(r.Form.Get("id"))
	data.OrderID = r.Form.Get("order_id")
	data.Status = r.Form.Get("status")
	data.PriceAmount = r.Form.Get("price_amount")
	data.PriceCurrency = r.Form.Get("price_currency")
	data.ReceiveCurrency = r.Form.Get("receive_currency")
	data.ReceiveAmount = r.Form.Get("receive_amount")
	data.PayAmount = r.Form.Get("pay_amount")
	data.PayCurrency = r.Form.Get("pay_currency")
	data.UnderpaidAmount = r.Form.Get("underpaid_amount")
	data.OverpaidAmount = r.Form.Get("overpaid_amount")
	data.IsRefundable, _ = strconv.ParseBool(r.Form.Get("is_refundable"))
	data.Token = r.Form.Get("token")
	data.CreatedAt, _ = time.Parse("2006-01-02T15:04:05-07:00", r.Form.Get("created_at"))

	return data, nil
}

func (c *Client) request(method string, uri string, payload interface{}) (string, error) {

	request := gorequest.New()
	url := fmt.Sprintf("%s/%s", c.BaseUrl, strings.Trim(uri, "/"))

	request.CustomMethod(method, url)

	if method == gorequest.POST {
		request.Type("urlencoded").Send(payload)
	} else if method == gorequest.GET {
		request.Query(payload)
	}

	request.Set("Accept", "application/json").
		Set("Content-Type", "application/x-www-form-urlencoded").
		Set("Authorization", "Bearer "+c.Token).
		Timeout(c.Timeout)

	resp, body, errs := request.End()

	if len(errs) > 0 {
		return "", errs[0]
	}

	if resp.StatusCode != 200 {
		var m ErrorResponse
		_ = json.Unmarshal([]byte(body), &m)

		return "", fmt.Errorf("Error %d: %s %s", resp.StatusCode, m.Reason, m.Message)
	}

	return body, nil
}
