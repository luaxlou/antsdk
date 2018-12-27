package open

import (
  "github.com/luaxlou/antsdk/api"
)

type AlipayOpenPublicAccountQueryResponse struct {
  api.AlipayResponse
  PublicBindAccounts []StdPublicBindAccount `json:"public_bind_accounts"` // 绑定账户列表
}
