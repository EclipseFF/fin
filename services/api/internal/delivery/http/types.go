package http

type GroupLoginRequestBody struct {
	Code string
}

type VideoAddRequestBody struct {
	Name    string
	Code    string
	Content []byte
}

type CreateGroupRequestBody struct {
	Code string
}

type RemoveVideoBody struct {
	Id string
}
