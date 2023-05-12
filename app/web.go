package app

import (
	"embed"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"web-installer-backend/global"
)

type StaticResource struct {
	path string
	fs   embed.FS
}

func NewResource() *StaticResource {
	return &StaticResource{
		path: "dist",
		fs:   global.Asset,
	}
}
func (r *StaticResource) Open(name string) (fs.File, error) {

	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}

	fullName := filepath.Join(r.path, filepath.FromSlash(path.Clean("/assets/"+name)))

	file, err := r.fs.Open(fullName)
	if err != nil {
		fmt.Println("fs.Open ERR:", err)

	}
	_, err1 := os.Open(fullName)
	if err1 != nil {
		fmt.Println("ERR1:", err1)
	}

	if err != nil {
		fmt.Println("Open Err:", err)
	}

	return file, err
}

func Resource(c *gin.Context) {
	c.Header("content-type", "text/html;charset=utf-8")
	c.String(200, string(global.Html))
}
