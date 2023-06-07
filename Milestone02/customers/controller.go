package customers

import "github.com/gin-gonic/gin"

type Controller struct {
	uc *UseCase
}

func (c Controller) CreateCustomer(cc *CreateCustomerRequest) (*CreateCustomerResponse, error) {
	cust := Customer{
		FirstName: cc.FirstName,
		LastName:  cc.LastName,
		Email:     cc.Email,
		Avatar:    cc.Avatar,
	}

	err := c.uc.CreateCustomer(&cust)

	if err != nil {
		return nil, err
	}

	res := &CreateCustomerResponse{
		Status: "sukses",
		Data: CustomerResponse{
			FirstName: cust.FirstName,
			LastName:  cust.LastName,
			Email:     cust.Email,
			Avatar:    cust.Avatar,
		},
	}

	return res, nil
}

func (c Controller) DeleteCustomer(d *DeleteCustomerRequest) (*DeleteCustomerResponse, error) {
	customer := Customer{}
	err := c.uc.DeleteCustomer(&customer, d)

	if err != nil {
		return nil, err
	}

	res := &DeleteCustomerResponse{
		Status:  "ok",
		Message: "customer berhasil dihapus",
	}
	return res, nil
}

func (c Controller) SearchCustomers(con *gin.Context) ([]Customer, error) {
	firstname := con.Query("first_name")
	email := con.Query("email")
	limit := con.DefaultQuery("limit", "5")
	page := con.DefaultQuery("page", "1")

	cust, err := c.uc.GetCustomersByNameAndEmail(firstname, email, limit, page)
	if err != nil {
		return nil, err
	}

	return cust, nil
}

func NewController(uc *UseCase) *Controller {
	return &Controller{
		uc: uc,
	}
}
