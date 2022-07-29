package helper

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func CurlPostFile(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("token", os.Getenv("TOKEN_TO_SERVER"))
	return req, err
}

func CurlGetClient(uri string, params map[string]string) (*http.Request, error) {
	paramQuery := ""

	for key, value := range params {
		if paramQuery != "" {
			paramQuery = paramQuery + key + "=" + value + "&"
		} else {
			paramQuery = key + "=" + value + "&"
		}

	}

	req, err := http.NewRequest("GET", uri+"?"+paramQuery, nil)
	req.Header.Set("token", os.Getenv("TOKEN_TO_SERVER"))

	return req, err
}

func CurlPostClient(uri string, params map[string]string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("token", os.Getenv("TOKEN_TO_SERVER"))
	return req, err
}
