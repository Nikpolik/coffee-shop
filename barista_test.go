package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBaristaMakeCoffee_EnoughStock(t *testing.T) {
	storage := NewBeansStorage(1000)
	// Create brewer and grinder channels
	grinders := make(chan *Grinder, 1)
	brewers := make(chan *Brewer, 1)

	// add a grinder and a brewer to the channels
	grinders <- NewGrinder(0, 20)
	brewers <- &Brewer{ouncesWaterPerSecond: 100}

	b := Barista{
		ID:           1,
		grinders:     grinders,
		brewers:      brewers,
		beansStorage: &storage,
	}
	order := Order{Size: SizeSmall, Recipe: RecipeIntense}
	coffee, error := b.MakeCoffee(order)
	require.NoError(t, error)

	// Check that the channels are not empty and that the barista has returned the grinder and brewer
	assert.Equal(t, 1, len(grinders))
	assert.Equal(t, 1, len(brewers))

	// Check that the coffee has been made
	assert.Equal(t, order.Size.Ounces(), coffee.Ounces)
}

func TestBaristaMakeCoffee_NotEnoughStock(t *testing.T) {
	storage := NewBeansStorage(0)
	// Create brewer and grinder channels
	grinders := make(chan *Grinder, 1)
	brewers := make(chan *Brewer, 1)

	// add a grinder and a brewer to the channels
	grinders <- NewGrinder(0, 20)
	brewers <- &Brewer{ouncesWaterPerSecond: 100}

	b := Barista{
		ID:           1,
		grinders:     grinders,
		brewers:      brewers,
		beansStorage: &storage,
	}
	order := Order{Size: SizeSmall, Recipe: RecipeIntense}
	coffee, error := b.MakeCoffee(order)
	require.Error(t, error)

	// Check that the channels are not empty and that the barista has returned the grinder and brewer
	assert.Equal(t, 1, len(grinders))
	assert.Equal(t, 1, len(brewers))

	// Check that the coffee has not been made
	assert.Equal(t, 0, coffee.Ounces)
}

func TestBaristaWork(t *testing.T) {
	group := sync.WaitGroup{}
	storage := NewBeansStorage(100)
	// Create brewer and grinder channels
	grinders := make(chan *Grinder, 1)
	brewers := make(chan *Brewer, 1)
	// add a grinder and a brewer to the channels
	grinders <- NewGrinder(0, 20)
	brewers <- &Brewer{ouncesWaterPerSecond: 20 * 12}

	b := Barista{1, grinders, brewers, &storage}
	orderQueue := make(chan Order, 1)
	coffeeQueue := make(chan Coffee, 1)

	go b.Work(orderQueue, coffeeQueue, &group)
	order := Order{Size: SizeSmall, Recipe: RecipeMild}
	orderQueue <- order
	coffee := <-coffeeQueue

	// Check that the coffee has been made
	assert.Equal(t, order.Size.Ounces(), coffee.Ounces)
}
