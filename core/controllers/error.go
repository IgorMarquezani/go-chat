package controllers

import "errors"

var (
	ErrBadRequest          = errors.New("bad request")
	ErrInternalServerError = errors.New("internal server error")
)

type Error struct {
	Message     string   `json:"error"`
	Description string   `json:"description"`
	Details     []string `json:"details"`
}

func (e Error) Error() string {
	return e.Message + ": " + e.Description
}
