package intcode

import "os"

type Intcode interface {
	Calculate(intput []int) []int
	Restore(input []int) []int
}

type intcode struct{}

func NewIntcode() Intcode {
	return &intcode{}
}

func (i *intcode) Calculate(input []int) []int {
	index := 0

	for index + 4 <= len(input) {
		opCode := input[index]

		switch opCode {
		case 1:
			i.add(input, index + 1)
		case 2:
			i.multiply(input, index + 1)
		case 99:
			//os.Exit(0)
		default:
			os.Exit(1)
		}

		index += 4
	}

	return input
}

func (i *intcode) Restore(input []int) []int {
	input[1] = 12
	input[2] = 2

	return input
}

func (i *intcode) add(input []int, index int) []int {
	pos1 := input[index]
	pos2 := input[index+1]
	result := input[index+2]

	input[result] = input[pos1] + input[pos2]

	return input
}

func (i *intcode) multiply(input []int, index int) []int {
	pos1 := input[index]
	pos2 := input[index+1]
	result := input[index+2]

	input[result] = input[pos1] * input[pos2]

	return input
}