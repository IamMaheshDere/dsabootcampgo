# Two sum

## Problem Statement
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

## Example

### Example 1

**Input:**
```
nums = [2,7,11,15], target = 9
```

**Output:**
```
Output: [0,1]
```

**Explanation:**
```
Because nums[0] + nums[1] == 9, we return [0, 1].
```

## Solution

### Pseudocode

```
map [number] idx
for idx, num := range nums {
    anotherNoIdx, ok := map[target-num]
    if ok{
        return []int{anotherNoIdx, idx}
    }

    map[num]=idx
}
```

### Time complexity: O(n)

### Space complexity O(n)
