package api

import (
  "github.com/luaxlou/antsdk/utils"
)

type IAlipayUploadRequest interface {
  IAlipayRequest
  GetFileParams() map[string]*utils.FileItem
}
