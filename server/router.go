package server

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

func (app *App) Get(
	route string,
	handler func(ctx *Context),
) {

	app.mux.HandleFunc(
		"GET "+route,
		func(w http.ResponseWriter, r *http.Request,
		) {
			handler(&Context{
				RWriter: w,
				Request: r,
				Ctx:     r.Context(),
			})
		})

	app.handlerCounter++
}

func (app *App) Post(
	route string,
	handler func(ctx *Context),
) {
	app.mux.HandleFunc(
		"POST "+route,
		func(w http.ResponseWriter, r *http.Request,
		) {
			handler(&Context{
				RWriter: w,
				Request: r,
				Ctx:     r.Context(),
			})
		})

	app.handlerCounter++
}

func (app *App) Put(
	route string,
	handler func(ctx *Context),
) {
	app.mux.HandleFunc(
		"PUT "+route,
		func(w http.ResponseWriter, r *http.Request,
		) {
			handler(&Context{
				RWriter: w,
				Request: r,
				Ctx:     r.Context(),
			})
		})

	app.handlerCounter++
}

func (app *App) Delete(
	route string,
	handler func(ctx *Context),
) {
	app.mux.HandleFunc(
		"DELETE "+route,
		func(w http.ResponseWriter, r *http.Request,
		) {
			handler(&Context{
				RWriter: w,
				Request: r,
				Ctx:     r.Context(),
			})
		})

	app.handlerCounter++
}
