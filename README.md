> deprecated project (~~maybe~~)

GOFUNC
======================
function-oriented goroutine pool. too many feafures unimplemented.

## preface

goroutine pool project is generally a pool implementing goroutine size control and kernel schedueling.

and with go, in most cases you use it like
```go
pool.Init(option)
pool.Run(func(arg interface{}){
    s, _ := arg.(string)
    fmt.Println("runing ", s)
}("goroutine"))
```
or use variables in function closure

it works but a little hard to use, and then this idea arose.

## what can this do

two points:

1. function-oriented, more convient to use although some performance loss
2. observable result

something like this:
```go
fm, err := gofunc.Register("math", func(i, j int){
    fmt.Printf("result: %d", i+j)
})

//two type call
gofunc.Call("math", 2, 3)
fm.Call(2, 3)

//you can use original params instead of interface{} decoding
res := fm.CallWithDone(2, 3)
if res.Done(); res.Err!=nil {
    fmt.Printf("Err while Calling: %s", res.Err)
}
```

## what needed to do

too many.

but the most is to find a better way to do with func.

i want to implement function-oriented feature with `reflect.Value.Call` by caching a function instance in memory  to reduce function reuse's cost, but it cost too too much than i thought instead.

so it became a bad project.

## last

learn to rebuild the core and wait for some change in go.