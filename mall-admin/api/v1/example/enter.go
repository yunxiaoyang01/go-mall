package example

import "github.com/yunxiaoyang01/go-mall/mall-admin/service"

type ApiGroup struct {
	ExcelApi
	CustomerApi
	FileUploadAndDownloadApi
}

var (
	excelService                 = service.ServiceGroupApp.ExampleServiceGroup.ExcelService
	customerService              = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
