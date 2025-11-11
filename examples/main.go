package main

import (
	"fmt"
	"os"
)

var examples = map[string]func(){
	"basic": basic,
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Bad Usage: you need to pass the example name.")
	}
	exampleFunc, ok := examples[os.Args[1]]
	if !ok {
		fmt.Println("Bad Usage: example not found.")
	}

	exampleFunc()
}
