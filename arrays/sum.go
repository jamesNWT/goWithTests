package arrays

func Sum(nums []int) int {
	var sum int = 0
	for _, num := range nums {
		sum += num 
	}
	return sum
}

func SumAll(slicesToSum... []int) []int {
	ans := make([]int, len(slicesToSum))
	for i, nums := range slicesToSum {
		ans[i] = Sum(nums)
	}
	return ans
}
