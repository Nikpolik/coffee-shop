package main

import (
	"fmt"
	"sync"
)

type Barista struct {
	ID           int
	grinders     chan *Grinder
	brewers      chan *Brewer
	beansStorage *BeansStorage
}

func (b *Barista) Work(orderQueue <-chan Order, coffeeQueue chan<- Coffee, group *sync.WaitGroup) {
	fmt.Println("Starting work for barista id: ", b.ID)
	group.Add(1)
	for order := range orderQueue {
		coffee, err := b.MakeCoffee(order)
		coffee.ID = order.ID
		if err != nil {
			// Do some error handling here
			fmt.Println(err)
			break
		}
		Log("Barista", b.ID, "done with order", order.ID)
		coffeeQueue <- coffee
	}
	group.Done()
}

func (b *Barista) MakeCoffee(order Order) (Coffee, error) {
	Log("Barista", b.ID, "is making order", order.ID, "of size", order.Size)
	amountOfBeans := order.Size.Ounces() * order.Recipe.BeansGrams
	ungroundBeans, err := b.beansStorage.GetBeans(amountOfBeans)
	if err != nil {
		return Coffee{}, err
	}
	// choose the next available grinder and grind the beans
	grinder := <-b.grinders
	Log("Barista", b.ID, "is grinding in", grinder.ID)
	groundBeans := grinder.Grind(ungroundBeans.weightGrams)
	// return the grinder to the channel
	b.grinders <- grinder
	brewer := <-b.brewers
	Log("Barista", b.ID, "is brewing in", brewer.ID)
	c := brewer.Brew(groundBeans, order.Recipe)
	b.brewers <- brewer
	return c, nil
}
