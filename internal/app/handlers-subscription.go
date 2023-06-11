package app

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/iamanders/subscriptions/internal/data"
	"gorm.io/gorm"
)

// GET /subscriptions
func (app *App) handleSubscriptionIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var subscriptions []data.Subscription
		data.DB.Preload("Category").Order("lower(title) ASC").Find(&subscriptions)
		app.respondJson(w, r, subscriptions, http.StatusOK)
	}
}

// GET /subscriptions/{uuid}
func (app *App) handleSubscriptionGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		var subscription data.Subscription
		// data.DB.Preload("Category").First(&subscription)
		result := data.DB.Where("id = ?", id).Preload("Category").First(&subscription)
		if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
			app.respondJson(w, r, "not found", http.StatusNotFound)
		} else if result.Error != nil {
			app.respondJson(w, r, result.Error, http.StatusInternalServerError)
		}

		app.respondJson(w, r, subscription, http.StatusOK)
	}
}

// POST /subscriptions
func (app *App) handleSubscriptionCreate() http.HandlerFunc {

	type request struct {
		Title      string `json:"title" validate:"required,min=1,max=100"`
		PriceMonth uint   `json:"price_month" validate:"required,numeric,min=1,max=9999999"`
		Category   string `json:"category" validate:"omitempty,min=1,max=100"`
		Paused     bool   `json:"paused" validate:"omitempty,boolean"`
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
		subscription, err := data.SubscriptionCreate(request.Title, request.PriceMonth, request.Category, request.Paused)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}

		// Repsonse
		app.respondJson(w, r, subscription, http.StatusOK)
	}
}

// PUT /subscriptions/{uuid}
func (app *App) handleSubscriptionPut() http.HandlerFunc {

	type request struct {
		Title      string `json:"title" validate:"required,min=1,max=100"`
		PriceMonth uint   `json:"price_month" validate:"required,numeric,min=1,max=9999999"`
		Category   string `json:"category" validate:"omitempty,min=1,max=100"`
		Paused     bool   `json:"paused" validate:"omitempty,boolean"`
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
		subscription, err := data.SubscriptionUpdate(chi.URLParam(r, "id"), request.Title, request.PriceMonth, request.Category, request.Paused)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}

		// Repsonse
		app.respondJson(w, r, subscription, http.StatusOK)
	}
}

// DELETE /subscriptions/{uuid}
func (app *App) handleSubscriptionDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subscription, err := data.SubscriptionFind(chi.URLParam(r, "id"))
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			app.respondJson(w, r, "not found", http.StatusNotFound)
		} else if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
		}

		err = data.SubscriptionDelete(subscription.ID)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
		}

		app.respondJson(w, r, true, http.StatusOK)
	}
}
