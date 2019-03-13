package main

import (
	"fmt"
	"math"
)

const DELTA = 1e-8

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can not Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	prez := 0.0
	curz := 1.0
	for {
		prez = curz
		curz = prez - (prez*prez-x)/(2*prez)
		if math.Abs(prez-curz) < DELTA {
			break
		}
	}
	return curz, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
