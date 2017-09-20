package httpget

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(srcFileUrl, dstFilePath string) (err error) {
	var rc io.ReadCloser
	if rc, err = OpenRemoteFile(srcFileUrl); err == nil {
		defer rc.Close()
		SaveToFile(rc, dstFilePath)
	}
	return
}

func OpenRemoteFile(srcFileUrl string) (src io.ReadCloser, err error) {
	var resp *http.Response
	if resp, err = new(http.Client).Get(srcFileUrl); (err == nil) && (resp != nil) {
		src = resp.Body
	}
	return
}

func SaveToFile(src io.Reader, dstFilePath string) (err error) {
	var file *os.File
	if file, err = os.Create(dstFilePath); file != nil {
		defer file.Close()
		if err == nil {
			_, err = io.Copy(file, src)
		}
	}
	return
}
