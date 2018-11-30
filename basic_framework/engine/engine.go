package engine

import (
	"jobbole_spider/basic_framework/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	// 只要有requests就一直做
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		reader, err := fetcher.Fetcher(r.Url)

		// 这里成千上万个requests出错了很正常, 记录日志,下一个
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		parserResult := r.ParserFunc(reader)
		// ... 就是把这个切片展开一个个加进去
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Fatalf("got item %v", item)
		}
	}
}
