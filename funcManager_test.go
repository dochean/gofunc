package gofunc_test

import (
	"fmt"
	"gofunc"
	"testing"
)

func hello() {
	fmt.Println("hello")
}

func TestFM(t *testing.T) {
	fm, err := gofunc.NewFM(hello)
	if err != nil {
		t.Log(err)
	}
	fm.Call()
	fr := fm.CallWithDone()
	if fr.Done() {
		t.Logf("%T, %[1]v\n", fr)
	}
}
