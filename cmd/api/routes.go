package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router{
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/", app.homeHandler)
	router.HandlerFunc(http.MethodGet, "/health/view/:id", app.showHealthHandler)
	router.HandlerFunc(http.MethodPost, "/health/daily", app.createActivityHandler)
	router.HandlerFunc(http.MethodPut, "/health/view/:id", app.updateHealthHandler)
	router.HandlerFunc(http.MethodDelete, "/health/view/:id", app.deleteHealthHandler)

	return router
}