package gofunc_test

import (
	"fmt"
	"gofunc"
	"testing"
)

func hellomath(i, j int) {
	fmt.Printf("hello %d, %d, %d\n", i, j, i+j)
}

func TestF(t *testing.T) {
	fm, err := gofunc.Register("math", hellomath)
	if err != nil {
		t.Log(err)
	}
	err = gofunc.Call("math", 2, 3)
	if err != nil {
		t.Log(err)
	}
	fm.Call(2, 9)
	done, err := gofunc.CallWithDone("math", 2, 5)
	if err != nil {
		t.Log(err)
	}
	if done.Done() {
		t.Logf("%T, %[1]v", done)
	}
}
