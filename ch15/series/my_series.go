package series

import (
	"fmt"
	"github.com/pkg/errors"
)

func init() {
	fmt.Println("init: 1")
}

func init() {
	fmt.Println("init: 2")
}

//func square(n int) int {
//	return n * n
//}

func GetFibonacci(n int) ([]int, error) {
	buff := []int{1, 1}
	if n < 2 || n > 100 {
		return nil, errors.New("parameter not valid, should be in [2,100]")
	}
	for i := 2; i < n; i++ {
		buff = append(buff, buff[i-2]+buff[i-1])
	}
	return buff, nil
}
