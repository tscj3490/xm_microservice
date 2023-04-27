package handlers

import "github.com/tscj3490/xm_microservice/entity"

// Bookings model response
// swagger:response companyResp
type swaggCreateCompanyResponse struct {
	// in:body
	Body entity.Company
}
