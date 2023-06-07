package customers

type CreateCustomerRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Avatar    string `json:"avatar" binding:"required"`
}
type CustomerResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}
type CreateCustomerResponse struct {
	Status string `json:"status"`
	Data   CustomerResponse
}
type DeleteCustomerRequest struct {
	CustomerID uint `json:"customer_id" binding:"required"`
}
type DeleteCustomerResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type Response struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
	Data       []UserData  `json:"data"`
	Support    SupportInfo `json:"support"`
}

type UserData struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type SupportInfo struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}
