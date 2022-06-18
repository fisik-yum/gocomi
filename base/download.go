package base

import (
	"io"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string) error {
	//blatantly stolen from https://golangcode.com/download-a-file-from-a-url/
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
