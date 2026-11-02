package main

import (
	"fmt"
	"sync"
)

type Products interface {
	Add()
	Delete()
	PrintAll()
	PrintByCategory()
}

type Product struct {
	name     string
	category string
}

type ProductList []Product

func (p ProductList) PrintByCategory(cat string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s :", cat)
	for _, i := range p {
		if i.category == cat {
			fmt.Printf("%s ", i.name)
		}
	}
	fmt.Println()
}

func main() {
	var wg sync.WaitGroup
	var allProducts = ProductList{{"Банан", "Фрукты"}, {"Помидор", "Овощи"}, {"Яблоко", "Фрукты"}, {"Груша", "Фрукты"}, {"Морковь", "Овощи"}}
	wg.Add(2)

	go allProducts.PrintByCategory("Фрукты", &wg)
	go allProducts.PrintByCategory("Овощи", &wg)

	wg.Wait()
}

// review
