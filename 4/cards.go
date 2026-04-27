package main

import "fmt"

func main() {

	name := "iPhone 15 Pro"
	brand := "Apple"
	price := int64(12990000)
	inStock := true

	monthly := price / 12

	fmt.Printf("===== Alifshop =====\n")
	fmt.Printf("Товар: %s\n", name)
	fmt.Printf("Бренд: %s\n", brand)
	fmt.Printf("Цена: %d сум\n", price)
	fmt.Printf("В наличии: %t\n", inStock)
	fmt.Printf("Рассрочка: 12 мес → %d сум/мес\n", monthly)
	fmt.Printf("====================\n")
}
