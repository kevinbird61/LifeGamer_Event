package engine

import (
	"fmt"
	"sort"
	"strings"

	"../../data"
	"../model"
)

// event queue
type Event struct {
	Event_model data.Object // model of event
	Event_type 	string		// type of event: e.g. food_gen, disaster
	Event_ts	float64		// timestamp of event
}

// Class: Engine
type Engine struct {
	Queue		[]Event 
	History		[]Event 
	Timeline	float64
	Doom		float64 	// Event Scheduler will end when "Timeline" reach the "Doom"
}

type SortEvent []Event 

// ========================================= Engine =========================================
func (e *Engine) Init(timeline, doom float64, event []data.Object) {
	e.Timeline = timeline
	e.Doom = doom

	// init all event into queue
	for _,obj := range event {
		// init new event
		var event Event
		event.Event_model = obj 
		event.Event_type = obj.Name 
		event.Event_ts = 0
		e.Schedule(event)
	}
}

func (e *Engine) Pop_front() Event {
	// Get the first event 
	event := e.Queue[0]
	// Eliminate the first event (e.g. pop)
	e.Queue = e.Queue[1:]
	return event 
}

func (e *Engine) Push_back(event Event) {
	// push into last of queue
	e.Queue = append(e.Queue, event)
}

/*
	sorting, need 3 interface  
	- Len
	- Swap
	- Less

	and call:
	- Sort_by_ts()
*/
func (q SortEvent) Len() int {
	return len(q)
}

func (q SortEvent) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q SortEvent) Less(i, j int) bool {
	if q[i].Event_ts <= q[j].Event_ts {
		return true
	}
	return false 
}

// schedule event obj 
func (e *Engine) Schedule(event_obj Event) {
	// check its model, and time 
	var new_event Event 
	new_event = event_obj
	// new_event.Event_model = event_obj
	// new_event.Event_type = event_obj.Name
	exp := model.Expon{}
	exp.Init(float64(event_obj.Event_model.Lambda))

	switch t := strings.ToUpper(event_obj.Event_model.Model); t {
		case "POISSON":
			new_event.Event_ts = exp.Rand() // interval is exponential
		case "EXPONENTIAL":
			new_event.Event_ts = exp.Rand_Var(event_obj.Event_ts)
		default:
			// error , need to return with an error code
			// But in here, we just assign an random float number
			new_event.Event_ts = model.Rand_gen_float()
	}

	// schedule into event queue
	e.Push_back(new_event)
}

// execute all event from given event object array (read & parse from parser.go)
// Notice: this func must be called after Init()
func (e *Engine) Start() {
	var first_event Event

	fmt.Println(e.Timeline, e.Doom)

	// infinite loop - check timeline and doom, 
	// if timeline > doom, break
	for e.Timeline < e.Doom {
		// First pop out the event from event queue
		first_event = e.Pop_front()
		// increment timeline
		e.Timeline += first_event.Event_ts
		// Schedule it, and then increase timeline by adding timestamp of this popping out event
		e.Schedule(first_event)

		// record into history, set Event timestamp sync with e.Timeline
		first_event.Event_ts = e.Timeline
		e.History = append(e.History, first_event)

		// Sort the event queue by timestamp (make sure the smallest one came first)
		sort.Sort(SortEvent(e.Queue))
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