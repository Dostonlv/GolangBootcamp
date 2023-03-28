package models

type Product struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID string `json:"category_id"`
}

type Products struct {
	Products []Product
}

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type Users struct {
	Users []User
}

type Benefit struct {
	Name    string `json:"name"`
	Benefit int    `json:"benefit"`
}
