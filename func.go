package gofunc

import (
	"fmt"
	"reflect"
	// "sync"
)

type Function struct {
	// rw      sync.RWMutex
	ft      reflect.Type
	fv      reflect.Value
	inNum   int
	outNum  int
	inType  []reflect.Type
	inValue []reflect.Value
	// outType []reflect.Type
	// rets    []reflect.Value
	err error
}

func (f *Function) Err() error {
	defer func() {
		f.err = nil
	}()
	return f.err
}

func (f *Function) Call(args ...interface{}) {
	// fmt.Println("call function")
	// error before call
	if len(args) != f.inNum {
		f.err = ErrParamNum
		return
	}
	defer func() {
		// TODO: panic handler for utility
		// should be nil after read
		// error in call
		if p := recover(); p != nil {
			f.err = fmt.Errorf("Err ocurred : %w", p)
		}
	}()
	// flag := true
	// first int

	//TODO: version control, function viewer will produce a proxy with version code which is key to read results.
	//TODO: param rollback
	//TODO: optimize param assignment
	// f.rw.Lock()
	// defer f.rw.Unlock()
	//TODO: more type supported after better way of reflect.Value.Call
	for i, _ := range args {
		switch f.inType[i].Kind() {
		case Int:
			v, ok := args[i].(int)
			if !ok {
				f.err = ErrType
				return
			}
			f.inValue[i].SetInt(int64(v))
		case Int8:
			v, ok := args[i].(int8)
			if !ok {
				f.err = ErrType
				return
			}
			f.inValue[i].SetInt(int64(v))
		case Int16:
			v, ok := args[i].(int16)
			if !ok {
				f.err = ErrType
				return
			}
			f.inValue[i].SetInt(int64(v))
		case Int32:
			v, ok := args[i].(int32)
			if !ok {
				f.err = ErrType
				return
			}
			f.inValue[i].SetInt(int64(v))
		case Int64:
			v, ok := args[i].(int64)
			if !ok {
				f.err = ErrType
				return
			}
			f.inValue[i].SetInt(int64(v))
		case Uint:
			v, ok := args[i].(uint)
			if !ok {
				f.err = ErrType
				return
			}
			f.inValue[i].SetUint(uint64(v))
		case Uint8:
			v, ok := args[i].(uint8)
			if !ok {
				f.err = ErrType
				return
			}
			f.inValue[i].SetUint(uint64(v))
		case Uint16:
			v, ok := args[i].(uint16)
			if !ok {
				f.err = ErrType
				return
			}
			f.inValue[i].SetUint(uint64(v))
		case Uint32:
			v, ok := args[i].(uint32)
			if !ok {
				f.err = ErrType
				return
			}
			f.inValue[i].SetUint(uint64(v))
		case Uint64:
			v, ok := args[i].(uint64)
			if !ok {
				f.err = ErrType
				return
			}
			f.inValue[i].SetUint(uint64(v))
		default:
			f.err = ErrType
			return
		}

	}
	f.fv.Call(f.inValue)
	// return nil
}

func (f *Function) Clone() *Function {
	cft, cfv := f.ft, f.fv
	cin, cout := f.inNum, f.outNum
	cinT := make([]reflect.Type, cin)
	copy(cinT, f.inType)
	inValue := make([]reflect.Value, cin)
	for i := 0; i < cin; i++ {
		inv := reflect.New(cinT[i]).Elem()
		inValue[i] = inv
	}

	return &Function{ft: cft, fv: cfv, inNum: cin, outNum: cout, inType: cinT, inValue: inValue, err: nil}
}

func newFunction(f interface{}) (*Function, error) {
	ft := reflect.TypeOf(f)
	if ft.Kind() != Func {
		return nil, fmt.Errorf("Func param needed but %s", ft.Kind())
	}

	fv := reflect.ValueOf(f)
	inNum, outNum := ft.NumIn(), ft.NumOut()
	if outNum != 0 {
		return nil, fmt.Errorf("Func cannot return result")
	}
	inType := make([]reflect.Type, inNum)
	// outType := make([]reflect.Type, outNum)
	for i := 0; i < inNum; i++ {
		inType[i] = ft.In(i)
	}
	// for i := 0; i < outNum; i++ {
	// 	outType[i] = ft.In(i)
	// }
	inValue := make([]reflect.Value, inNum)
	for i := 0; i < inNum; i++ {
		inv := reflect.New(inType[i]).Elem()
		inValue[i] = inv
	}

	return &Function{ft, fv, inNum, outNum, inType, inValue, nil}, nil
}
