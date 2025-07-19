package main

import (
	"fmt"
	//"os"
	//"strings"
	"time"
)

func main() {
	start := time.Now()
	
	//fmt.Println(strings.Join(os.Args[1:], " "))
	
	fmt.Printf("Versão eficiente demorou: %.5fs\n", time.Since(start).Seconds())

}
