package gofunc

import "fmt"

type FuncPool struct {
	pool map[string]*FunctionManager
}

var fp = &FuncPool{make(map[string]*FunctionManager)}

func Register(name string, f interface{}) (*FunctionManager, error) {
	return fp.Register(name, f)
}

func Call(name string, args ...interface{}) error {
	return fp.Call(name, args...)
}

func CallWithDone(name string, args ...interface{}) (*FunctionResult, error) {
	return fp.CallWithDone(name, args...)
}

func (fp *FuncPool) Register(name string, f interface{}) (*FunctionManager, error) {
	fn, err := NewFM(f)
	if err != nil {
		return nil, fmt.Errorf("FunctionPool creation err: %w", err)
	}
	fp.pool[name] = fn
	return fn, nil
}

func (fp *FuncPool) Call(name string, args ...interface{}) error {
	fn, ok := fp.pool[name]
	if !ok {
		return fmt.Errorf("FuncPool Err: no such function registered.")
	}
	fn.Call(args...)
	return nil
}

func (fp *FuncPool) CallWithDone(name string, args ...interface{}) (*FunctionResult, error) {
	fn, ok := fp.pool[name]
	if !ok {
		return nil, fmt.Errorf("FuncPool Err: no such function registered.")
	}
	fr := fn.CallWithDone(args...)
	return fr, nil
}
