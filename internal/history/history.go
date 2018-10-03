/*
	Record all the event with timestamp, describe all event like a history.
*/

package history

import (
	"../data"
)

type History struct {
	Map		map[string]int
}

func (h *History) Init(event_list []data.Event) {
	h.Map = make(map[string]int)
	for _,obj := range event_list {
		h.Map[obj.Event_type]++
	}
}