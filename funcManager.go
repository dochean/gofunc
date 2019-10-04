package gofunc

import "fmt"

type FunctionManager struct {
	size        int64 // size of pool to call
	workersize  int64 // size of worker
	workerqueue chan []interface{}
	// len   int64 // len of pool used
	// index int64 //last index of pool
	// num   *Set  // unused index less than index
	// working *Set      //pool in use
	idle  chan int64 //pool idle
	f     *Function  // base function
	funcs map[int64]*Function
}

//FunctionResult cancel context support needed?
type FunctionResult struct {
	Err  error
	done chan struct{}
}

func (fr *FunctionResult) Done() bool {
	select {
	case <-fr.done:
		return true
	}
	return false
}

const (
	DefaultWorkerSize  = 20
	DefaultWorkerQueue = 1000
)

func NewFM(f interface{}) (*FunctionManager, error) {
	return NewFunctionManager(f, DefaultWorkerQueue, DefaultWorkerSize)
}

//Init do configiuration, more option needed
func NewFunctionManager(f interface{}, size int64, workersize int64) (*FunctionManager, error) {
	if size < DefaultWorkerQueue {
		size = DefaultWorkerQueue
	}
	workerqueue := make(chan []interface{}, size)
	if workersize < DefaultWorkerSize {
		workersize = DefaultWorkerSize
	}
	// fm.num = NewSet()
	idle := make(chan int64, workersize)

	ff, err := newFunction(f)
	if err != nil {
		return nil, fmt.Errorf("Err while function init: %w", err)
	}
	// fm.f = ff
	funcs := make(map[int64]*Function)
	for i := int64(0); i < workersize; i++ {
		funcs[i] = ff.Clone()
		idle <- i
	}
	fm := &FunctionManager{size, workersize, workerqueue, idle, ff, funcs}
	go func() {
		for i := range fm.idle {

			go func() {
				args := <-fm.workerqueue
				fm.WorkerCall(i, args...)
			}()
		}
	}()
	return fm, nil
}

func (fm *FunctionManager) Call(args ...interface{}) {
	fm.workerqueue <- args
}

func (fm *FunctionManager) WorkerCall(index int64, args ...interface{}) {
	fm.callWithDoneIndex(false, index, args...)
}

func (fm *FunctionManager) CallWithDone(args ...interface{}) *FunctionResult {
	i := <-fm.idle
	return fm.callWithDoneIndex(true, i, args...)
}

func (fm *FunctionManager) callWithDoneIndex(done bool, index int64, args ...interface{}) *FunctionResult {
	fr := &FunctionResult{nil, make(chan struct{})}

	go func() {
		defer func() {
			fm.idle <- index
		}()

		f, ok := fm.funcs[index]
		if !ok {
			f = fm.f.Clone()
			fm.funcs[index] = f
		}
		f.Call(args...)
		fr.Err = f.Err()
		close(fr.done)
	}()

	if done {
		return fr
	}
	return nil
}

//Close cancelContext/waitgroup needed
func (fm *FunctionManager) Close() {
	close(fm.workerqueue)
	close(fm.idle)
	fm.funcs = nil
	fm.f = nil
	fm = nil
}
