package counter

type Counter interface {
	CountFuel(masses []int) int
	CountExtraFuel(masses []int) int
}

type counter struct{}

func NewCounter() Counter {
	return &counter{}
}

func (c *counter) CountFuel(masses []int) int {
	var result int
	for _, mass := range masses {
		result += count(mass)
	}

	return result
}


func (c *counter) CountExtraFuel(masses []int) int {
	var result int
	for _, mass := range masses {
		result += countRecursive(mass)
	}

	return result
}

func count(input int) int {
	return input / 3 - 2
}

func countRecursive(input int) int {
	var result int
	for input > 0 {
		value := count(input)
		if  value > 0 {
			result += value
		}
		input = value
	}

	return result
}
