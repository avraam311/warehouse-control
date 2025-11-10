package models

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
