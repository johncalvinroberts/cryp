package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:build
var embeddedFiles embed.FS

func GetHandler() http.Handler {
	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	fsys, err := fs.Sub(embeddedFiles, "build")
	if err != nil {
		panic(err)
	}

	fs := http.FS(fsys)
	return http.FileServer(fs)
}
