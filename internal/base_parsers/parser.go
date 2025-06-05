package base_parsers

type Parser interface {
	SetNext(next Parser) Parser
	Parse(line string, num int) string
}
