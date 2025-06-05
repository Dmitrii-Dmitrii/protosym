package base_parsers

import (
	"testing"
)

func TestBaseParser_Parse_NoMatch_NoNext(t *testing.T) {
	settings := ParserSettings{
		Command:        "enum",
		CheckFirstOnly: true,
		NameDelimiter:  "{",
		NameOffset:     1,
	}
	parser := NewBaseParser(settings)

	result := parser.Parse("message TestMessage {", 1)
	if result != "" {
		t.Errorf("Expected empty result when no match and no next parser, got '%s'", result)
	}
}

func TestBaseParser_Parse_NoMatch_WithNext(t *testing.T) {
	parser1 := NewBaseParser(ParserSettings{
		Command:        "enum",
		CheckFirstOnly: true,
		NameDelimiter:  "{",
		NameOffset:     1,
	})
	parser2 := NewBaseParser(ParserSettings{
		Command:        "message",
		CheckFirstOnly: true,
		NameDelimiter:  "{",
		NameOffset:     1,
	})

	parser1.SetNext(parser2)

	result := parser1.Parse("message TestMessage {", 1)
	expected := "TestMessage message 1:9-20"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestBaseParser_Parse_WithMatch_CheckFirstOnly(t *testing.T) {
	settings := ParserSettings{
		Command:        "enum",
		CheckFirstOnly: true,
		NameDelimiter:  "{",
		NameOffset:     1,
	}
	parser := NewBaseParser(settings)

	result := parser.Parse("enum TestEnum {", 1)
	expected := "TestEnum enum 1:6-14"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestBaseParser_Parse_WithMatch_NotCheckFirstOnly(t *testing.T) {
	settings := ParserSettings{
		Command:        "rpc",
		CommandAlias:   "method",
		CheckFirstOnly: false,
		NameDelimiter:  "(",
		NameOffset:     1,
	}
	parser := NewBaseParser(settings)

	result := parser.Parse("  rpc GetUser(", 1)
	expected := "GetUser method 1:7-14"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestBaseParser_Parse_WithQuoteDelimiter(t *testing.T) {
	settings := ParserSettings{
		Command:        "import",
		CheckFirstOnly: true,
		NameDelimiter:  "\"",
		NameOffset:     2,
	}
	parser := NewBaseParser(settings)

	result := parser.Parse("import \"google/protobuf/timestamp.proto\";", 1)
	expected := "google/protobuf/timestamp.proto import 1:8-41"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestBaseParser_Parse_CommandNotFound(t *testing.T) {
	settings := ParserSettings{
		Command:        "service",
		CheckFirstOnly: true,
		NameDelimiter:  "{",
		NameOffset:     1,
	}
	parser := NewBaseParser(settings)

	result := parser.Parse("message TestMessage {", 1)
	if result != "" {
		t.Errorf("Expected empty result when command not found, got '%s'", result)
	}
}

func TestBaseParser_Parse_EmptyLine(t *testing.T) {
	settings := ParserSettings{
		Command:        "message",
		CheckFirstOnly: true,
		NameDelimiter:  "{",
		NameOffset:     1,
	}
	parser := NewBaseParser(settings)

	result := parser.Parse("", 1)
	if result != "" {
		t.Errorf("Expected empty result for empty line, got '%s'", result)
	}
}

func TestParserChain(t *testing.T) {
	enumParser := NewBaseParser(ParserSettings{
		Command:        "enum",
		CheckFirstOnly: true,
		NameDelimiter:  "{",
		NameOffset:     1,
	})

	messageParser := NewBaseParser(ParserSettings{
		Command:        "message",
		CheckFirstOnly: true,
		NameDelimiter:  "{",
		NameOffset:     1,
	})

	serviceParser := NewBaseParser(ParserSettings{
		Command:        "service",
		CheckFirstOnly: true,
		NameDelimiter:  "{",
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
		result := enumParser.Parse(test.input, 1)
		if result != test.expected {
			t.Errorf("For input '%s': expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}
