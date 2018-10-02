package data 

type Object struct {
	Name 	string	`json: "name"`
	Model 	string 	`json: "model"`
	Lambda 	int		`json: "lambda"`
	X		int 	`json: "X"`
}