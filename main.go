package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/dop251/goja"
)

func main() {
	var wg sync.WaitGroup

	data, err := os.ReadFile("./index.js")

	if err != nil {
		panic(err)
	}

	script := string(data)

	ctx := goja.New()

	console := ctx.NewObject()

	errSet := console.Set("println", func(call goja.FunctionCall) goja.Value {
		for _, arg := range call.Arguments {
			fmt.Println(arg)
		}
		return goja.NaN()
	})
	if errSet != nil {
		panic(errSet)
	}

	if err := ctx.Set("console", console); err != nil {
		panic(err)
	}

	ctx.Set("setTimeoutCustom", func(call goja.FunctionCall) goja.Value {
		callable, isfunc := goja.AssertFunction(call.Arguments[0])

		if !isfunc {
			panic(ctx.ToValue("first argument must be a function"))
		}

		timeout := call.Arguments[1]
		time := goja.Value(timeout).ToInteger()

		wg.Add(1)
		go TimeoutThread(callable, time, &wg)

		return goja.Null()
	})
	defer wg.Wait()

	_, errScript := ctx.RunScript("", script)
	if errScript != nil {
		panic(errScript)
	}

}

func TimeoutThread(callable goja.Callable, timeValue int64, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(timeValue) * time.Millisecond)
	callable(goja.Undefined())
}
