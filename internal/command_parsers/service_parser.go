package command_parsers

import "protosym/internal/base_parser"

type ServiceParser struct {
	*base_parser.BaseParser
}

func NewServiceParser() *ServiceParser {
	return &ServiceParser{
		BaseParser: base_parser.NewBaseParser(&base_parser.ParserSettings{
			Command:        base_parser.ServiceCommand,
			CheckFirstOnly: true,
			NameDelimiter:  base_parser.BraceDelimiter,
			NameOffset:     1,
		}),
	}
}
