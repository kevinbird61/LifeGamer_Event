package main

import (
	"fmt"
	"flag"
	"strings"
	"../internal/parser"
	"../internal/logger"
	"../internal/event/engine"
)

func main(){
	// argparse options
	itype := flag.String("itype", "json", "Specify type of input file.")
	ifile := flag.String("ifile", "test.json", "Specify filename of input file.")
	debug := flag.Bool("debug", false, "Debug flag, default is false.")

	// argparse parsing 
	flag.Parse()

	// create parser
	p := parser.Parser{}
	// create logger & initial 
	lg := logger.LG_Logger{}
	lg.Init()

	// Test 
	lg.Info.Println("Welcome to use LifeGamer event engine!")

	// switch case - base on file type
	switch t := strings.ToUpper(*itype); t {
		case "JSON":
			p.ReadJSON(*ifile)
		case "YAML":
			p.ReadYAML(*ifile)
		default:
			fmt.Println("%v - Not support yet!", *itype)
	}

	// Print for debug
	if *debug {
		fmt.Println(p.Obj)
	}

	// get something useful from source file
	fmt.Println("Poisson: ",engine.Debug_poisson())
}