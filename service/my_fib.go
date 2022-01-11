package service

import (
	"errors"
	"fmt"
)

var LessThanTwoError = errors.New("n should be not less than 2")

var LargerThan100Error = errors.New("n should be not larger than 100")

func init() {
	fmt.Println("my_fib init!")
}

func init() {
	fmt.Println("my_fib init again!")
}

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThan100Error
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
