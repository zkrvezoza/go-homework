package main

import "fmt"

type ProductID int
type Price int
type Months int

type Product struct {
	ID     ProductID
	Name   string
	Price  Price
	Months Months
}

func monthlyPayment(p Product) int {
	return int(p.Price) / int(p.Months)
}

func printProduct(p Product) {
	fmt.Println("===== Alif Shop =====")
	fmt.Println("ID:", p.ID)
	fmt.Println("Товар:", p.Name)
	fmt.Println("Цена:", p.Price, "сум")
	fmt.Println("Рассрочка:", p.Months, "месяцев")
	fmt.Println("В месяц:", monthlyPayment(p), "сум")
	fmt.Println("=====================")
}

func main() {
	product := Product{
		ID:     7291,
		Name:   "Xiaomi Redmi Note 14",
		Price:  2499000,
		Months: 12,
	}

	if product.Months < 3 || product.Months > 24 {
		fmt.Println("Ошибка: срок от 3 до 24 месяцев")
		return
	}

	printProduct(product)
}
