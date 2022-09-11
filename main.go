package main

import (
	"fmt"
	"os"

	"github.com/moritzkalwa/go-parse/Parser"
)

func main() {
	data, err := os.ReadFile("STOP_AREA.bin")
	if err != nil {
		fmt.Print(err)
	}
	data = data[162193:]
	result := Parser.ParseBytesFormat5(&data)
	for i := 0; i < 100; i++ {
		fmt.Printf("%+v\n", result[i])
	}
}
