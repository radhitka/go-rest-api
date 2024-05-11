package utils

import "net/http"

type ResponseData struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}

func NewResponseData() *ResponseData {
	return &ResponseData{}
}

func (r *ResponseData) SuccessOk() {
	r.StatusCode = http.StatusOK
	r.Message = http.StatusText(http.StatusOK)
}

func (r *ResponseData) SuccessCreated() {
	r.StatusCode = http.StatusCreated
	r.Message = http.StatusText(http.StatusCreated)
}

func (r *ResponseData) InternalServerError() {
	r.StatusCode = http.StatusInternalServerError
	r.Message = http.StatusText(http.StatusInternalServerError)
}

func (r *ResponseData) BadRequest() {
	r.StatusCode = http.StatusBadRequest
	r.Message = http.StatusText(http.StatusBadRequest)
}

func (r *ResponseData) WithData(data any) *ResponseData {
	r.Data = data
	return r
}

func (r *ResponseData) WithMessage(message string) *ResponseData {
	r.Message = message
	return r
}
