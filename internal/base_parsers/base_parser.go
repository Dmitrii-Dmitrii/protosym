package base_parsers

import (
	"slices"
	"strconv"
	"strings"
)

type BaseParser struct {
	nextParser Parser
	settings   ParserSettings
}

func NewBaseParser(settings ParserSettings) *BaseParser {
	return &BaseParser{settings: settings}
}

func (p *BaseParser) SetNext(parser Parser) Parser {
	p.nextParser = parser
	return parser
}

func (p *BaseParser) Parse(line string, num int) string {
	if result := p.tryParse(line, num); result != "" {
		return result
	}

	if p.nextParser == nil {
		return ""
	}

	return p.nextParser.Parse(line, num)
}

func (p *BaseParser) tryParse(line string, num int) string {
	command := p.settings.Command
	symbols := strings.Split(line, " ")

	var hasCommand bool
	if p.settings.CheckFirstOnly {
		hasCommand = len(symbols) > 0 && symbols[0] == command
	} else {
		hasCommand = slices.Contains(symbols, command)
	}

	if !hasCommand {
		return ""
	}

	nameSymbols := symbols[slices.Index(symbols, command)+1]

	var name string
	var startIdx, endIdx int
	if p.settings.NameDelimiter == "\"" {
		name = strings.Split(nameSymbols, "\"")[1]
		startIdx = strings.Index(line, name)
		endIdx = startIdx + len(name) + p.settings.NameOffset
	} else {
		name = strings.Split(nameSymbols, p.settings.NameDelimiter)[0]
		startIdx = strings.Index(line, name) + p.settings.NameOffset
		endIdx = startIdx + len(name)
	}

	commandAlias := p.settings.Command
	if p.settings.CommandAlias != "" {
		commandAlias = p.settings.CommandAlias
	}

	result := name + " " +
		commandAlias + " " +
		strconv.Itoa(num) + ":" +
		strconv.Itoa(startIdx) + "-" +
		strconv.Itoa(endIdx)

	return result
}
