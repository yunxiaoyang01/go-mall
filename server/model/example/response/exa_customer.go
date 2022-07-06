package response

import "github.com/yunxiaoyang01/go-mall/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
