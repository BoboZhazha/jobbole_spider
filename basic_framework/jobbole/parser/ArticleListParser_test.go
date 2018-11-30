package parser

import (
	"fmt"
	"net/http"
	"testing"
)

func TestParserArticleList(t *testing.T) {
	url := "http://blog.jobbole.com/all-posts"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)

	}

}
