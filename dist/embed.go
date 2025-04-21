// File: dist/embed.go
package dist

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:*
var embeddedFiles embed.FS

var StaticFiles http.FileSystem

func init() {
	// Strip the prefix from dist/, so you can access /index.html, /assets, etc.
	fsys, err := fs.Sub(embeddedFiles, ".")
	if err != nil {
		panic(err)
	}
	StaticFiles = http.FS(fsys)
}
