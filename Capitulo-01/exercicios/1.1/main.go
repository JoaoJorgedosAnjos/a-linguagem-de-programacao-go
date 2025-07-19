package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
	//go run main.go ola teste
	//main go ola teste
	//fmt.Println(strings.Join(os.Args[0:], "."))
	//go run main.go ola teste
	///tmp/go-build3362075384/b001/exe/main.ola.teste
}
