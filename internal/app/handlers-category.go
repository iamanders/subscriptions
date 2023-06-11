package app

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/iamanders/subscriptions/internal/data"
	"gorm.io/gorm"
)

// GET /categories
func (app *App) handleCategoryIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var categories []data.Category
		data.DB.Order("lower(title) ASC").Find(&categories)
		app.respondJson(w, r, categories, http.StatusOK)
	}
}

// GET /categories/{uuid}
func (app *App) handleCategoryGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category, err := data.CategoryFind(chi.URLParam(r, "id"))
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			app.respondJson(w, r, "not found", http.StatusNotFound)
		} else if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
		}

		app.respondJson(w, r, category, http.StatusOK)
	}
}

// POST /categories
func (app *App) handleCategoryCreate() http.HandlerFunc {

	type request struct {
		Title string `json:"title" validate:"required,min=1,max=100"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Decode
		var request request
		err := app.decodeJson(w, r, &request)
		if err != nil {
			app.ErrorLog.Println(err)
		}

		// Validate
		validate := validator.New()
		err = validate.Struct(request)
		if err != nil {
			validationErrors := app.extractValidationErrors(err.(validator.ValidationErrors), request)
			app.respondJson(w, r, validationErrors, http.StatusUnprocessableEntity)
			return
		}

		// Insert
		category, err := data.CategoryCreate(request.Title)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}

		// Repsonse
		app.respondJson(w, r, category, http.StatusOK)
	}
}

// PUT /category/{uuid}
func (app *App) handleCategoryPut() http.HandlerFunc {

	type request struct {
		Title string `json:"title" validate:"required,min=1,max=100"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Decode
		var request request
		err := app.decodeJson(w, r, &request)
		if err != nil {
			app.ErrorLog.Println(err)
		}

		// Validate
		validate := validator.New()
		err = validate.Struct(request)
		if err != nil {
			validationErrors := app.extractValidationErrors(err.(validator.ValidationErrors), request)
			app.respondJson(w, r, validationErrors, http.StatusUnprocessableEntity)
			return
		}

		// Update
		category, err := data.CategoryUpdate(chi.URLParam(r, "id"), request.Title)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}

		// Repsonse
		app.respondJson(w, r, category, http.StatusOK)
	}
}

// DELETE /categories/{uuid}
func (app *App) handleCategoryDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category, err := data.CategoryFind(chi.URLParam(r, "id"))
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			app.respondJson(w, r, "not found", http.StatusNotFound)
		} else if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
		}

		err = data.CategoryDelete(category.ID)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
		}

		app.respondJson(w, r, true, http.StatusOK)
	}
}
