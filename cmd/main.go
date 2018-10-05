package main

import (
	"fmt"
	// "log"
	"net/http"
	//"io/ioutil"
	"flag"
	"strings"
	"strconv"
	"encoding/json"
	"../internal/data"
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
	e := engine.Engine{}
	e.Init(0, *doom, p.Obj)

	// start simulation process
	e.Start()

	// assign history event list into engine
	history := history.History{}
	history.Init(0, *doom, e.History)

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
			ctx.Text(http.StatusOK, "text/plain", "OK")
		})
		app.Handle(`^/fetch$`, func(ctx *monitor.Context){
			// return all 
			jsonstr,_ := json.Marshal(history)
			ctx.Text(http.StatusOK, "text/plain", fmt.Sprintf("%s", string(jsonstr)))
		})
		app.Handle(`^/insert`, func(ctx *monitor.Context){
			ctx.Request.ParseForm()
			// insert new event 
			if ctx.Request.Method == http.MethodPost {
				// TODO
				// schedule new event into history
				var NewEvent data.Object
				for key,value := range ctx.Request.Form {
					// fmt.Printf("%s = %s\n", key, value)
					switch key {
						case "event_name":
							NewEvent.Name = value[0]
						case "event_model":
							NewEvent.Model = value[0]
						case "lambda":
							NewEvent.Lambda,_ = strconv.Atoi(value[0])
						case "x": 
							NewEvent.X,_ = strconv.Atoi(value[0])
						default:
							ctx.Text(404, "text/plain", "Error input form.")
					}
				}

				fmt.Println(NewEvent)
				var event_arr []data.Object
				event_arr = append(event_arr, NewEvent)
				// And then regenerate Event 
				newEngine := engine.Engine{}
				newEngine.Init(0, *doom, event_arr)
				newEngine.Start()
				// Push new event into history
				history.Add_event_list(newEngine.History)

				// Print out total event summation 
				fmt.Println(history.Map)

			} else {
				ctx.Text(http.StatusMethodNotAllowed, "text/plain", "Please using POST method")
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
