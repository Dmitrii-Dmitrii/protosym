package parsers

import (
	"slices"
	"strconv"
	"strings"
)

type ImportParser struct {
	ProtoParser
}

func (p *ImportParser) Parse(line string, num int) string {
	command := "import"
	symbols := strings.Split(line, " ")
	if symbols[0] == command {
		nameSymbols := symbols[slices.Index(symbols, command)+1]
		name := nameSymbols[1 : len(nameSymbols)-2]
		startIdx := strings.Index(line, name)
		endIdx := startIdx + len(name) + 2

		result := name + " " +
			command + " " +
			strconv.Itoa(num) + ":" +
			strconv.Itoa(startIdx) + "-" +
			strconv.Itoa(endIdx)

		return result
	}

	return p.ProtoParser.Parse(line, num)
}
