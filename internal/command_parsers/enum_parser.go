package command_parsers

import "protosym/internal/base_parsers"

type EnumParser struct {
	*base_parsers.BaseParser
}

func NewEnumParser() *EnumParser {
	return &EnumParser{
		BaseParser: base_parsers.NewBaseParser(base_parsers.ParserSettings{
			Command:        "enum",
			CheckFirstOnly: true,
			NameDelimiter:  "{",
			NameOffset:     1,
		}),
	}
}
