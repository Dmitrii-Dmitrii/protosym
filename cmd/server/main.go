package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"protosym/internal/parsers"
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

	protoParser := &parsers.ProtoParser{}
	importParser := &parsers.ImportParser{}
	serviceParser := &parsers.ServiceParser{}
	rpcParser := &parsers.RpcParser{}
	enumParser := &parsers.EnumParser{}
	messageParser := &parsers.MessageParser{}

	protoParser.
		SetNext(importParser).
		SetNext(serviceParser).
		SetNext(rpcParser).
		SetNext(enumParser).
		SetNext(messageParser)

	scanner := bufio.NewScanner(file)
	num := 1
	for scanner.Scan() {
		line := scanner.Text()
		result := protoParser.Parse(line, num)
		if result != "" {
			fmt.Println(result)
		}
		num++
	}
}
