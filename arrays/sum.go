package arrays

func Sum(nums []int) int {
	var sum int = 0
	for _, num := range nums {
		sum += num 
	}
	return sum
}

func SumAll(slicesToSum... []int) []int {
	var ans []int
	for _, nums := range slicesToSum {
		ans = append(ans, Sum(nums))
	}
	return ans
}

func SumAllTails(slicesToSum... []int) []int {
	var ans []int
	for _, slice := range slicesToSum {
		if len(slice) >= 1 {
			ans = append(ans, Sum(slice[1:]))
		} else {
			ans = append(ans, 0)
		}
	}
	return ans
}
