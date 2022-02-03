package main

import "math"

var decay = .5

func makeShouldIterate() func(speed int, tick int) bool {
	speedTickIterationMap := make(map[int]map[int]int)
	for i := 1; i < 10; i++ {
		speedTickIterationMap[i] = make(map[int]int)
	}
	mode := 1000
	return func(speed int, tick int) bool {
		step := tick % mode

		if speed == 0 {
			return false
		}

		tickIterationMap := speedTickIterationMap[speed]
		var iteration int
		if iter, ok := tickIterationMap[step]; ok {
			iteration = iter
		} else {
			factor := math.Pow(decay, float64(9-speed))
			iteration = int(float64(step) * factor)
			tickIterationMap[step] = iteration
		}

		idx := mode - 1
		if step > 0 {
			idx = step - 1
		}
		prev := tickIterationMap[idx]
		return iteration > prev
	}
}
