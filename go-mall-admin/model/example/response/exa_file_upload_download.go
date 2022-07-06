package response

import "github.com/yunxiaoyang01/go-mall/go-mall-admin/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
