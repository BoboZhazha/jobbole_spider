package parser

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"jobbole_spider/basic_framework/engine"
	"regexp"
)

type Item struct {
	Title    string
	Url      string
	ImageUrl string
	Date     string
	Content  string
}

func ParseArticleList(reader io.Reader) engine.ParserResult {

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		panic(err)
	}
	result := engine.ParserResult{}

	doc.Find("#archive").Find(".post.floated-thumb").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find("a.archive-title").Text()
		href, _ := selection.Find("a.archive-title").Attr("href")
		imageUrl, _ := selection.Find("img").Attr("src")
		dateTemp := selection.Find("p").Text()
		dateRe := regexp.MustCompile(`\d{4}/\d{2}/\d{2}`)
		date := dateRe.FindString(dateTemp)
		content := selection.Find("span.excerpt").Text()
		item := Item{Title: title, Url: href, ImageUrl: imageUrl, Date: date, Content: content}

		result.Items = append(result.Items, item)

		// 每个list页面解析出来多个detail页的Request
		result.Requests = append(result.Requests, engine.Request{
			Url:        href,
			ParserFunc: ParseArticleDetail,
		})
	})
	return result

}
