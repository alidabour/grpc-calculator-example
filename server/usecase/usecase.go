package usecase

import "github.com/alfredxing/calc/compute"

//Eval usecase interface
type Eval interface {
	Eval(in string) (float64, error)
}

type eval struct{}

//New return new struct implement Eval usecase interface
func New() Eval {
	return eval{}
}

//Eval evaluate expression
//input ex: 1+2*2/3
func (eval) Eval(in string) (float64, error) {
	res, err := compute.Evaluate(in)
	if err != nil {
		return 0, err
	}
	return res, nil
}
