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

func DefaultSuccessOK() *ResponseData {
	return &ResponseData{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
	}
}

func (r *ResponseData) SuccessOk() {
	r.StatusCode = http.StatusOK
	if r.Message == "" {
		r.Message = http.StatusText(http.StatusOK)
	}
}

func (r *ResponseData) SuccessCreated() {
	r.StatusCode = http.StatusCreated
	if r.Message == "" {
		r.Message = http.StatusText(http.StatusCreated)
	}
}

func (r *ResponseData) NotFound() {
	r.StatusCode = http.StatusNotFound
	if r.Message == "" {
		r.Message = http.StatusText(http.StatusNotFound)
	}
}

func (r *ResponseData) InternalServerError() {
	r.StatusCode = http.StatusInternalServerError
	if r.Message == "" {
		r.Message = http.StatusText(http.StatusInternalServerError)
	}
}

func (r *ResponseData) BadRequest() {
	r.StatusCode = http.StatusBadRequest
	if r.Message == "" {
		r.Message = http.StatusText(http.StatusBadRequest)
	}
}

func (r *ResponseData) WithData(data any) *ResponseData {
	r.Data = data
	return r
}

func (r *ResponseData) WithMessage(message string) *ResponseData {
	r.Message = message
	return r
}
