# Best Sightseeing Pair

## Problem Statement

You are given an integer array `values` where values[i] represents the value of the i-th sightseeing spot. Two sightseeing spots i and j have a distance `j - i` between them.

A pair of sightseeing spots (i < j) has a score of `values[i] + values[j] + i - j`. The task is to find the maximum score of a pair of sightseeing spots.

## Example

### Example 1

**Input:**
```
values = [8, 1, 5, 2, 6]
```

**Output:**
```
11
```

**Explanation:**
```
i = 0, j = 2, values[i] + values[j] + i - j = 8 + 5 + 0 - 2 = 11
```

### Example 2

**Input:**
```
values = [1, 2]
```

**Output:**
```
2
```

**Explanation:**
```
i = 0, j = 1, values[i] + values[j] + i - j = 1 + 2 + 0 - 1 = 2
```

## Constraints

- `2 <= values.length <= 5 * 10^4`
- `1 <= values[i] <= 1000`

## Solution

Need to find the max value of values[i] + i + value[j] - j
Lets consider imax = values[i] + i and jmax = values[j] - j

We iterate throuth the array starting and for each index of j, calculate the score for each pair of (i,j) as imax + jmax
and for each iteration update the value of imax by taking maximum value imax at index j and index j-1



### Pseudocode

```
init the variable imax and score
imax = values[0] + 0 
maxScore = 0

range over values from index 1
for j:=0; j<len(values) ; j++{
    maxScore = max(maxScore, imax+values[j]-j)
    imax = max(imax, values[i] + i) 
}

return maxScore

```

### Time complexity: O(n)

 ### Space complexity 0(1)
