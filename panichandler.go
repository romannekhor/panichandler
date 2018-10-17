package panichandler

import (
	"github.com/valyala/fasthttp"
)

// PanicRecoveryHandler is an extension of standard fasthttp.RequestHandler.
//  `recoverDetails` contains all the details from 'recover()' call.
type PanicRecoveryHandler func(ctx *fasthttp.RequestCtx, recoverDetails interface{})

// PanicHandler wraps original request handler with custom panic recovery logic (`PanicRecoveryHandler`)
func PanicHandler(handler fasthttp.RequestHandler, onRecover PanicRecoveryHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if onRecover != nil {
			defer recoverIfNeeded(ctx, onRecover)
		}

		handler(ctx)
	}
}

func recoverIfNeeded(ctx *fasthttp.RequestCtx, onRecover PanicRecoveryHandler) {
	if r := recover(); r != nil {
		onRecover(ctx, r)
	}
}

// SimplePanicHandler covers the most simple case when FastHTTP server needs to return HTTP 500 on unhandled panics.
func SimplePanicHandler(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return PanicHandler(handler, func(ctx *fasthttp.RequestCtx, recoverDetails interface{}) {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBody([]byte("Unexpected error. Recovered from 'panic'.`"))
	})
}
