# trades
U本位合約的交易功能

## example code
使用範例

### 實體化websocket物件
```
var ws *trades.Websocket = trades.NewWebsocket("{websocket url}")
```

### 實體化account物件
需要先在Binance設定中設定api key，並打開合約功能才能使用
```
var account binancetypes.Account = binancetypes.Account{
    APIKey:"{binance api key}",
    APISecret:"{binance api secret}",
}
```

### 下訂訂單

```
var (
    order trades.PlaceOrderFormBinance = trades.PlaceOrderFormBinance{}
    err error
    createdOrder trades.PlacedOrder
)

order.NewClientOrderID = uuid.NewString()
order.Symbol = "NOTUSDT"
order.Quantity = decimal.NewFromFloat32(500)
order.SetSideBUY()
order.RecvWindow = 5000
order.Timestamp = time.Now().UnixMilli()
order.SetPositionSideBOTH()
order.SetTypeLIMIT()
order.SetTimeInForceGTC()
order.Price = decimal.NewFromFloat32(0.01)


if createdOrder,err = ws.CreateOrder(ctx,account,&order) ; err != nil{
    panic(err)
}
```

### 取消未成交訂單
```
var(
    orderCancel trades.CancelOrderFormBinance = trades.CancelOrderFormBinance{}
    canceledOrder trades.CanceledOrder
    err error
)

orderCancel.Symbol = "NOTUSDT"
orderCancel.OrderID = {order id from created order}
orderCancel.OriginClientOrderID = "{order new client id from created order}"
orderCancel.Timestamp = time.Now().UnixMilli()
if canceledOrder,err = ws.CancelOrder(ctx, account, &orderCancel) ; err != nil{
    panic(err)
}

```