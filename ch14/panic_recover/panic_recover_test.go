package panic_recover

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	//defer func() {
	//	fmt.Println("Finally")
	//}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()

	fmt.Println("Start")

	panic(errors.New("Something wrong!"))
	//os.Exit(-1)
}