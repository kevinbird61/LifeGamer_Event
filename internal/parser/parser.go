package parser

import (
	_ "bufio"
	_ "fmt"
	_ "io"
	"io/ioutil"
	_ "os"

	"encoding/json"
	"gopkg.in/yaml.v2"

	"../data"
)

// Structure definition
type Event struct {
	Event []data.Object
}

// Class: Parser
type Parser struct {
	Raw string
	Obj []data.Object
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