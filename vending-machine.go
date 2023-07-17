package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Item struct {
	Name   string
	Price  float64
	Stock  int
	Amount int
}

type VendingMachine struct {
	Items       []Item
	TotalAmount float64
}

func NewVendingMachine() *VendingMachine {
	return &VendingMachine{}
}

func (vm *VendingMachine) AddItems(items ...Item) {
	vm.Items = append(vm.Items, items...)
}

func (vm *VendingMachine) displayItems() {
	fmt.Println("Available Items:")
	for i, item := range vm.Items {
		fmt.Printf("%d. %s - $%.2f (Stock: %d)\n", i+1, item.Name, item.Price, item.Stock)
	}
	fmt.Println()
}

func (vm *VendingMachine) selectItems() ([]int, []int, error) {
	fmt.Print("Enter item numbers separated by commas: ")
	var itemNumbersInput string
	_, err := fmt.Scanln(&itemNumbersInput)
	if err != nil {
		return nil, nil, err
	}
	itemNumbersInput = strings.TrimSpace(itemNumbersInput)
	itemNumbersStr := strings.Split(itemNumbersInput, ",")
	itemNumbers := make([]int, len(itemNumbersStr))
	for i, numStr := range itemNumbersStr {
		itemNumber, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil || itemNumber < 1 || itemNumber > len(vm.Items) {
			return nil, nil, fmt.Errorf("invalid item number: %s", numStr)
		}
		itemNumbers[i] = itemNumber
	}

	fmt.Print("Enter quantities separated by commas: ")
	var quantitiesInput string
	_, err = fmt.Scanln(&quantitiesInput)
	if err != nil {
		return nil, nil, err
	}
	quantitiesInput = strings.TrimSpace(quantitiesInput)
	quantityStr := strings.Split(quantitiesInput, ",")
	quantities := make([]int, len(quantityStr))
	for i, qtyStr := range quantityStr {
		quantity, err := strconv.Atoi(strings.TrimSpace(qtyStr))
		if err != nil || quantity < 1 {
			return nil, nil, fmt.Errorf("invalid quantity: %s", qtyStr)
		}
		quantities[i] = quantity
	}

	if len(itemNumbers) != len(quantities) {
		return nil, nil, fmt.Errorf("mismatched number of items and quantities")
	}

	return itemNumbers, quantities, nil
}

func (vm *VendingMachine) insertCoin() {
	fmt.Print("Insert a coin (in dollars): ")
	var input string
	_, _ = fmt.Scanln(&input)
	input = strings.TrimSpace(input)
	coin, err := strconv.ParseFloat(input, 64)
	if err != nil || coin <= 0 {
		fmt.Println("Invalid coin")
		return
	}
	vm.TotalAmount += coin
	fmt.Printf("Coin of $%.2f inserted successfully\n", coin)
}

func (vm *VendingMachine) purchaseItems(itemNumbers []int, quantities []int) {
	totalPrice := 0.0
	purchasedItems := make([]Item, 0, len(itemNumbers))
	for i, itemNumber := range itemNumbers {
		item := &vm.Items[itemNumber-1]
		if item.Stock == 0 {
			fmt.Printf("Item '%s' is out of stock. Skipping...\n", item.Name)
			continue
		}
		if quantities[i] > item.Stock {
			fmt.Printf("Insufficient stock for '%s'. Skipping...\n", item.Name)
			continue
		}
		totalPrice += item.Price * float64(quantities[i])
		purchasedItems = append(purchasedItems, *item)
	}
	if totalPrice > vm.TotalAmount {
		fmt.Println("Insufficient amount. Please insert more coins.")
		return
	}

	fmt.Println("You have purchased the following items:")
	for i, item := range purchasedItems {
		quantity := quantities[i]
		cost := item.Price * float64(quantity)
		fmt.Printf("- %d %s for $%.2f\n", quantity, item.Name, cost)
		item.Stock -= quantity
	}
	change := vm.TotalAmount - totalPrice
	if change > 0 {
		fmt.Printf("Please collect your change: $%.2f\n", change)
	}
	vm.TotalAmount = 0
}

func main() {
	vendingMachine := NewVendingMachine()

	vendingMachine.AddItems(
		Item{Name: "Coke", Price: 1.50, Stock: 3},
		Item{Name: "Chips", Price: 1.20, Stock: 5},
		Item{Name: "Water", Price: 0.80, Stock: 2},
	)

	for {
		vendingMachine.displayItems()
		itemNumbers, quantities, err := vendingMachine.selectItems()
		if err != nil {
			fmt.Println(err)
			continue
		}
		vendingMachine.insertCoin()
		vendingMachine.purchaseItems(itemNumbers, quantities)
		fmt.Println()
	}
}
