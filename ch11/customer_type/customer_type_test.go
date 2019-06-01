package customer_type

import (
	"time"
	"fmt"
	"testing"
)

type InnerFuction func (op int) int

func timeSpent(inner InnerFuction) InnerFuction {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		// end := time.Now()
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second*1)
	return op
}

func TestFunc(t *testing.T) {
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}