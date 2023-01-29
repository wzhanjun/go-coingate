## go-coingate

[https://developer.coingate.com/reference/cryptocurrency-payment-api](https://developer.coingate.com/reference/cryptocurrency-payment-api)

```golang

client, _ = NewClient(1, "Api Token", WithTimeout(5*time.Second), WithSandBox(true))

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
    log.Fatalln(err)
}

log.Println(order)

```