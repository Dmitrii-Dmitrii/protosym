package command_parsers

import "protosym/internal/base_parser"

type RpcParser struct {
	*base_parser.BaseParser
}

func NewRpcParser() *RpcParser {
	return &RpcParser{
		BaseParser: base_parser.NewBaseParser(&base_parser.ParserSettings{
			Command:        base_parser.RpcCommand,
			CommandAlias:   base_parser.RpcAlias,
			CheckFirstOnly: false,
			NameDelimiter:  base_parser.ParenthesisDelimiter,
			NameOffset:     1,
		}),
	}
}
