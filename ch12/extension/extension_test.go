package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" Pet ", host)
}

type Dog struct {
	Pet //匿名嵌套类型
}

// func (d *Dog) Speak() {
// 	fmt.Println("Wang!")
// }

// func (d *Dog) SpeakTo(host string) {
// 	d.Speak()
// 	fmt.Println(" Dog ", host)
// }

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Chao")
}
