package responses

type SuccessResponse struct {
	Data any `json:"data"`
}

type SuccessArrayResponse struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Data any `json:"data"`
}
