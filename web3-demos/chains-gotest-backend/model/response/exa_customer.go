package response

import "chains-gotest-backend/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
