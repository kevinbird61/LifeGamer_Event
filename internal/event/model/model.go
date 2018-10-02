/*
	Distribution Model 

	Support model:
	- Poisson 	(Poisson distribution)
	- Expon 	(Exponential distribution)
*/

package model

import (
	"math"
)

// utils
func factorial(k float64) float64 {
	sum:=1.0
	for i:=1.0; i<k; i++ {
		sum *= i
	}

	return sum
}

// ==================================== Poisson ====================================
type Poisson struct {
	Lambda 	float64
}

func (p *Poisson) Init(lambda float64) {
	p.Lambda = lambda
}

func (p *Poisson) Get(k float64) float64{
	// Get P(X=k) value
	return (math.Exp(-p.Lambda) * math.Pow(p.Lambda, k)) / factorial(k)
}

// ==================================== Exponential ====================================
type Expon struct {
	Lambda	int
}