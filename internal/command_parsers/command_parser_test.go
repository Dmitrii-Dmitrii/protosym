package command_parsers

import "testing"

func TestEnumParser_Parse_Valid(t *testing.T) {
	parser := NewEnumParser()

	tests := []struct {
		input    string
		lineNum  int
		expected string
	}{
		{"enum Status {", 1, "Status enum 1:6-12"},
		{"enum Color {", 5, "Color enum 5:6-11"},
	}

	for _, test := range tests {
		result := parser.Parse(test.input, test.lineNum)
		if result != test.expected {
			t.Errorf("For input '%s': expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestEnumParser_Parse_Invalid(t *testing.T) {
	parser := NewEnumParser()

	tests := []string{
		"message User {",
		"service UserService {",
		"rpc GetUser(",
		"import \"test.proto\";",
		"some enum Status {",
		" enum Priority {",
	}

	for _, input := range tests {
		result := parser.Parse(input, 1)
		if result != "" {
			t.Errorf("For input '%s': expected empty result, got '%s'", input, result)
		}
	}
}

func TestImportParser_Parse_Valid(t *testing.T) {
	parser := NewImportParser()

	tests := []struct {
		input    string
		lineNum  int
		expected string
	}{
		{"import \"google/protobuf/timestamp.proto\";", 1, "google/protobuf/timestamp.proto import 1:8-41"},
		{"import \"common/user.proto\";", 2, "common/user.proto import 2:8-27"},
	}

	for _, test := range tests {
		result := parser.Parse(test.input, test.lineNum)
		if result != test.expected {
			t.Errorf("For input '%s': expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestImportParser_Parse_Invalid(t *testing.T) {
	parser := NewImportParser()

	tests := []string{
		"message User {",
		"enum Status {",
		"service UserService {",
		"some import \"test.proto\";",
	}

	for _, input := range tests {
		result := parser.Parse(input, 1)
		if result != "" {
			t.Errorf("For input '%s': expected empty result, got '%s'", input, result)
		}
	}
}

func TestMessageParser_Parse_Valid(t *testing.T) {
	parser := NewMessageParser()

	tests := []struct {
		input    string
		lineNum  int
		expected string
	}{
		{"message User {", 1, "User message 1:9-13"},
		{"message CreateUserRequest {", 5, "CreateUserRequest message 5:9-26"},
	}

	for _, test := range tests {
		result := parser.Parse(test.input, test.lineNum)
		if result != test.expected {
			t.Errorf("For input '%s': expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestMessageParser_Parse_Invalid(t *testing.T) {
	parser := NewMessageParser()

	tests := []string{
		"enum Status {",
		"service UserService {",
		"rpc GetUser(",
		"import \"test.proto\";",
		"some message User {",
	}

	for _, input := range tests {
		result := parser.Parse(input, 1)
		if result != "" {
			t.Errorf("For input '%s': expected empty result, got '%s'", input, result)
		}
	}
}

func TestRpcParser_Parse_Valid(t *testing.T) {
	parser := NewRpcParser()

	tests := []struct {
		input    string
		lineNum  int
		expected string
	}{
		{"rpc GetUser(", 1, "GetUser method 1:5-12"},
		{"  rpc CreateUser(", 5, "CreateUser method 5:7-17"},
		{"    rpc DeleteUser(", 10, "DeleteUser method 10:9-19"},
		{"some other rpc UpdateUser(", 15, "UpdateUser method 15:16-26"},
	}

	for _, test := range tests {
		result := parser.Parse(test.input, test.lineNum)
		if result != test.expected {
			t.Errorf("For input '%s': expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestRpcParser_Parse_Invalid(t *testing.T) {
	parser := NewRpcParser()

	tests := []string{
		"enum Status {",
		"message User {",
		"service UserService {",
		"import \"test.proto\";",
		"method GetUser(",
	}

	for _, input := range tests {
		result := parser.Parse(input, 1)
		if result != "" {
			t.Errorf("For input '%s': expected empty result, got '%s'", input, result)
		}
	}
}

func TestServiceParser_Parse_Valid(t *testing.T) {
	parser := NewServiceParser()

	tests := []struct {
		input    string
		lineNum  int
		expected string
	}{
		{"service UserService {", 1, "UserService service 1:9-20"},
		{"service AuthService {", 5, "AuthService service 5:9-20"},
	}

	for _, test := range tests {
		result := parser.Parse(test.input, test.lineNum)
		if result != test.expected {
			t.Errorf("For input '%s': expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestServiceParser_Parse_Invalid(t *testing.T) {
	parser := NewServiceParser()

	tests := []string{
		"enum Status {",
		"message User {",
		"rpc GetUser(",
		"import \"test.proto\";",
		"some service UserService {",
	}

	for _, input := range tests {
		result := parser.Parse(input, 1)
		if result != "" {
			t.Errorf("For input '%s': expected empty result, got '%s'", input, result)
		}
	}
}

func TestCommandParserChain(t *testing.T) {
	enumParser := NewEnumParser()
	importParser := NewImportParser()
	messageParser := NewMessageParser()
	rpcParser := NewRpcParser()
	serviceParser := NewServiceParser()

	enumParser.SetNext(importParser).
		SetNext(messageParser).
		SetNext(rpcParser).
		SetNext(serviceParser)

	tests := []struct {
		input    string
		lineNum  int
		expected string
	}{
		{"enum Status {", 1, "Status enum 1:6-12"},
		{"import \"test.proto\";", 2, "test.proto import 2:8-20"},
		{"message User {", 3, "User message 3:9-13"},
		{"rpc GetUser(", 4, "GetUser method 4:5-12"},
		{"service UserService {", 5, "UserService service 5:9-20"},
		{"unknown command", 6, ""},
	}

	for _, test := range tests {
		result := enumParser.Parse(test.input, test.lineNum)
		if result != test.expected {
			t.Errorf("For input '%s': expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}
