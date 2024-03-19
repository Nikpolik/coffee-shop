package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBrewer_Brew(t *testing.T) {
	b := &Brewer{ouncesWaterPerSecond: 12}
	startTime := time.Now()
	coffee := b.Brew(Beans{weightGrams: 48}, Recipe{WaterOunces: 0.5, BeansGrams: 2})
	endTime := time.Now()
	assert.Equal(t, 24, coffee.Ounces)
	assert.GreaterOrEqual(t, endTime.Sub(startTime), 2*time.Second)
	assert.LessOrEqual(t, endTime.Sub(startTime), 3*time.Second)
}
