package main

import (
	"fmt"
	// "log"
	"net/http"
	"flag"
	"strings"
	"encoding/json"
	"../internal/parser"
	"../internal/logger"
	"../internal/history"
	"../internal/runtime"
	"../internal/event/engine"
)

func main(){
	// argparse options
	itype := flag.String("itype", "json", "Specify type of input file.")
	ifile := flag.String("ifile", "test.json", "Specify filename of input file.")
	debug := flag.Bool("debug", false, "Debug flag, default is false.")
	runtime := flag.Bool("runtime", false, "Runtime server enable/disable, default is disable.")
	port := flag.Int("port", 9000, "Port number for runtime server.")
	doom := flag.Float64("doom", 0, "Doom of this simulation.")

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
	// create engine & initialize it
	engine := engine.Engine{}
	engine.Init(0, *doom, p.Obj)

	// start simulation process
	engine.Start()

	// assign history event list into engine
	history := history.History{}
	history.Init(0, *doom, engine.History)

	// Print out total event summation 
	fmt.Println(history.Map)

	// ================================================= 
	/* 
	 * Runtime control 
	 * - if not specify runtime flag, then event engine will not support runtime controll
	 * - only support output "history"
	*/
	// =================================================
	if *runtime {
		// create monitor, aim for peaker program/runtime control 
		app := monitor.CreateMonitor()

		// Handler
		/*
			/health: 	[GET] health check
			/fetch:		[GET] get the history
			/insert:	[POST] insert new event into history
		*/
		app.Handle(`^/health$`, func(ctx *monitor.Context){
			// fmt.Println(ctx.Request.Method)
			ctx.Text(http.StatusOK, "OK")
		})
		app.Handle(`^/fetch$`, func(ctx *monitor.Context){
			// return all 
			jsonstr,_ := json.Marshal(history)
			ctx.Text(http.StatusOK, fmt.Sprintf("%s", string(jsonstr)))
		})
		app.Handle(`^/insert`, func(ctx *monitor.Context){
			// insert new event 
			if ctx.Request.Method == http.MethodPost {
				// TODO
				// schedule new event into history
			} else {
				ctx.Text(http.StatusMethodNotAllowed, "Please using POST method")
			}
		})

		fmt.Println("Runtime server starting ...")
		// Running server in the background
		go func() {
			http.ListenAndServe(fmt.Sprintf(":%d", *port), app)
		}()
	}

	/*
		Dealing with other things
	*/

	// need an endless loop to maintain background server 
	for {}
}
