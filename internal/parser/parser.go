package parser

import (
	_ "bufio"
	_ "fmt"
	_ "io"
	"io/ioutil"
	_ "os"

	"encoding/json"
	"gopkg.in/yaml.v2"
)

// Structure definition
type Object struct {
	Name 	string	`json: "name"`
	Model 	string 	`json: "model"`
	Lambda 	int		`json: "lambda"`
	X		int 	`json: "X"`
}

type Event struct {
	Event []Object
}

// Class: Parser
type Parser struct {
	Raw string
	Obj []Object
}

// =============================================== Member function of Parser ===============================================

// Provide function usage - load from file and read json
func (p *Parser) ReadJSON(url string) {
	// read file into string
	data, err := ioutil.ReadFile(url)
	if err != nil {
		panic(err)
	}
	// fmt.Print(string(data))
	p.Raw = string(data)
	json.Unmarshal(data, &p.Obj)

	// fmt.Println(p.obj[0])
}

// Provide function usage - load from file and read yaml
func (p *Parser) ReadYAML(url string) {
	// read file into string
	data, err := ioutil.ReadFile(url)
	if err != nil {
		panic(err)
	}
	
	p.Raw = string(data)
	yaml.Unmarshal(data, &p.Obj)
	// fmt.Println(p.obj[0])
}