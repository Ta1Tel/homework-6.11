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

func main() {
	var wg sync.WaitGroup
	var allProducts = []Product{{"Банан", "Фрукты"}, {"Помидор", "Овощи"}, {"Яблоко", "Фрукты"}, {"Груша", "Фрукты"}, {"Морковь", "Овощи"}}

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Print("Фрукты:")
		for _, i := range allProducts {
			if i.category == "Фрукты" {
				fmt.Printf("%s ", i.name)
			}
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Print("Овощи:")
		for _, i := range allProducts {
			if i.category == "Овощи" {
				fmt.Printf("%s ", i.name)
			}
		}
	}()

	wg.Wait()
}
