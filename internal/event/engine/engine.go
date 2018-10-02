package engine

import (
	"../../data"
	"../model"
)

// execute all event from given event object array (read & parse from parser.go)
func Execute(event []data.Object) {

}

// debug 
func Debug_poisson() float64 {
	p := model.Poisson{} 
	p.Init(3)

	return p.Get(4)
}