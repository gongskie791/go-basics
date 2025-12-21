package utils

import (
	"ecommerce/models"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func DisplayProducts(products []models.Product) {
	fmt.Println("\n=== PRODUCTS ===")
	fmt.Printf("%-4s %-20s %-10s %-6s %s\n", "ID", "Name", "Price", "Stock", "Description")

	fmt.Println("----------------------------------------------------------------")

	for _, p := range products {
		fmt.Printf("%-4d %-20s $%-9.2f %-6d %s\n",
			p.ID, p.Name, p.Price, p.Stock, p.Description)
	}
}

func DisplayCart(items []models.Product, cart models.Cart, total float64) {
	fmt.Println("\n=== YOUR CART ===")
	if len(items) == 0 {
		fmt.Println("Cart is empty")
		return
	}

	fmt.Printf("%-20s %-10s %-8s %-10s\n", "Product", "Price", "Qty", "Subtotal")
	fmt.Println("--------------------------------------------------")

	for i, item := range items {
		cartItem := cart.Items[i]
		subtotal := item.Price * float64(cartItem.Quantity)
		fmt.Printf("%-20s $%-9.2f %-8d $%-9.2f\n",
			item.Name, item.Price, cartItem.Quantity, subtotal)
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("TOTAL: $%.2f\n", total)
}
