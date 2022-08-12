package goid

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

func TestGoid(t *testing.T) {
	var data [256]byte

	id := Goid()
	a := fmt.Sprintf("goroutine %d ", id)
	b := data[:]
	b = b[:runtime.Stack(b, false)]

	if !strings.HasPrefix(string(b), a) {
		t.Errorf("runtime.Stack return %s, does contains goid %d", b, id)
	}
}

func ExampleGoid() {
	fmt.Println(Goid())
	// Output: 1
}
