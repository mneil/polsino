package request

import "io/ioutil"

type (
	// File contains a file as part of an HTTP request or response
	File struct {
		ContentType string
		Content     []byte
		Name        string
	}
)

func (f *File) save() error {
	return ioutil.WriteFile(f.Name, f.Content, 0600)
}
