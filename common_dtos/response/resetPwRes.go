package response

type ResetPasswordResponse struct {
	RedirectUrl string `json:"redirect_url"`
	ErrorMsg    string `json:"error_msg"`
}
