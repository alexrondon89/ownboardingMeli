package dto

type HttpResponse struct {
	StatusCode	int
	Body		[]byte
}

func NewHttpResponse(body []byte, statusCode int) *HttpResponse {
	return &HttpResponse{Body: body, StatusCode: statusCode}
}