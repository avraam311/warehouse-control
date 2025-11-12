package models

import "github.com/golang-jwt/jwt/v5"

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
}

type Claims struct {
	UserID      uint
	Permissions []string
	jwt.RegisteredClaims
}
