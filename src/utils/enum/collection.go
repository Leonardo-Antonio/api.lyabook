package enum

type collection struct {
	Users        string
	Sales        string
	Books        string
	Categories   string
	ShoppingCart string
}

var Collection = &collection{
	Users:        "users",
	Sales:        "sales",
	Books:        "books",
	Categories:   "categories",
	ShoppingCart: "shopping_cart",
}
