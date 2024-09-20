package httpclient

import (
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	var data []byte

	res, err := http.Get(url)
	if err != nil {
		return data, err
	}

	data, err = io.ReadAll(res.Body)
	res.Body.Close()

	return data, err
}
