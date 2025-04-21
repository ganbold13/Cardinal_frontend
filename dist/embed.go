package dist

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:*
var embeddedFiles embed.FS

// New returns the embedded http.FileSystem for the main frontend.
func New() http.FileSystem {
	fsys, err := fs.Sub(embeddedFiles, ".")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}
