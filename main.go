package main

import (
	"fmt"
	"os"
)

//  //  //

func main() {
	params := newParams()

	if err := params.parse(); err != nil {
		fmt.Printf("Error en los parámetros utilizados: %s\n", err)
		os.Exit(1)
	}
}
