package trade

import (
  "github.com/luaxlou/antsdk/api"
)

type AlipayTradePreCreateResponse struct {
  api.AlipayResponse
  OutTradeNo  string  `json:"out_trade_no"` // 商户的订单号
  QRCode      string  `json:"qr_code"`      // 当前预下单请求生成的二维码码串，可以用二维码生成工具根据该码串值生成对应的二维码
}
