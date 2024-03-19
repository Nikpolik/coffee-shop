## Coffee Shop Simulation

This code simulates a coffee shop order process with baristas, grinders, brewers, and bean storage. It uses channels and goroutines to approximate the scarcity of each item in the shop.

### Running the Simulation

This code requires Go installed on your system.

1.  Clone or download the repository containing the code.
2.  Open a terminal and navigate to the directory containing the code.
3.  Run the following command to build and execute the program:

```bash
go run .
go run . --help
```

Optional flags:

- `-baristas int` Number of baristas working at the coffee shop (default 2)
- `-beans int` Total amount of coffee beans (in grams) available (default 1000)
- `-brewSpeed int` Brew speed in ounces per second (default 100)
- `-brewers int` Number of coffee brewers available (default 2)
- `-grindSpeed int` Grind speed in grams per second (default 20)
- `-grinders int` Number of coffee grinders available (default 2)

This will start the coffee shop simulation and print order processing information to the console.

### Testing

The code includes unit tests for most of the functions. You can run the tests using the following command:

```bash
go test ./...
```

### Potential Improvements

Better handling of how the coffee shop starts and stops. Not 100% sure that returning the channel is the best way to handle this.

Use context to handle the cancellation of the simulation.

Better tests for the channels and routines with multiple baristas, grinders, and brewers. The current tests only check the happy path with one of each.

Add more complex logic to the simulation, such as different types of coffee beans, multiple types of coffee drinks, and more detailed order processing steps.

Add statistics and metrics tracking to the simulation to analyze the performance of the coffee shop and identify areas for improvement.
