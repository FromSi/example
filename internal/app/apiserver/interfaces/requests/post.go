package requests

type GinCreatePostRequestBody struct {
	Text string `from:"text" binding:"required"`
}

type GinShowPostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type GinUpdatePostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type GinUpdatePostRequestBody struct {
	Text *string `from:"text"`
}

type GinDeletePostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type GinResetPostRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}
