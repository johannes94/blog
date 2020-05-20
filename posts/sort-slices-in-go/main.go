package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Name  string
	Price float64
}

// Define a custom slice type
type ProductsByPrice []Product

// Implement sort.Interface
func (p ProductsByPrice) Len() int {
	return len(p)
}

// The Less function is just like the comparator function above
func (p ProductsByPrice) Less(i, j int) bool {
	return p[i].Price < p[j].Price
}

// Function two swap to elements in a slice
func (p ProductsByPrice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {

	// Sort for primitive slices
	int_slice := []int{3, 2, 4, 2}
	sort.Ints(int_slice)
	fmt.Println(int_slice)

	// Sort with comperator function
	products := []struct {
		Name  string
		Price float64
	}{
		{"Microphone", 10.0},
		{"Mouse", 20.0},
		{"Keyboard", 20.0},
		{"Headphone", 15.0},
	}

	// Sort by price ascending
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
	fmt.Println("Products sorted by price ascending:")
	fmt.Println(products)
	// [{Microphone 10} {Headphone 15} {Mouse 20} {Keyboard 20}]

	// Sort by price descending
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price > products[j].Price
	})
	fmt.Println("Products sorted by price descending:")
	fmt.Println(products)
	// [{Mouse 20} {Keyboard 20} {Headphone 15} {Microphone 10}]

	// sort with custom datastructure
	productsByPrice := ProductsByPrice{
		{"Microphone", 10.0},
		{"Mouse", 20.0},
		{"Keyboard", 20.0},
		{"Headphone", 15.0},
	}
	// Sort with custom datastructure and sort.Interface
	fmt.Println("Products sorted by price with custom datastructure:")
	sort.Sort(productsByPrice)
	fmt.Println(productsByPrice)
	// [{Microphone 10} {Headphone 15} {Mouse 20} {Keyboard 20}]

	// Binary search on sorted slices
	// For int, float64 and string
	float_slice := []float64{3.0, 2.0, 4.0, 2.0}
	sort.Float64s(float_slice)
	index := sort.SearchFloat64s(float_slice, 4)
	fmt.Printf("Number %v found at index: %v\n", 4, index)

	// With custom function

	search_products := []struct {
		Name  string
		Price int
	}{
		{"Microphone", 10},
		{"Mouse", 20},
		{"Keyboard", 20},
		{"Headphone", 15},
	}
	searchFor := 20

	// ascending sort and search
	sort.Slice(search_products, func(i, j int) bool {
		return search_products[i].Price < search_products[j].Price
	})
	fmt.Println(search_products)
	// [{Microphone 10} {Headphone 15} {Mouse 20} {Keyboard 20}]
	index = sort.Search(len(search_products), func(i int) bool {
		return search_products[i].Price >= searchFor
	})
	// Test for exact match
	if index < len(search_products) && search_products[index].Price == searchFor {
		fmt.Printf("Product with price %v found at index: %v\n", searchFor, index)
	} else {
		fmt.Printf("Value not found\n")
	}

	// descending sort and search
	sort.Slice(search_products, func(i, j int) bool {
		return search_products[i].Price > search_products[j].Price
	})
	fmt.Println(search_products)
	// [{Mouse 20} {Keyboard 20} {Headphone 15} {Microphone 10}]
	index = sort.Search(len(search_products), func(i int) bool {
		return search_products[i].Price <= searchFor
	})
	if index < len(search_products) && search_products[index].Price == searchFor {
		fmt.Printf("Product with price %v found at index: %v\n", searchFor, index)
	} else {
		fmt.Printf("Value not found\n")
	}
}
