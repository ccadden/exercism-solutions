package cars

const MinutesPerHour int = 60
const CostOfTenCars uint = 95000
const CostOfOneCar uint = 10000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate/100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / float64(MinutesPerHour))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOfTen := carsCount / 10
	remainder := carsCount % 10
	return uint(groupsOfTen) * CostOfTenCars + uint(remainder) * CostOfOneCar
}
