package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGrinderGrind(t *testing.T) {
	// grinder can grind 20 grams of beans in 1 second
	g := NewGrinder(0, 20)

	startTime := time.Now()
	groundBeans := g.Grind(20)
	endTime := time.Now()

	assert.Equal(t, BeansStateGround, groundBeans.state)
	assert.Equal(t, 0, g.stock.weightGrams)
	assert.Equal(t, 20, groundBeans.weightGrams)
	assert.GreaterOrEqual(t, endTime.Sub(startTime), time.Second)
	assert.LessOrEqual(t, endTime.Sub(startTime), 2*time.Second)
}
