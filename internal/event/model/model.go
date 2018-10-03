/*
	Distribution Model 

	Support model:
	- Poisson 	(Poisson distribution)
	- Expon 	(Exponential distribution)
*/

package model

import (
	"math"
	"math/rand"
	"time"
)

// utils
func factorial(k float64) float64 {
	sum:=1.0
	for i:=1.0; i<k; i++ {
		sum *= i
	}

	return sum
}

// 0 ~ upbound
func Rand_gen_int(upbound int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(upbound)
}

// 0.0~1.0
func Rand_gen_float() float64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Float64()
}

// interface 
type Model interface {
	Init()
	Get()
	Rand()
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

func (p *Poisson) Rand() float64 {
	var k float64
	k = Rand_gen_float()
	return (math.Exp(-p.Lambda) * math.Pow(p.Lambda, k)) / factorial(k)
}

// ==================================== Exponential ====================================
type Expon struct {
	Lambda	float64
}

func (e *Expon) Init(lambda float64) {
	e.Lambda = lambda
}

// Return an value on exponential distribution
func (e *Expon) Get(k float64) float64 {
	if k >= 0 {
		return e.Lambda * math.Exp(-e.Lambda*k)
	}
	return 0
}

// Rand on exponential distribution
func (e *Expon) Rand() float64{ 
	return e.Lambda * math.Exp(-e.Lambda*Rand_gen_float())	
}

// Exponential random variable
func (e *Expon) Rand_Var(t float64) float64 {
	if t == float64(0) {
		// assign a small number as "t"
		return math.Exp(-e.Lambda*float64(0.0001))
	}
	return math.Exp(-e.Lambda*t)
}

// ==================================== Uniform ====================================
type Uniform struct {
	Boundary float64
}

func (u *Uniform) Init(lower, upper float64) {
	if upper >= lower {
		u.Boundary = upper - lower 
	} else {
		u.Boundary = lower - upper
	}
}

func (u *Uniform) Rand() float64 {
	return float64(1)/u.Boundary
}