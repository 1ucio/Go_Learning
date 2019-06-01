package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))

	s = "\xE4\xB8\xA5"
	t.Log(s)
	s = "中"
	t.Log(len(s))

	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

	z := "\xE4\xB8\xAD"
	t.Log(z)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"

	for _, c := range s {
		t.Logf("%[1]c %[1]d %[1]x", c)
	}

	z := "中"
	t.Logf("%[1]x", z)
}

func TestStringArraySlice(t *testing.T) {
	s := "Hello"
	a := [3]int{1,2,3}
	b := []int{1,2,3}
	t.Log(len(s), cap(s))
	t.Log(len(a), cap(a))
	t.Log(len(b), cap(b))
}
