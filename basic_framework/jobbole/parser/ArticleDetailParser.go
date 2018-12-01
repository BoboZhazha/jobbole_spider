package parser

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"jobbole_spider/basic_framework/engine"
)

type DetailItem struct {
	Title    string
	Contents string
}

func ParseArticleDetail(reader io.Reader) engine.ParserResult {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		panic(err)
	}
	result := engine.ParserResult{}
	var items []interface{}
	title := doc.Find("div.entry-header").Find("h1").Text()
	text := doc.Find("div.entry").Text()
	detailItem := DetailItem{title, text}
	items = append(result.Items, detailItem)
	result.Items = items
	return result
}
