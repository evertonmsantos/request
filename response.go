package request

import "net/http"

type Response struct {
	Status        string
	StatusCode    int
	ContentLength int64
	Body          string
	Header        http.Header
}
