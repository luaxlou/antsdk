package open

import (
  "github.com/luaxlou/antsdk/api"
)

type AlipayOpenPublicAccountCreateResponse struct {
  api.AlipayResponse
  AgreementId string `json:"agreement_id"` // 协议号，商户会员在支付宝服务窗账号中的唯一标识
}
