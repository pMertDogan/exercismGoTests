package cars

const multiplier = 221

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if speed < 1 {
		return 0.0
	} else if speed < 5 {
		return 1.0
	} else if speed <9 {
		return 0.9
	}else{
		return 0.77
	}

}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return SuccessRate(speed) * float64(speed)* multiplier
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed)/ 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	x := CalculateProductionRatePerHour(speed)
	if x == 0 {
		return x
	}else if x > limit {
		return limit
	}else{
		return x
	}
}
