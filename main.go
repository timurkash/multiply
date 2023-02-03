package main

import "fmt"

func main() {
	num := []int8{1, 0, 0, 0, 0, 1}
	num2 := []int8{9, 9, 9, 9, 9}
	fmt.Println(multiply(num, num2))
}

func multiply(num1, num2 []int8) (res []int8) {
	reverse := func(num []int8) (res []int8) {
		l := len(num)
		res = make([]int8, l)
		for i, d := range num {
			res[l-i-1] = d
		}
		if res[0] == 0 {
			res = res[1:]
		}
		return
	}
	l := len(num2)
	var nums [][]int8
	for i := range num2 {
		var term []int8
		for j := 0; j < i; j++ {
			term = append(term, 0)
		}
		term = append(term, multiplyDigit(num1, num2[l-i-1])...)
		for j := 0; j < l-i; j++ {
			term = append(term, 0)
		}
		nums = append(nums, term)
	}
	l2 := len(nums[0])
	var (
		sum  int8
		over int8
		d    int8
	)
	for j := 0; j < l2; j++ {
		sum = over
		for i := 0; i < l; i++ {
			sum += nums[i][j]
		}
		d = sum % 10
		over = (sum - d) / 10
		res = append(res, d)
	}
	return reverse(res)
}

func multiplyDigit(num []int8, digit int8) (res []int8) {
	var (
		over int8
		d    int8
		r    int8
	)
	for i := len(num) - 1; i >= 0; i-- {
		r = num[i]*digit + over
		d = r % 10
		over = (r - d) / 10
		res = append(res, d)
	}
	if over > 0 {
		res = append(res, over)
	}
	return
}
