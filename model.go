package main

type Product struct {
	ID    int
	Name  string
	Price float64
}

var products = []Product{
	{ID: 1, Name: "Product A", Price: 10.99},
	{ID: 2, Name: "Product B", Price: 20.99},
	{ID: 3, Name: "Product C", Price: 15.49},
}
