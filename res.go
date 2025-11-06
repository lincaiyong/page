package page

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/log"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

//go:embed res.zip
var resZip []byte

func readZip(b []byte) (map[string][]byte, error) {
	ret := make(map[string][]byte)
	zipr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, err
	}

	for _, z := range zipr.File {
		rr, openErr := z.Open()
		if openErr != nil {
			return nil, openErr
		}

		b, readErr := io.ReadAll(rr)
		if readErr != nil {
			return nil, readErr
		}
		_ = rr.Close()
		if !strings.HasSuffix(z.Name, "/") {
			ret[z.Name] = b
		}
	}
	return ret, nil
}

var resFileMap map[string][]byte

func HandleRes(baseUrl string) gin.HandlerFunc {
	return func(c *gin.Context) {
		filePath := c.Param("filepath")[1:]
		if !strings.HasPrefix(filePath, "svg/") && !strings.HasPrefix(filePath, "vs/") {
			b, err := os.ReadFile(path.Join("res/", path.Base(filePath)))
			if err != nil {
				log.ErrorLog("fail to read index.js: %v", err)
				c.String(http.StatusNotFound, "file not found")
				return
			}
			content := strings.ReplaceAll(string(b), "<base_url>", baseUrl)
			if strings.HasSuffix(filePath, ".html") {
				c.Data(http.StatusOK, "text/html", []byte(content))
			} else {
				c.Data(http.StatusOK, "application/javascript", []byte(content))
			}
			return
		}
		if resFileMap == nil {
			resFileMap, _ = readZip(resZip)
		}
		b, ok := resFileMap[filePath]
		if !ok {
			c.String(http.StatusNotFound, "resource not found")
		}
		ext := path.Ext(filePath)
		contentType := "text/plain"
		if ext == ".css" {
			contentType = "text/css"
		} else if ext == ".js" {
			contentType = "application/javascript"
		} else if ext == ".svg" {
			contentType = "image/svg+xml"
		}
		c.Data(http.StatusOK, contentType, b)
	}
}
