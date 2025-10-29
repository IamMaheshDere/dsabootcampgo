package main

func main() {

}

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int) //map[number]idx
	for idx, num := range nums {
		pairNoIdx, ok := numsMap[target-num]
		if ok {
			return []int{idx, pairNoIdx}
		}

		numsMap[num] = idx
	}

	return []int{}
}
