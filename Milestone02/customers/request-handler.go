package customers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type RequestHandler struct {
	ctrl *Controller
}

func (rh RequestHandler) CreateCustomer(c *gin.Context) {
	cust := CreateCustomerRequest{}
	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := rh.ctrl.CreateCustomer(&cust)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) DeleteCustomer(c *gin.Context) {
	dcr := DeleteCustomerRequest{}
	if err := c.ShouldBindJSON(&dcr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := rh.ctrl.DeleteCustomer(&dcr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) GetAllCustomers(c *gin.Context) {
	res, err := rh.ctrl.SearchCustomers(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func NewRequestHandler(ctrl *Controller) *RequestHandler {
	return &RequestHandler{ctrl: ctrl}
}

func DefaultRequestHandler(db *gorm.DB) *RequestHandler {
	return NewRequestHandler(
		NewController(
			NewUseCase(
				NewRepository(db),
			),
		),
	)
}
