package models

import "shop/models"

type Product models.Product
type User models.User

type Shop interface {
	PrintProduct() []Product
	SendUsers() []User
}
