package grasp_patterns

import (
	"fmt"
)

// InventoryManager handles inventory-related operations
type InventoryManager struct {
	inventory map[string]int
}

func NewInventoryManager() *InventoryManager {
	return &InventoryManager{
		inventory: make(map[string]int),
	}
}

func (im *InventoryManager) AddItem(item string, quantity int) {
	im.inventory[item] += quantity
	fmt.Printf("Added %d %s(s) to inventory\n", quantity, item)
}

func (im *InventoryManager) CheckStock(item string) int {
	stock := im.inventory[item]
	fmt.Printf("Stock of %s: %d\n", item, stock)
	return stock
}

// SalesManager handles sales-related operations
type SalesManager struct {
	sales map[string]int
}

func NewSalesManager() *SalesManager {
	return &SalesManager{
		sales: make(map[string]int),
	}
}

func (sm *SalesManager) RecordSale(item string, quantity int) {
	sm.sales[item] += quantity
	fmt.Printf("Recorded sale of %d %s(s)\n", quantity, item)
}

func main() {
	inventory := NewInventoryManager()
	sales := NewSalesManager()

	inventory.AddItem("Widget", 50)
	inventory.CheckStock("Widget")

	sales.RecordSale("Widget", 10)
	sales.RecordSale("Gadget", 5)
}
