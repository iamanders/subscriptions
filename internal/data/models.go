package data

import (
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string         `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Set UUID when inserting a BaseModel model
func (model *BaseModel) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", strings.ToLower(uuid.New().String()))
	return nil
}

type Category struct {
	BaseModel
	Title         string         `json:"title"`
	Subscriptions []Subscription `json:"-"`
}

type Subscription struct {
	// Category   Category       `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BaseModel
	CategoryID sql.NullString `json:"-"`
	Category   Category       `json:"category"`
	Title      string         `json:"title"`
	PriceMonth uint           `json:"price_month"`
	Paused     bool           `json:"paused"`
	Comment    string         `json:"comment"`
}

type Settings struct {
	BaseModel
	Locale   string `json:"locale"`
	Currency string `json:"currency"`
	Decimals int8   `json:"decimals"`
}
