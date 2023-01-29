package coingate

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var client *Client

func init() {
	client, _ = NewClient(1, "API Token", WithTimeout(5*time.Second), WithSandBox(true))
}

func TestClient_CreateOrder(t *testing.T) {
	order, err := client.CreateOrder(CreateOrderRequest{
		OrderID:         fmt.Sprintf("%d", rand.Int31()),
		PriceAmount:     100.01,
		PriceCurrency:   "USD",
		ReceiveCurrency: "BTC",
		Title:           "test order",
		CancelURL:       "https://pay-sandbox.coingate.com/",
		SuccessURL:      "https://pay-sandbox.coingate.com/",
		CallbackURL:     "https://pay-sandbox.coingate.com/",
	})
	if err != nil {
		t.Fatalf("error on create order, err:%v", err)
	}
	t.Log(order)
}

func TestClient_Checkout(t *testing.T) {
	order, err := client.Checkout(58743, CheckoutRequest{
		PayCurrency: "BTC",
	})
	if err != nil {
		t.Fatalf("error on check order, err:%v", err)
	}
	t.Log(order)
}

func TestClient_GetOrder(t *testing.T) {
	order, err := client.GetOrder(58739)
	if err != nil {
		t.Fatalf("error on get order, err:%v", err)
	}
	t.Log(order)
}

func TestClient_ListOrders(t *testing.T) {
	order, err := client.ListOrders(ListOrdersRequest{
		PerPage: 10,
		Page:    1,
	})
	if err != nil {
		t.Fatalf("error on get order, err:%v", err)
	}
	t.Log(order)
}
