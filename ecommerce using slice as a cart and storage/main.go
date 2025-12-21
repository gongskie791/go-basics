package main

import (
	"bufio"
	"ecommerce/services"
	"ecommerce/storage"
	"ecommerce/utils"
	"fmt"
	"os"
)

func displayProducts(productService *services.ProductService) {
	products := productService.ListProducts()
	utils.DisplayProducts(products)
}

func addToCart(scanner *bufio.Scanner, cartService *services.CartService, productService *services.ProductService) {
	products := productService.ListProducts()
	utils.DisplayProducts(products)

	if len(products) == 0 {
		fmt.Println("No products available")
		return
	}

	productID := utils.GetIntInput("\nEnter product ID: ")
	quantity := utils.GetIntInput("Enter quantity: ")

	_, message := cartService.AddItem(productID, quantity)
	fmt.Println(message)
}

func viewCart(cartService *services.CartService) {
	cart, total, items := cartService.GetCart()
	utils.DisplayCart(items, cart, total)
}

func checkout(cartService *services.CartService) {
	if utils.Confirm("Are you sure you want to checkout?") {
		_, message, total := cartService.Checkout()
		fmt.Printf("%s Total: $%.2f\n", message, total)
	}
}

func addNewProduct(scanner *bufio.Scanner, productService *services.ProductService) {
	fmt.Println("\n=== ADD NEW PRODUCT ===")

	name := utils.GetStringInput("Product name: ")
	description := utils.GetStringInput("Description: ")
	price := utils.GetFloatInput("Price: ")
	stock := utils.GetIntInput("Stock: ")

	product := productService.AddProduct(name, description, price, stock)
	fmt.Printf("Product added! ID: %d\n", product.ID)
}

func main() {
	productStore := storage.NewProductStore()
	cartStore := storage.NewCartStore()

	// Initialize services
	productService := services.NewProductService(productStore)
	cartService := services.NewCartService(cartStore, productStore)

	// Seed with sample data
	seedProducts(productService)

	// Main menu loop
	scanner := bufio.NewScanner(os.Stdin)

	for {
		utils.ClearScreen()
		fmt.Println("=== ECOMMERCE CONSOLE ===")
		fmt.Println("1. View Products")
		fmt.Println("2. Add Product to Cart")
		fmt.Println("3. View Cart")
		fmt.Println("4. Checkout")
		fmt.Println("5. Add New Product (Admin)")
		fmt.Println("6. Exit")
		fmt.Print("Select option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			displayProducts(productService)
		case "2":
			addToCart(scanner, cartService, productService)
		case "3":
			viewCart(cartService)
		case "4":
			checkout(cartService)
		case "5":
			addNewProduct(scanner, productService)
		case "6":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice!")
		}

		fmt.Print("\nPress Enter to continue...")
		scanner.Scan()
	}
}

func seedProducts(ps *services.ProductService) {
	ps.AddProduct("Laptop", "High-performance laptop", 999.99, 10)
	ps.AddProduct("Mouse", "Wireless mouse", 25.50, 50)
	ps.AddProduct("Keyboard", "Mechanical keyboard", 89.99, 30)
}
