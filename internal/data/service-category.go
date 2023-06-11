package data

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// Create a category if title doesn't exist
// Returns the category
func CategoryCreate(title string) (*Category, error) {
	category := &Category{
		Title: title,
	}

	// Already in database?
	categoryDB, err := CategoryFindByTitle(title)
	fmt.Println(categoryDB)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err // Some error that is not ErrRecordNotFound
	} else if err == nil {
		return categoryDB, nil
	}

	// Create new and return
	tx := DB.Create(&category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return category, nil
}

func CategoryFind(id string) (*Category, error) {
	var category Category
	result := DB.Where("id = ?", id).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func CategoryFindByTitle(title string) (*Category, error) {
	var category Category
	result := DB.Where("lower(title) = ?", strings.ToLower(title)).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func CategoryUpdate(id string, title string) (*Category, error) {
	category, err := CategoryFind(id)
	if err != nil {
		return nil, err
	}

	category.Title = title

	// Update and return
	tx := DB.Save(&category)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return category, nil
}

func CategoryDelete(id string) error {
	category, err := CategoryFind(id)
	if err != nil {
		return err
	}

	// Delete all references
	DB.Model(&Subscription{}).Where("lower(category_id) = ?", category.ID).Update("category_id", nil)

	// Delete
	DB.Where("lower(id) = ?", strings.ToLower(category.ID)).Delete(&Category{})
	return nil
}
