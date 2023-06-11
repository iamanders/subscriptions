package data

import (
	"errors"

	"gorm.io/gorm"
)

func initSettings() (*Settings, error) {
	var settings = &Settings{
		Locale:   "en-US",
		Currency: "USD",
		Decimals: 2,
	}
	result := DB.Save(settings)
	if result.Error != nil {
		return nil, result.Error
	}
	return settings, nil
}

// Gets the settings, if no settings exists it will be created
func SettingsFind() (*Settings, error) {
	var settings Settings
	result := DB.First(&settings)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		newSettings, _ := initSettings()
		return newSettings, nil
	} else if result.Error != nil {
		return nil, result.Error
	}
	return &settings, nil
}

func SettingsUpdate(locale string, currency string, decimals int8) (*Settings, error) {
	settings, err := SettingsFind()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// Continue
	} else if err != nil {
		return nil, err
	}

	settings.Locale = locale
	settings.Currency = currency
	settings.Decimals = decimals

	// Update and return
	tx := DB.Save(&settings)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return settings, nil
}
