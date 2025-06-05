package command_parsers

import "protosym/internal/base_parsers"

type ServiceParser struct {
	*base_parsers.BaseParser
}

func NewServiceParser() *ServiceParser {
	return &ServiceParser{
		BaseParser: base_parsers.NewBaseParser(base_parsers.ParserSettings{
			Command:        "service",
			CheckFirstOnly: true,
			NameDelimiter:  "{",
			NameOffset:     1,
		}),
	}
}
