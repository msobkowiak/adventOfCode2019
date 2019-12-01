package main

import (
	"advent_of_code_2019/day_1/counter"
	"advent_of_code_2019/day_1/reader"
)

func main() {
	r := reader.NewFileReader()
	moduleMasses := r.Read("input.txt")

	c := counter.NewCounter()
	result := c.CountExtraFuel(moduleMasses)

	println(result)
}
