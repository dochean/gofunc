package gofunc

import (
	"fmt"
	// "gofunc"
	"log"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup

func helloForTest(t int) {
	// fmt.Println("hello go.")
	// var c int
	defer wg.Done()
	fmt.Printf("hello[], %d\n", t)
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("end")
	// c++
}

func sumForBenchmark(i, j int) {
	_ = i + j
	// fmt.Println("res: ", k)
}

// for now assume that Call is not blocked with goroutine
func TestNewFunction(t *testing.T) {
	f, err := newFunction(helloForTest)
	if err != nil {
		t.Errorf("Err in newFunction %w", err)
	}
	fc := f.Clone()
	// fc, err := newFunction(helloForTest)
	log.Printf("%T, %[1]v\n", f)
	log.Printf("%T, %[1]v \n", fc)

	//maybe f/fc is a same function, which causes
	//Function.Call call the same blocked process
	wg.Add(2)
	go f.Call(3)
	go fc.fv.Call(fc.inValue)
	wg.Wait()
	// t.Errorf("%T, %[1]v\n", f)
}

func BenchmarkFunctionCall(b *testing.B) {
	f, _ := newFunction(sumForBenchmark)
	b.ResetTimer()

	// too spensive cost in value.call
	for i := 0; i < b.N; i++ {
		f.Call(i, i)
	}
}

func BenchmarkFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumForBenchmark(i, i)
	}
}
