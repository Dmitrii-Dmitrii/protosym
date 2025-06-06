package command_parsers

import "protosym/internal/base_parser"

type EnumParser struct {
	*base_parser.BaseParser
}

func NewEnumParser() *EnumParser {
	return &EnumParser{
		BaseParser: base_parser.NewBaseParser(&base_parser.ParserSettings{
			Command:        base_parser.EnumCommand,
			CheckFirstOnly: true,
			NameDelimiter:  base_parser.BraceDelimiter,
			NameOffset:     1,
		}),
	}
}
