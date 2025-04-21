package dist

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:*
var embeddedFiles embed.FS

func New() http.FileSystem {
	fsys, err := fs.Sub(embeddedFiles, ".")
	if err != nil {
		panic(err)
	}
	return WithSPA(http.FS(fsys))
}

