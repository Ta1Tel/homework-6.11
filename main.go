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

	work := func(cat string) {
		defer wg.Done()
		fmt.Printf("%s : ", cat)
		for _, i := range allProducts {
			if i.category == cat {
				fmt.Printf("%s ", i.name)
			}
		}
	}

	go work("Фрукты")
	go work("Овощи")

	wg.Wait()
}
