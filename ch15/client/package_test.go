package client

import (
	"Go_Basic/src/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	if value, err := series.GetFibonacci(5); err != nil {
		t.Error(err)
	} else {
		t.Log(value)
	}
}
