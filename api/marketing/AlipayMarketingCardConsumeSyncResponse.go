package marketing

import (
  "github.com/luaxlou/antsdk/api"
)

type AlipayMarketingCardConsumeSyncResponse struct {
  api.AlipayResponse
  ExternalCardNo string `json:"external_card_no"` // 外部卡号
}
