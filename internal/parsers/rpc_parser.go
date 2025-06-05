package parsers

import (
	"slices"
	"strconv"
	"strings"
)

type RpcParser struct {
	ProtoParser
}

func (p *RpcParser) Parse(line string, num int) string {
	command := "rpc"
	commandAlias := "method"
	symbols := strings.Split(line, " ")
	if slices.Contains(symbols, command) {
		nameSymbols := symbols[slices.Index(symbols, command)+1]
		name := strings.Split(nameSymbols, "(")[0]
		startIdx := strings.Index(line, name) + 1
		endIdx := startIdx + len(name)

		result := name + " " +
			commandAlias + " " +
			strconv.Itoa(num) + ":" +
			strconv.Itoa(startIdx) + "-" +
			strconv.Itoa(endIdx)

		return result
	}

	return p.ProtoParser.Parse(line, num)
}
