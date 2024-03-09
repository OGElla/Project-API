package main

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/OGElla/Project-API/internal/data"
	"github.com/OGElla/Project-API/internal/validator"
)

func (app *application) createActivityHandler(w http.ResponseWriter, r *http.Request){
	var input struct{
		Walking data.Walking `json:"walking"`
		Hydrate data.Hydrate `json:"hydrate"`
		Sleep data.Sleep `json:"sleep"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil{
		app.badRequestResponse(w, r, err)
		return
	}

	health := &data.Health{
		Walking: input.Walking,
		Hydrate: input.Hydrate,
		Sleep: input.Sleep,
	}

	//VALIDATION (99)
	v := validator.New()
	if data.ValidateDaily(v, health); !v.Valid(){
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Health.Insert(health)
	if err!=nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/health/view/%d", health.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"health": health}, headers)
	if err != nil{
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showHealthHandler(w http.ResponseWriter, r *http.Request) {
 	id, err := app.readIDParam(r)
	if err != nil{
		app.notFoundResponse(w, r)
		return
	}

	health, err := app.models.Health.Get(id)
	if err != nil{
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"healh":health}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateHealthHandler(w http.ResponseWriter, r *http.Request){
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return 
	}

	health, err := app.models.Health.Get(id)
	if err != nil{
		switch{
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default: 
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Walking int64 `json:"walking"`
		Hydrate int32 `json:"hydrate"`
		Sleep int32 `json:"sleep"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	health.Walking = data.Walking(input.Walking)
	health.Hydrate = data.Hydrate(input.Hydrate)
	health.Sleep = data.Sleep(input.Sleep)

	v := validator.New()
	if data.ValidateDaily(v, health); !v.Valid(){
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Health.Update(health)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"health": health}, nil)
	if err != nil{
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteHealthHandler(w http.ResponseWriter, r *http.Request){
	id, err := app.readIDParam(r)
	if err != nil{
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Health.Delete(id)
	if err != nil{
		switch{
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "healthtracker successfully deleted"}, nil)
	if err != nil{
		app.serverErrorResponse(w, r, err)
	}
}