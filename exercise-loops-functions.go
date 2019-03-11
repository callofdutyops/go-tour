package main

import (
	"fmt"
    "math"
)

const DELTA = 1e-8

func Sqrt(x float64) (res float64, round int) {
	prez := 0.0
    curz := 1.0
    r := 0
	for {
        prez = curz
		curz = prez - (prez*prez - x) / (2 * prez)
        r++
        if math.Abs(prez - curz) < DELTA {
            break
        }
	}
	return curz, r
}

func main() {
    calculated_v, round := Sqrt(2)
    fmt.Println("The calculated sqrt(2) is:", calculated_v, "after", round, "round.")
    fmt.Println("The real sqrt(2) is:", math.Sqrt(2))
}
