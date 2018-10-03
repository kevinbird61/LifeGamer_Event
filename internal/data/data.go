package data 

// event model
type Object struct {
	Name 	string	`json: "name"`
	Model 	string 	`json: "model"`
	Lambda 	int		`json: "lambda"`
	X		int 	`json: "X"`
}

// event queue
type Event struct {
	Event_model Object // model of event
	Event_type 	string		// type of event: e.g. food_gen, disaster
	Event_ts	float64		// timestamp of event
}

/*
	sorting, need 3 interface  
	- Len
	- Swap
	- Less

	and call:
	- Sort_by_ts()
*/
type SortEvent []Event 

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