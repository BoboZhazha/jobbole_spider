package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic"
	"net/http"
	"regexp"
	"strconv"
)

type Item struct {
	Title    string
	Url      string
	ImageUrl string
	Date     string
	Content  string
}

func main() {

	// 此处应该是从网页上拿到
	page_num := 563

	for i := 1; i <= page_num; i++ {
		url := "http://blog.jobbole.com/all-posts/page/" + strconv.FormatInt(int64(i), 10)
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.StatusCode)

		}
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		doc.Find("#archive").Find(".post.floated-thumb").Each(func(i int, selection *goquery.Selection) {

			title := selection.Find("a.archive-title").Text()
			href, _ := selection.Find("a.archive-title").Attr("href")
			imageUrl, _ := selection.Find("img").Attr("src")
			dateTemp := selection.Find("p").Text()
			dateRe := regexp.MustCompile(`\d{4}/\d{2}/\d{2}`)
			date := dateRe.FindString(dateTemp)
			content := selection.Find("span.excerpt").Text()

			item := Item{Title: title, Url: href, ImageUrl: imageUrl, Date: date, Content: content}
			saveItem(item)
		})
	}

}

func saveItem(item interface{}) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	client.Index().Index("dating_profile").Type("jobbole").BodyJson(item).Do(context.Background())
}
