
// Cheat Sheet

// 1. To check if a element exsists in a HashMap:
if requiredValue, isPresent := HashMap[target_key]; isPresent {
	Return
} else {
	Return
}

// 2. For Loop to iterate through a HashMap with Key and Value:
for index, value := range HashMap {
}

// 3. To declare an array:
nums := []int{2, 7, 11, 15}

// 4. To initalize an array:
suffixList := make([]int, len(nums))

// 5. To copy without reference of array:
copy(suffixList, nums)

// 6. Max of 2 int
int(math.Max(float64(res), float64(prices[rightPointer]-prices[leftPointer])))

// 7. Make a 2D array
dp := make([][]int, len(text1)+1)
for v := range dp {
	dp[v] = make([]int, len(text2)+1)
}
