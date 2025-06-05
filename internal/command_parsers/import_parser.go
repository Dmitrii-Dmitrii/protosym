package command_parsers

import "protosym/internal/base_parsers"

type ImportParser struct {
	*base_parsers.BaseParser
}

func NewImportParser() *ImportParser {
	return &ImportParser{
		BaseParser: base_parsers.NewBaseParser(base_parsers.ParserSettings{
			Command:        "import",
			CheckFirstOnly: true,
			NameDelimiter:  "\"",
			NameOffset:     2,
		}),
	}
}
