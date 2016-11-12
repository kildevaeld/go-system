package system

import (
	"fmt"
	"testing"
)

func TestUname(t *testing.T) {

	i, e := Uname()
	if e != nil {
		t.Fatal(e)
	}

	fmt.Printf("%#v\n", i)
}
