package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//if i,ok := p.(int); ok {
	//	fmt.Println("integer: ", i)
	//	return
	//}
	//if s,ok := p.(string); ok {
	//	fmt.Println("string: ", s)
	//	return
	//}
	//fmt.Println("Unknown Type")
	switch i := p.(type) {
	case int:
		fmt.Println("integer: ", i)
	case string:
		fmt.Println("string: ", i)
	default:
		fmt.Println("Unknown Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("Hello World")
	DoSomething(3.5)
}
