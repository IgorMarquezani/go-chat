package ports

import "io"

type Storage interface {
	Create(string, io.Reader) error
	Delete(string) error
}
