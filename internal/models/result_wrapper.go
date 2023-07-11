package models

import codes "github.com/LucasToledoPereira/go-boilerplate/internal/enums/codes"

type ResultWrapper[T any] struct {
	Code    codes.Codes `json:"code"`
	Success bool        `json:"success"`
	Errors  []string    `json:"error"`
	Data    T           `json:"data"`
}

func NewResultWrapper[T any](c codes.Codes, s bool, e []string, d T) (rw *ResultWrapper[T]) {
	return &ResultWrapper[T]{
		Code:    c,
		Success: s,
		Errors:  e,
		Data:    d,
	}
}
