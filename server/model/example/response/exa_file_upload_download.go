package response

import "github.com/yunxiaoyang01/go-mall/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
