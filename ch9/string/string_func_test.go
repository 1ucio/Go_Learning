package string_test

import (
	"testing"
	"strings"
	"strconv"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _,part := range parts {
		t.Logf("%[1]x", part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConvert(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log(s, len(s))
	if i,err := strconv.Atoi(s); err == nil {
		t.Log(10+i)
	}
}