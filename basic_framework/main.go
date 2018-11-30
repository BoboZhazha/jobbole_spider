package main

import (
	"jobbole_spider/basic_framework/engine"
	"jobbole_spider/basic_framework/jobbole/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://blog.jobbole.com/all-posts",
		ParserFunc: parser.ParseArticleList,
	})

}
