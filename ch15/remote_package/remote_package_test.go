package remote_package

import (
	cmap "github.com/orcaman/concurrent-map"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	m := cmap.New()
	m.Set("foo", "bar")
	if value, ok := m.Get("foo"); !ok {
		t.Log("Not Found!")
	} else {
		t.Log(value)
	}
}
