package getinterest

import (
	"math"
)

// Total Time in Years
type Myvalues struct {
	Amount       float64
	InterestRate float64
	Totaltime    float64
}

func (a Myvalues) SimpleInterest() float64 {
	return (a.Amount * a.InterestRate * a.Totaltime) / 100
}

func (a Myvalues) ComppoundInterest() float64 {
	var amount float64 = a.Amount * (math.Pow((1 + a.InterestRate/100), a.Totaltime))
	return amount
}
