package main

import (
	"jobbole_spider/basic_framework/engine"
	"jobbole_spider/basic_framework/jobbole/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://blog.jobbole.com/all-posts",
		ParserFunc: parser.ParseArticleList,
	}, engine.Request{
		Url:        "http://blog.jobbole.com/all-posts/page/2",
		ParserFunc: parser.ParseArticleList,
	}, engine.Request{
		Url:        "http://blog.jobbole.com/all-posts/page/3",
		ParserFunc: parser.ParseArticleList,
	},
	)

}
