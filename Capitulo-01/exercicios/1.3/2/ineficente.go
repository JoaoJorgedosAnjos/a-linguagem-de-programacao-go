package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now() // Marca o início
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	// Opcional: Imprimir a string para garantir que o trabalho foi feito
	// fmt.Println(s) 
	
	// Imprime o tempo que a tarefa levou
	fmt.Printf("Versão ineficiente demorou: %.5fs\n", time.Since(start).Seconds())
}