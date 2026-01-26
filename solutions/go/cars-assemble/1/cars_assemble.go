package cars
// import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    var x int = productionRate
    test := float64(x)
    return test * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var carPerHour float64 = CalculateWorkingCarsPerHour(productionRate, successRate);
    carPerHourInt := int(carPerHour)
    return carPerHourInt / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	//10,000 per car
    // 95, 000 for every 10 cars
    tens := carsCount / 10
    unit := carsCount % 10
    result := (tens * 95000) + (unit * 10000)
    return uint(result)
}
