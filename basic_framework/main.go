package main

import (
	"jobbole_spider/basic_framework/engine"
	"jobbole_spider/basic_framework/jobbole/parser"
	"strconv"
)

func main() {

	const baseUrl = "http://blog.jobbole.com/all-posts"
	var seeds []engine.Request

	for i := 1; i < 500; i++ {
		urlStr := baseUrl + "/page/" + strconv.FormatInt(int64(i), 10)
		request := engine.Request{
			Url:        urlStr,
			ParserFunc: parser.ParseArticleList,
		}
		seeds = append(seeds, request)
	}

	//engine.Run(engine.Request{
	//		Url:        "http://blog.jobbole.com/all-posts",
	//		ParserFunc: parser.ParseArticleList,
	//	}, engine.Request{
	//		Url:        "http://blog.jobbole.com/all-posts/page/2",
	//		ParserFunc: parser.ParseArticleList,
	//	}, engine.Request{
	//		Url:        "http://blog.jobbole.com/all-posts/page/3",
	//		ParserFunc: parser.ParseArticleList,
	//	},
	//)

	engine.Run(seeds...)

}
