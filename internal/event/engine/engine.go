package engine

import (
	"fmt"
	"sort"
	"strings"

	"../../data"
	"../model"
)

// Class: Engine
type Engine struct {
	Queue		[]data.Event 
	History		[]data.Event 
	Timeline	float64
	Doom		float64 	// Event Scheduler will end when "Timeline" reach the "Doom"
}

// ========================================= Engine =========================================
func (e *Engine) Init(timeline, doom float64, event []data.Object) {
	e.Timeline = timeline
	e.Doom = doom

	// init all event into queue
	for _,obj := range event {
		// init new event
		var event data.Event
		event.Event_model = obj 
		event.Event_type = obj.Name 
		event.Event_ts = 0 // (doom - timeline) / doom*100
		e.Schedule(event)
	}
}

func (e *Engine) Pop_front() data.Event {
	// Get the first event 
	event := e.Queue[0]
	// Eliminate the first event (e.g. pop)
	e.Queue = e.Queue[1:]
	return event 
}

func (e *Engine) Push_back(event data.Event) {
	// push into last of queue
	e.Queue = append(e.Queue, event)
}

// schedule event obj 
func (e *Engine) Schedule(event_obj data.Event) {
	// check its model, and time 
	var new_event data.Event 
	new_event = event_obj
	// create random distribution model
	exp := model.Expon{}
	exp.Init(float64(event_obj.Event_model.Lambda))
	poi := model.Poisson{}
	poi.Init(float64(event_obj.Event_model.Lambda))
	uni := model.Uniform{}
	uni.Init(float64(event_obj.Event_model.Lambda), float64(event_obj.Event_model.X)) // usign Lambda as lower, X as upper

	switch t := strings.ToUpper(event_obj.Event_model.Model); t {
		case "POISSON":
			new_event.Event_ts = event_obj.Event_ts + poi.Rand() // interval is exponential
		case "EXPONENTIAL":
			new_event.Event_ts = event_obj.Event_ts + exp.Rand() //exp.Rand_Var(event_obj.Event_ts)
		case "UNIFORM": 
			new_event.Event_ts = event_obj.Event_ts + uni.Rand()
		default:
			// error , need to return with an error code
			// But in here, we just assign an random float number
			new_event.Event_ts = event_obj.Event_ts + model.Rand_gen_float()
	}

	// schedule into event queue
	e.Push_back(new_event)
}

// execute all event from given event object array (read & parse from parser.go)
// Notice: this func must be called after Init()
func (e *Engine) Start() {
	var first_event data.Event

	fmt.Println(e.Timeline, e.Doom)

	// infinite loop - check timeline and doom, 
	// if timeline > doom, break
	for e.Timeline < e.Doom {
		// First pop out the event from event queue
		first_event = e.Pop_front()
		// increment timeline
		e.Timeline = first_event.Event_ts
		// Schedule it, and then increase timeline by adding timestamp of this popping out event
		e.Schedule(first_event)
		fmt.Println("[Event Engine Process] Timeline: ",e.Timeline)
		// record into history, set Event timestamp sync with e.Timeline
		e.History = append(e.History, first_event)

		// Sort the event queue by timestamp (make sure the smallest one came first)
		sort.Sort(data.SortEvent(e.Queue))
	}
}

// ============================================= debug =============================================
func Debug_poisson() float64 {
	p := model.Poisson{} 
	p.Init(3)

	return p.Get(4)
}
func Debug_expon() float64 {
	e := model.Expon{}
	e.Init(3)

	return e.Rand()
}