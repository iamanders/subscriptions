package app

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/iamanders/subscriptions/internal/data"
)

// GET /settings
func (app *App) handleSettingsIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		settings, err := data.SettingsFind()
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}
		app.respondJson(w, r, settings, http.StatusOK)
	}
}

// PUT /settings
func (app *App) handleSettingsPut() http.HandlerFunc {

	type request struct {
		Locale   string `json:"locale" validate:"required,min=1,max=10,oneof=sv-SE en-US"`
		Currency string `json:"currency" validate:"required,min=1,max=10,oneof=SEK EUR USD"`
		Decimals int8   `json:"decimals" validate:"number,min=0,max=2"`
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
		updatedSettings, err := data.SettingsUpdate(request.Locale, request.Currency, request.Decimals)
		if err != nil {
			app.respondJson(w, r, err, http.StatusInternalServerError)
			return
		}

		// Repsonse
		app.respondJson(w, r, updatedSettings, http.StatusOK)
	}
}
