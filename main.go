package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// Define flags
	var numBaristas int
	var totalBeans int
	var numGrinders int
	var numBrewers int
	var grindSpeed int
	var brewSpeed int

	flag.IntVar(&numBaristas, "baristas", 2, "Number of baristas working at the coffee shop")
	flag.IntVar(&totalBeans, "beans", 1000, "Total amount of coffee beans (in grams) available")
	flag.IntVar(&numGrinders, "grinders", 2, "Number of coffee grinders available")
	flag.IntVar(&numBrewers, "brewers", 2, "Number of coffee brewers available")
	flag.IntVar(&grindSpeed, "grindSpeed", 20, "Grind speed in grams per second")
	flag.IntVar(&brewSpeed, "brewSpeed", 100, "Brew speed in ounces per second")
	flag.Parse() // Parse command line flags
	Log("Starting coffee shop with", numBaristas, "baristas,", totalBeans, "grams of beans,", numGrinders, "grinders,", numBrewers, "brewers,", "grind speed", grindSpeed, "grams per second, and brew speed", brewSpeed, "ounces per second")

	if numBaristas <= 0 || totalBeans <= 0 || numGrinders <= 0 || numBrewers <= 0 {
		fmt.Println("Invalid arguments. All values must be positive integers.")
		return
	}

	// Create coffee shop with specified resources
	cs := NewCoffeeShop(createGrinders(numGrinders, grindSpeed), createBrewers(numBrewers, brewSpeed), totalBeans, numBaristas)

	// Start the coffee shop and get order/coffee channels
	orderQueue, _ := cs.Start()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		os.Exit(1)
	}()

	id := 1
	for {
		orderQueue <- Order{Size: SizeSmall, ID: id, Recipe: RecipeIntense}
	}
}

// Helper functions to create grinders and brewers
func createGrinders(num int, speed int) []*Grinder {
	grinders := make([]*Grinder, num)
	for i := 0; i < num; i++ {
		grinders[i] = NewGrinder(i, speed) // Adjust grind speed (grams per second) as needed
	}
	return grinders
}

func createBrewers(num int, speed int) []*Brewer {
	brewers := make([]*Brewer, num)
	for i := 0; i < num; i++ {
		brewers[i] = &Brewer{ouncesWaterPerSecond: speed} // Adjust brewing speed (ounces per second) as needed
	}
	return brewers
}
