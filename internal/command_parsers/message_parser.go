package command_parsers

import "protosym/internal/base_parser"

type MessageParser struct {
	*base_parser.BaseParser
}

func NewMessageParser() *MessageParser {
	return &MessageParser{
		BaseParser: base_parser.NewBaseParser(&base_parser.ParserSettings{
			Command:        base_parser.MessageCommand,
			CheckFirstOnly: true,
			NameDelimiter:  base_parser.BraceDelimiter,
			NameOffset:     1,
		}),
	}
}
