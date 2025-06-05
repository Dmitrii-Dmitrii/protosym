package command_parsers

import "protosym/internal/base_parsers"

type MessageParser struct {
	*base_parsers.BaseParser
}

func NewMessageParser() *MessageParser {
	return &MessageParser{
		BaseParser: base_parsers.NewBaseParser(base_parsers.ParserSettings{
			Command:        "message",
			CheckFirstOnly: true,
			NameDelimiter:  "{",
			NameOffset:     1,
		}),
	}
}
