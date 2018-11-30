package fetcher

import (
	"fmt"
	"io"
	"net/http"
)

// 传入一个url, 返回一个请求得到的内容(返回一个reader是配合goquery),和错误, 你可以拿到错误做你想做的
func Fetcher(url string) (io.Reader, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	// 有异常返回nil,和错误
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code : %d", resp.StatusCode)
	}

	return resp.Body, err

}
