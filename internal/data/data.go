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