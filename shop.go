package main

import (
	"fmt"
	"sync"
)

// import "math/rand"

type Size string

const (
	SizeSmall  Size = "small"
	SizeMedium Size = "medium"
	SizeLarge  Size = "large"
)

func (c *Size) Ounces() int {
	if *c == SizeLarge {
		return 16
	}
	if *c == SizeMedium {
		return 12
	}
	return 8
}

type Order struct {
	ID     int
	Size   Size
	Recipe Recipe
}

var RecipeIntense = Recipe{
	BeansGrams:  4,
	WaterOunces: 1,
}

var RecipeBalanced = Recipe{
	BeansGrams:  2,
	WaterOunces: 1,
}

var RecipeMild = Recipe{
	BeansGrams:  2,
	WaterOunces: 1.5,
}

type Recipe struct {
	BeansGrams  int
	WaterOunces float64
}

type BeansStorage struct {
	totalAmount int
	mutex       *sync.Mutex
}

func NewBeansStorage(totalAmount int) BeansStorage {
	return BeansStorage{
		totalAmount: totalAmount,
		mutex:       &sync.Mutex{},
	}
}

func (bs *BeansStorage) GetBeans(amount int) (Beans, error) {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()
	if bs.totalAmount < amount {
		return Beans{}, fmt.Errorf("Not enough beans")
	}
	bs.totalAmount -= amount
	return Beans{
		weightGrams: amount,
	}, nil
}

type CoffeeShop struct {
	orders   chan Order
	coffees  chan Coffee
	grinders chan *Grinder
	brewers  chan *Brewer
	baristas []Barista
	storage  BeansStorage
	wg       *sync.WaitGroup
}

func NewCoffeeShop(grinders []*Grinder, brewers []*Brewer, totalBeans, numOfBaristas int) CoffeeShop {
	// create a channel for each grinder and brewer
	grinderChans := make(chan *Grinder, len(grinders))
	brewerChans := make(chan *Brewer, len(brewers))
	storage := NewBeansStorage(totalBeans)
	wg := &sync.WaitGroup{}

	for i := range grinders {
		grinderChans <- grinders[i]
	}

	for i := range brewers {
		brewerChans <- brewers[i]
	}

	baristas := make([]Barista, numOfBaristas)
	for i := 0; i < numOfBaristas; i++ {
		id := i
		fmt.Println("barista id: ", id)
		baristas[i] = Barista{
			ID:           id,
			grinders:     grinderChans,
			brewers:      brewerChans,
			beansStorage: &storage,
		}
	}

	return CoffeeShop{
		grinders: grinderChans,
		brewers:  brewerChans,
		wg:       wg,
		baristas: baristas,
		storage:  storage,
	}
}

func (cs *CoffeeShop) Start() (chan Order, chan Coffee) {
	orderQueue := make(chan Order, 100)
	coffeeQueue := make(chan Coffee, 100)
	// start the baristas
	for _, barista := range cs.baristas {
		barista := barista
		go barista.Work(orderQueue, coffeeQueue, cs.wg)
	}
	return orderQueue, coffeeQueue
}

func (cs *CoffeeShop) Close() {
	close(cs.orders)
	cs.wg.Wait()
}
