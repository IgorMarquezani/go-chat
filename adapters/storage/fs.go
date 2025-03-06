package storage

import (
	"io"
	"os"
)

type FS struct{}

func (f FS) Create(key string, src io.Reader) error {
	file, err := os.Create("/var/gochat/" + key)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, src)

	return err
}

func (f FS) Delete(key string) error {
	return os.Remove("/var/gochat/" + key)
}
