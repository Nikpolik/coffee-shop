package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoffeeShop_Start(t *testing.T) {
	// Create baristas, grinders, and brewers
	grinders := []*Grinder{
		{ID: 0, gramsPerSecond: 20},
		{ID: 1, gramsPerSecond: 20},
		{ID: 2, gramsPerSecond: 20},
		{ID: 2, gramsPerSecond: 20},
	}
	brewers := []*Brewer{
		{ouncesWaterPerSecond: 100, ID: 0},
		{ouncesWaterPerSecond: 100, ID: 1},
		{ouncesWaterPerSecond: 100, ID: 2},
	}

	// create a channel for each grinder and brewer
	cs := NewCoffeeShop(grinders, brewers, 1000, 5)

	orderQueue, coffeeQueue := cs.Start()
	// Check that the order and coffee queues are not nil
	assert.NotNil(t, orderQueue)
	assert.NotNil(t, coffeeQueue)
	orders := []Order{
		{Size: SizeSmall, ID: 0, Recipe: RecipeIntense},
		{Size: SizeSmall, ID: 1, Recipe: RecipeIntense},
		{Size: SizeSmall, ID: 2, Recipe: RecipeMild},
		{Size: SizeLarge, ID: 3, Recipe: RecipeBalanced},
		{Size: SizeMedium, ID: 4, Recipe: RecipeBalanced},
		{Size: SizeLarge, ID: 5, Recipe: RecipeMild},
		{Size: SizeMedium, ID: 6, Recipe: RecipeIntense},
		{Size: SizeSmall, ID: 2, Recipe: RecipeMild},
	}
	for i := 0; i < len(orders); i++ {
		orderQueue <- orders[i]
	}

	for i := 0; i < len(orders); i++ {
		coffee := <-coffeeQueue
		found := false
		for j := 0; j < len(orders); j++ {
			if orders[j].ID == coffee.ID {
				found = true
				assert.Equal(t, orders[j].Size.Ounces(), coffee.Ounces)
				break
			}
		}
		assert.True(t, found)
	}
}
