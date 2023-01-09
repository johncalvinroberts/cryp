package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist
var embeddedFiles embed.FS

func GetHandler() http.Handler {
	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	fsys, err := fs.Sub(embeddedFiles, "dist")
	if err != nil {
		panic(err)
	}

	fs := http.FS(fsys)
	// return embedFolder(embeddedFiles, "dist")
	return http.FileServer(fs)
}
