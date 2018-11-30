package engine

import "io"

type Request struct {
	Url        string
	ParserFunc func(reader io.Reader) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser(reader io.Reader) ParserResult {
	return ParserResult{}
}
