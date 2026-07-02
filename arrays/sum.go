package arrays

func Sum(nums []int) int {
	var sum int = 0
	for _, num := range nums {
		sum += num 
	}
	return sum
}
