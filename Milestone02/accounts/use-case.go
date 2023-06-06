package accounts

type IAccount interface {
	Login(a *Actor, lr *LoginRequest) error
}
type UseCase struct {
	repo *Repository
}

func (u UseCase) Login(a *Actor, lr *LoginRequest) error {
	return u.repo.Login(a, lr)
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{repo: repo}
}
