package main

import (
	"github.com/sviterok/panichandler"
	"github.com/valyala/fasthttp"
)

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/hello":
			helloHandler(ctx)
		case "/danger":
			dangerousHandler(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	// Create a limiter struct.
	safeRequestHandler := panichandler.SimplePanicHandler(requestHandler)
	fasthttp.ListenAndServe(":4444", safeRequestHandler)

}

func helloHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Hello, World!"))
}

func dangerousHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("Doing science..."))
	panic("oh no!!!")
}
