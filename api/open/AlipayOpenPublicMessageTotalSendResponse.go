package open

import (
  "github.com/gionna/antsdk/api"
)

type AlipayOpenPublicMessageTotalSendResponse struct {
  api.AlipayResponse
  MessageId string `json:"message_id"`  // 消息ID
}
