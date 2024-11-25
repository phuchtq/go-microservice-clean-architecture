package request

type SignUpModel struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required, min=8"`
	RoleId   string `json:"role_id"`
}
