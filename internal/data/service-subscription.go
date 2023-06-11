package data

import (
	"database/sql"
	"strings"
)

func SubscriptionFind(id string) (*Subscription, error) {
	var subscription Subscription
	result := DB.Where("id = ?", id).First(&subscription)
	if result.Error != nil {
		return nil, result.Error
	}
	return &subscription, nil
}

func SubscriptionCreate(title string, priceMonth uint, category string, paused bool) (*Subscription, error) {
	subscription := &Subscription{
		Title:      title,
		PriceMonth: priceMonth,
		Paused:     paused,
	}

	// Assign category?
	if category != "" {
		category, err := CategoryCreate(category)
		if err != nil {
			return nil, err
		}
		subscription.CategoryID = sql.NullString{String: category.ID, Valid: true}
	}

	// Create and return
	tx := DB.Create(&subscription)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return subscription, nil
}

func SubscriptionUpdate(id string, title string, priceMonth uint, category string, paused bool) (*Subscription, error) {
	subscription, err := SubscriptionFind(id)
	if err != nil {
		return nil, err
	}

	subscription.Title = title
	subscription.PriceMonth = priceMonth
	subscription.Paused = paused

	// Assign category?
	if category == "" {
		subscription.CategoryID = sql.NullString{Valid: false}
		subscription.Category = Category{}
	} else {
		category, err := CategoryCreate(category)
		if err != nil {
			return nil, err
		}
		subscription.CategoryID = sql.NullString{String: category.ID, Valid: true}
	}

	// Update and return
	tx := DB.Save(&subscription)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return subscription, nil
}

func SubscriptionDelete(id string) error {
	subscription, err := SubscriptionFind(id)
	if err != nil {
		return err
	}

	DB.Where("lower(id) = ?", strings.ToLower(subscription.ID)).Delete(&Subscription{})
	return nil
}
