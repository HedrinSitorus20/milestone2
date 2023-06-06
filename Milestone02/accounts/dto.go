package accounts

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type ActorResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	RoleID   uint   `json:"role_id" `
	Verified string `json:"verified" gorm:"type:enum('true','false');default:'false'"`
	Active   string `json:"active" gorm:"type:enum('true','false');default:'false'"`
}

type LoginResponse struct {
	User  ActorResponse
	Token string `json:"token"`
}
