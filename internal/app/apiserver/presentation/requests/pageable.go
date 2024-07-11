package requests

type GinPageableRequest struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
