package parsers

import (
	"slices"
	"strconv"
	"strings"
)

type MessageParser struct {
	ProtoParser
}

func (p *MessageParser) Parse(line string, num int) string {
	command := "message"
	symbols := strings.Split(line, " ")
	if symbols[0] == command {
		nameSymbols := symbols[slices.Index(symbols, command)+1]
		name := strings.Split(nameSymbols, "{")[0]
		startIdx := strings.Index(line, name) + 1
		endIdx := startIdx + len(name)

		result := name + " " +
			command + " " +
			strconv.Itoa(num) + ":" +
			strconv.Itoa(startIdx) + "-" +
			strconv.Itoa(endIdx)

		return result
	}

	return p.ProtoParser.Parse(line, num)
}
