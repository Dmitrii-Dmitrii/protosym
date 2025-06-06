package command_parsers

import "protosym/internal/base_parser"

type ImportParser struct {
	*base_parser.BaseParser
}

func NewImportParser() *ImportParser {
	return &ImportParser{
		BaseParser: base_parser.NewBaseParser(&base_parser.ParserSettings{
			Command:        base_parser.ImportCommand,
			CheckFirstOnly: true,
			NameDelimiter:  base_parser.QuoteDelimiter,
			NameOffset:     2,
		}),
	}
}
