package main

func main() {
	maxSubsequence([]int{8, 1, 5, 2, 6}, 2)
}

func maxSubsequence(nums []int, k int) []int {
	maxsubSeq := make([]int, k)
	if len(nums) == 0 {
		return maxsubSeq
	}

	sum, maxSum := 0, 0
	for i, j := 0, 0; j < len(nums); j++ {
		if j-i+1 < k {
			sum += nums[j]
		} else if j-i+1 == k {
			sum = sum + nums[j]
			if sum > maxSum {
				maxSum = sum
				copy(maxsubSeq, nums[i:j+1])
			}
			sum = sum - nums[i]
			i++
		}
	}
	return maxsubSeq
}
