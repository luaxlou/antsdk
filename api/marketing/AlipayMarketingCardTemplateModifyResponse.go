package marketing

import (
  "github.com/gionna/antsdk/api"
)

type AlipayMarketingCardTemplateModifyResponse struct {
  api.AlipayResponse
  TemplateId string `json:"template_id"`  // 模板ID
}
