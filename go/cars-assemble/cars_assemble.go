package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var successPercentage = successRate / 100
	var carsProducedSuccessfully = float64(productionRate) * successPercentage
	return float64(carsProducedSuccessfully)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var res = int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
	return res
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var res = carsCount / 10
	var remainedCars = carsCount % 10
	var finalCost = (res * 95000) + (remainedCars * 10000)
	return uint(finalCost)

}
