package main

import (
	"math"
	"time"
)

type Coffee struct {
	ID     int
	Ounces int
}

type Brewer struct {
	ID int
	// assume we have unlimited water, but we can only run a certain amount of water per second into our brewer + beans
	ouncesWaterPerSecond int
}

func (b *Brewer) Brew(beans Beans, recipe Recipe) Coffee {
	water := math.Ceil(float64(beans.weightGrams) * recipe.WaterOunces)
	brewTime := roundUpDivide(int(water), b.ouncesWaterPerSecond)
	time.Sleep(time.Duration(brewTime) * time.Second)
	result := beans.weightGrams / recipe.BeansGrams
	return Coffee{
		Ounces: result,
	}
}
