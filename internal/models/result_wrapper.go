package models

type ResultWrapper[T any] struct {
	Code    string   `json:"code"`
	Success bool     `json:"success"`
	Errors  []string `json:"error"`
	Data    T        `json:"data"`
}

func NewResultWrapper[T any](c string, s bool, e []string, d T) (rw *ResultWrapper[T]) {
	return &ResultWrapper[T]{
		Code:    c,
		Success: s,
		Errors:  e,
		Data:    d,
	}
}
