package main

import "math"

func main() {
	maxScoreSightseeingPair([]int{8, 1, 5, 2, 6})
}

// func maxScoreSightseeingPair(values []int) int {
// 	maxScore := 0
// 	for i := 0; i < len(values); i++ {
// 		for j := i + 1; j < len(values); j++ {
// 			if values[i]+values[j]+i-j > 0 {
// 				maxScore = values[i] + values[j] + i - j
// 			}
// 		}
// 	}
// 	return maxScore
// }

func maxScoreSightseeingPair(values []int) int {
	maxScore := 0
	if len(values) == 0 {
		return maxScore
	}

	imax := values[0]

	for j := 1; j < len(values); j++ {
		maxScore = int(math.Max(float64(maxScore), float64(imax+values[j]-j)))
		imax = int(math.Max(float64(imax), float64(values[j]+j)))

	}
	return maxScore
}
