package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ItemDTO struct {
	Name        string  `json:"name" db:"name" validate:"required"`
	Description string  `json:"description" db:"description" validate:"required"`
	Price       float64 `json:"price" db:"price" validate:"required"`
}

type ItemDB struct {
	ID          string  `json:"id" db:"id" validate:"required"`
	Name        string  `json:"name" db:"name" validate:"required"`
	Description string  `json:"description" db:"description" validate:"required"`
	Price       float64 `json:"price" db:"price" validate:"required"`
}

type UserWithRoleDTO struct {
	Email    string `json:"email" db:"email" validate:"required"`
	Password string `json:"password" db:"password" validate:"required"`
	Role     string `json:"role" db:"role" validate:"required"`
}

type UserDTO struct {
	Email    string `json:"email" db:"email" validate:"required"`
	Password string `json:"password" db:"password" validate:"required"`
}

type UserPermissionsDB struct {
	Routes []string `json:"routes" db:"routes" validate:"required"`
}

type UserWithHashDomain struct {
	Email string `json:"email" db:"email" validate:"required"`
	Hash  []byte `json:"hash" db:"hash" validate:"required"`
	Role  string `json:"role" db:"role" validate:"required"`
}

type UserWithHashDB struct {
	ID   uint   `json:"id" db:"id" validate:"required"`
	Hash []byte `json:"hash" db:"hash" validate:"required"`
	Role string `json:"role" db:"role" validate:"required"`
}

type Claims struct {
	Role   string `json:"role" validate:"required"`
	UserID uint   `json:"user_id" validate:"required"`
	jwt.RegisteredClaims
}

type HistoryItem struct {
	HistoryID      uint      `json:"history_id" db:"history_id"`
	ItemID         uint      `json:"item_id" db:"item_id"`
	Action         string    `json:"action" db:"action"`
	ChangedAt      time.Time `json:"changed_at" db:"changed_at"`
	ChangedBy      uint      `json:"changed_by" db:"changed_by"`
	ChangedByEmail string    `json:"changed_by_email" db:"email"`
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description" db:"description"`
	Price          float64   `json:"price" db:"price"`
}

type HistoryFilter struct {
	ItemID   *uint      `json:"item_id" form:"item_id"`
	DateFrom *time.Time `json:"date_from" form:"date_from"`
	DateTo   *time.Time `json:"date_to" form:"date_to"`
	UserID   *uint      `json:"user_id" form:"user_id"`
	Action   *string    `json:"action" form:"action"`
	Limit    *int       `json:"limit" form:"limit"`
	Offset   *int       `json:"offset" form:"offset"`
}

type VersionDiff struct {
	ItemID    uint     `json:"item_id"`
	Version1  uint     `json:"version1"`
	Version2  uint     `json:"version2"`
	NameDiff  *string  `json:"name_diff"` // nil если одинаковые
	DescDiff  *string  `json:"desc_diff"`
	PriceDiff *float64 `json:"price_diff"`
}
