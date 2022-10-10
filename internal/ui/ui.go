package ui

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/static"
)

type embedFileSystem struct {
	http.FileSystem
}

//go:embed dist
var embeddedFiles embed.FS

func GetUIFileSystem() static.ServeFileSystem {
	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	return embedFolder(embeddedFiles, "dist")
}

// needed to fulfill the interface of gin-contrib/static
func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

// embed folder to FS
func embedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
