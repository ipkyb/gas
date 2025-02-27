package internal

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/ipkyb/gas/api"
)

func HandlerRecovery() api.Handler {
	coroutine := func(ctx api.Context) {
		catch := recover()
		if catch == nil {
			return
		}

		var fname string
		pc, file, line, ok := runtime.Caller(3)
		if ok {
			fname = runtime.FuncForPC(pc).Name()
		}

		trace := string(debug.Stack())
		msg := fmt.Sprintf(
			"Runtime recovers from panic: %v. Source: %s:%d: %s"+
				"\nStack Trace: %s",
			catch, file, line, fname, trace,
		)
		fmt.Println(msg)

		if ctx != nil {
			//ctx.Response().Reset()
			ctx.Status(api.StatusInternalServerError)
		}
	}

	return func(ctx api.Context) {
		defer coroutine(ctx)
		ctx.Next()
	}
}
