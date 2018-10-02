package main

import (
	"fmt"

	"../internal/parser"
)

func main(){
	// create parser
	p := parser.Parser{}
	p.ReadJSON("../test/test.json")

	fmt.Println(p.Obj)

	// get something useful from source file
}