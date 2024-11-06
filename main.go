package main

import (
	"fmt"
	"github.com/karlmoad/AntlrParser/lib/parser"
)

func main() {
	fmt.Println("ANTLR Pl/SQL Parser...")
	ast, err := parser.ParseString("SELECT * FROM FAKE WHERE 1=1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ast)
}
