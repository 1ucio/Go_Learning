package fib

import (
	"testing"
)

// var a int;
// var b int;

func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1
	// var (
	// 	a int = 1
	// 	b = 2  // Go 具有一定类型推断的能力
	// )
	// a = 1
	a := 1
	b := 1
	// fmt.Println(a, " ")
	t.Log(a)
	for i := 0; i < 5; i++ {
		// fmt.Println(" ", b)
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	// fmt.Println()
	t.Log()

}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp
	a, b = b, a
	t.Log(a, b)
}
