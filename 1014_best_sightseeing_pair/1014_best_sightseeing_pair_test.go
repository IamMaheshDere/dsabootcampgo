package main

import "testing"

func TestMaxScoreSightseeingPair(t *testing.T) {
	tests := []struct {
		values []int
		want   int
	}{
		{[]int{8, 1, 5, 2, 6}, 11},
		{[]int{1, 2}, 2},
		{[]int{}, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxScoreSightseeingPair(tt.values); got != tt.want {
				t.Errorf("maxScoreSightseeingPair(%v) = %v, want %v", tt.values, got, tt.want)
			}
		})
	}
}
