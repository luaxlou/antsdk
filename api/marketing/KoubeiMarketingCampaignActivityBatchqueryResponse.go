package marketing

import (
  "github.com/luaxlou/antsdk/api"
)

type KoubeiMarketingCampaignActivityBatchqueryResponse struct {
  api.AlipayResponse
  TotalNumber int           `json:"total_number,string"`  // 总数量
  CampSets    []CampBaseDto `json:"camp_sets"`            // 活动列表
}
