package base_parser

import (
	"testing"
)

func TestBaseParser_Parse_NoMatch_NoNext(t *testing.T) {
	settings := &ParserSettings{
		Command:        EnumCommand,
		CheckFirstOnly: true,
		NameDelimiter:  BraceDelimiter,
		NameOffset:     1,
	}
	parser := NewBaseParser(settings)

	result, ok := parser.Parse("message TestMessage {", 1)
	if ok {
		t.Errorf("Expected empty result when no match and no next parser, got '%s'", result)
	}
}

func TestBaseParser_Parse_NoMatch_WithNext(t *testing.T) {
	parser1 := NewBaseParser(&ParserSettings{
		Command:        EnumCommand,
		CheckFirstOnly: true,
		NameDelimiter:  BraceDelimiter,
		NameOffset:     1,
	})
	parser2 := NewBaseParser(&ParserSettings{
		Command:        MessageCommand,
		CheckFirstOnly: true,
		NameDelimiter:  BraceDelimiter,
		NameOffset:     1,
	})

	parser1.SetNext(parser2)

	result, ok := parser1.Parse("message TestMessage {", 1)
	expected := "TestMessage message 1:9-20"
	if !ok && result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestBaseParser_Parse_WithMatch_CheckFirstOnly(t *testing.T) {
	settings := &ParserSettings{
		Command:        EnumCommand,
		CheckFirstOnly: true,
		NameDelimiter:  BraceDelimiter,
		NameOffset:     1,
	}
	parser := NewBaseParser(settings)

	result, ok := parser.Parse("enum TestEnum {", 1)
	expected := "TestEnum enum 1:6-14"
	if !ok && result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestBaseParser_Parse_WithMatch_NotCheckFirstOnly(t *testing.T) {
	settings := &ParserSettings{
		Command:        RpcCommand,
		CommandAlias:   RpcAlias,
		CheckFirstOnly: false,
		NameDelimiter:  ParenthesisDelimiter,
		NameOffset:     1,
	}
	parser := NewBaseParser(settings)

	result, ok := parser.Parse("  rpc GetUser(", 1)
	expected := "GetUser method 1:7-14"
	if !ok && result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestBaseParser_Parse_WithQuoteDelimiter(t *testing.T) {
	settings := &ParserSettings{
		Command:        ImportCommand,
		CheckFirstOnly: true,
		NameDelimiter:  QuoteDelimiter,
		NameOffset:     2,
	}
	parser := NewBaseParser(settings)

	result, ok := parser.Parse("import \"google/protobuf/timestamp.proto\";", 1)
	expected := "google/protobuf/timestamp.proto import 1:8-41"
	if !ok && result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestBaseParser_Parse_CommandNotFound(t *testing.T) {
	settings := &ParserSettings{
		Command:        ServiceCommand,
		CheckFirstOnly: true,
		NameDelimiter:  BraceDelimiter,
		NameOffset:     1,
	}
	parser := NewBaseParser(settings)

	result, ok := parser.Parse("message TestMessage {", 1)
	if ok {
		t.Errorf("Expected empty result when command not found, got '%s'", result)
	}
}

func TestBaseParser_Parse_EmptyLine(t *testing.T) {
	settings := &ParserSettings{
		Command:        MessageCommand,
		CheckFirstOnly: true,
		NameDelimiter:  BraceDelimiter,
		NameOffset:     1,
	}
	parser := NewBaseParser(settings)

	result, ok := parser.Parse("", 1)
	if ok {
		t.Errorf("Expected empty result for empty line, got '%s'", result)
	}
}

func TestParserChain(t *testing.T) {
	enumParser := NewBaseParser(&ParserSettings{
		Command:        EnumCommand,
		CheckFirstOnly: true,
		NameDelimiter:  BraceDelimiter,
		NameOffset:     1,
	})

	messageParser := NewBaseParser(&ParserSettings{
		Command:        MessageCommand,
		CheckFirstOnly: true,
		NameDelimiter:  BraceDelimiter,
		NameOffset:     1,
	})

	serviceParser := NewBaseParser(&ParserSettings{
		Command:        ServiceCommand,
		CheckFirstOnly: true,
		NameDelimiter:  BraceDelimiter,
		NameOffset:     1,
	})

	enumParser.SetNext(messageParser).SetNext(serviceParser)

	tests := []struct {
		input    string
		expected string
	}{
		{"enum Status {", "Status enum 1:6-12"},
		{"message User {", "User message 1:9-13"},
		{"service UserService {", "UserService service 1:9-20"},
		{"unknown command", ""},
	}

	for _, test := range tests {
		result, ok := enumParser.Parse(test.input, 1)
		if !ok && result != test.expected {
			t.Errorf("For input '%s': expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}
