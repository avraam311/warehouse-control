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

type UserDTOWithRole struct {
	Email    string `json:"email" db:"email" validate:"required"`
	Password string `json:"password" db:"password" validate:"required"`
	Role     string `json:"role" db:"role" validate:"required"`
}

type UserDTO struct {
	Email    string `json:"email" db:"email" validate:"required"`
	Password string `json:"password" db:"password" validate:"required"`
}
