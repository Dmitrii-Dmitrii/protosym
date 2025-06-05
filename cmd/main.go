package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"protosym/internal/command_parsers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Wrong number of arguments.")
		os.Exit(1)
	}

	protoPath := os.Args[1]
	file, err := os.Open(protoPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	importParser := command_parsers.NewImportParser()
	serviceParser := command_parsers.NewServiceParser()
	rpcParser := command_parsers.NewRpcParser()
	enumParser := command_parsers.NewEnumParser()
	messageParser := command_parsers.NewMessageParser()

	importParser.
		SetNext(serviceParser).
		SetNext(rpcParser).
		SetNext(enumParser).
		SetNext(messageParser)

	scanner := bufio.NewScanner(file)
	num := 1
	for scanner.Scan() {
		line := scanner.Text()
		result := importParser.Parse(line, num)
		if result != "" {
			fmt.Println(result)
		}
		num++
	}
}
