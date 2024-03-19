package main

import (
	"time"
)

type BeansState string

const (
	BeansStateGround BeansState = "ground"
	BeansStateWhole  BeansState = "whole"
)

type Beans struct {
	weightGrams int
	// Indicates if the beans are ground or whole
	state BeansState
}

type Grinder struct {
	ID               int
	gramsPerSecond   int
	stock            Beans
	capacity         int
	restockPerSecond int
}

func NewGrinder(id, gramsPerSecond int) *Grinder {
	return &Grinder{
		ID:             id,
		gramsPerSecond: gramsPerSecond,
	}
}

func (g *Grinder) Stock() int {
	return g.stock.weightGrams
}

func (g *Grinder) Grind(grams int) Beans {
	grindTime := roundUpDivide(grams, g.gramsPerSecond)
	time.Sleep(time.Duration(grindTime) * time.Second)

	return Beans{
		weightGrams: grams,
		state:       BeansStateGround,
	}
}
