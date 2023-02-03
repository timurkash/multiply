package main

import (
	"fmt"
	"strconv"
)

func verifyNum(num []int8) error {
	if len(num) == 0 {
		return fmt.Errorf("num is empty")
	}
	if num[0] == 0 {
		return fmt.Errorf("leading 0")
	}
	for _, d := range num {
		if !(d>=0 && d <= 9) {
			return fmt.Errorf("%d is wrong digit", d)
		}
	}
	return nil
}

func stringToSlice(str string) ([]int8, error) {
	res := make([]int8, len(str))
	for i, r := range str {
		if !(r >= '0' && r <= '9') {
			return nil, fmt.Errorf("%v is wrong rune", r)
		}
		res[i] = int8(r - '0')
	}
	return res, nil
}

func sliceToString(num []int8) (res string) {
	for _, n := range num {
		res += strconv.Itoa(int(n))
	}
	return res
}