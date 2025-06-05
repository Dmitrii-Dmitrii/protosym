package parsers

type ProtoParser struct {
	nextParser Parser
}

func (p *ProtoParser) SetNext(parser Parser) Parser {
	p.nextParser = parser
	return parser
}

func (p *ProtoParser) Parse(line string, num int) string {
	if p.nextParser == nil {
		return ""
	}

	return p.nextParser.Parse(line, num)
}
