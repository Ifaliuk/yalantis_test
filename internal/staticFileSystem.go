package internal

import "net/http"

type StaticFileSystem struct {
	http.Dir
}

func (fs StaticFileSystem) Open(name string) (result http.File, err error) {
	f, err := fs.Dir.Open(name)
	if err != nil {
		return
	}
	fi, err := f.Stat()
	if err != nil {
		return
	}
	if fi.IsDir() {
		// Return a response that would have been if directory would not exist:
		return fs.Dir.Open("does-not-exist")
	}
	return f, nil
}
