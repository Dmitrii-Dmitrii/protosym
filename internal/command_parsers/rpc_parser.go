package command_parsers

import "protosym/internal/base_parsers"

type RpcParser struct {
	*base_parsers.BaseParser
}

func NewRpcParser() *RpcParser {
	return &RpcParser{
		BaseParser: base_parsers.NewBaseParser(base_parsers.ParserSettings{
			Command:        "rpc",
			CommandAlias:   "method",
			CheckFirstOnly: false,
			NameDelimiter:  "(",
			NameOffset:     1,
		}),
	}
}
