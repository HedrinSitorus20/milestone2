package customers

type ICustomer interface {
	CreateCustomer(c *Customer) error
	DeleteCustomer(c *Customer, d *DeleteCustomerRequest) error
	GetCustomersByNameAndEmail(firstname string, email string, limit string, page string) ([]Customer, error)
}

type UseCase struct {
	repo *Repository
}

func (u UseCase) CreateCustomer(c *Customer) error {
	return u.repo.CreateCustomer(c)
}

func (u UseCase) DeleteCustomer(c *Customer, d *DeleteCustomerRequest) error {
	return u.repo.DeleteCustomer(c, d)
}

func (u UseCase) GetCustomersByNameAndEmail(firstname string, email string, limit string, page string) ([]Customer, error) {
	return u.repo.FindAllCustomers(firstname, email, limit, page)
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{repo: repo}
}
