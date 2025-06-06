package base_parser

import (
	"slices"
	"strconv"
	"strings"
)

type Parser interface {
	SetNext(next Parser) Parser
	Parse(line string, num int) (string, bool)
}

type BaseParser struct {
	nextParser Parser
	settings   *ParserSettings
}

func NewBaseParser(settings *ParserSettings) *BaseParser {
	return &BaseParser{settings: settings}
}

func (p *BaseParser) SetNext(parser Parser) Parser {
	p.nextParser = parser
	return parser
}

func (p *BaseParser) Parse(line string, num int) (string, bool) {
	if result, ok := p.tryParse(line, num); ok {
		return result, ok
	}

	if p.nextParser == nil {
		return "", false
	}

	return p.nextParser.Parse(line, num)
}

func (p *BaseParser) tryParse(line string, num int) (string, bool) {
	command := p.settings.Command
	symbols := strings.Split(line, " ")

	var hasCommand bool
	if p.settings.CheckFirstOnly {
		hasCommand = len(symbols) > 0 && symbols[0] == command
	} else {
		hasCommand = slices.Contains(symbols, command)
	}

	if !hasCommand {
		return "", false
	}

	nameIdx := slices.Index(symbols, command) + 1
	nameSymbols := symbols[nameIdx]

	var name string
	var startIdx, endIdx int
	nameSplit := strings.Split(nameSymbols, p.settings.NameDelimiter)

	if p.settings.NameDelimiter == "\"" {
		name = ""
		if len(nameSplit) > 1 {
			name = nameSplit[1]
		}

		startIdx = strings.Index(line, name)
		endIdx = startIdx + len(name) + p.settings.NameOffset
	} else {
		name = nameSplit[0]
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

	return result, true
}
