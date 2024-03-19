package main

import (
	"fmt"
	"math"
	"time"
)

// roundUpDivide returns the result of a divided by b, rounded up to the nearest integer
func roundUpDivide(a, b int) float64 {
	return math.Ceil(float64(a) / float64(b))
}

func Log(a ...any) {
	now := time.Now().Format(time.RFC3339)
	fmt.Println(now, a)
}
