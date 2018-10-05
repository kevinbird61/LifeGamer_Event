/*
	Record all the event with timestamp, describe all event like a history.
*/

package history

import (
	"sort"

	"../data"
)

type History struct {
	Map		map[string]int		// total number of each event
	Sim		[]data.Event 		// Simulation (event list) of History, the game will base on this event list to compute
	Tstart	float64				// start time
	Tend	float64				// end time
	Story	map[float64][]data.Event 		// Story mode - using timestamp as index, to describe events
}

// Count the event number 
func (h *History) Init(tstart, tend float64, event_list []data.Event) {
	// record duration - for scale (game compute)
	h.Tstart = tstart 
	h.Tend = tend 

	// count
	h.Map = make(map[string]int)
	h.Story = make(map[float64][]data.Event)
	for _,obj := range event_list {
		// type as index
		h.Map[obj.Event_type]++
		// timestamp as index
		h.Story[obj.Event_ts] = append(h.Story[obj.Event_ts], obj)
	}

	// record history
	h.Sim = event_list
}

// Let new event list add into History (For runtime additional event)
func (h *History) Add_event_list(event_list []data.Event) {
	// append 
	h.Sim = append(h.Sim, event_list...)
	// re-count
	for _,obj := range event_list {
		h.Map[obj.Event_type]++
	}
	// sort 
	sort.Sort(data.SortEvent(h.Sim))
}
