package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Customer struct {
	ID      int
	Name    string
	Address string
	Phone   string
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "customers.db")
	if err != nil {
		return nil, err
	}

	// Create customers table if it doesn't exist
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS customers (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        address TEXT,
        phone TEXT
    );`
	if _, err := db.Exec(createTableQuery); err != nil {
		return nil, err
	}
	return db, nil
}

// Add a new customer
func addCustomer(db *sql.DB, name, address, phone string) {
	query := `INSERT INTO customers (name, address, phone) VALUES (?, ?, ?)`
	_, err := db.Exec(query, name, address, phone)
	if err != nil {
		fmt.Println("Error adding customer:", err)
		return
	}
	fmt.Println("Customer added successfully.")
}

// List all customers
func listCustomers(db *sql.DB) {
	rows, err := db.Query(`SELECT id, name, address, phone FROM customers`)
	if err != nil {
		fmt.Println("Error listing customers:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.Address, &c.Phone)
		if err != nil {
			fmt.Println("Error scanning customer:", err)
			continue
		}
		fmt.Printf("%d: %s, %s, %s\n", c.ID, c.Name, c.Address, c.Phone)
	}
}

// Edit a customer's details
func editCustomer(db *sql.DB, id int, newName, newAddress, newPhone string) {
	query := `UPDATE customers SET name = ?, address = ?, phone = ? WHERE id = ?`
	_, err := db.Exec(query, newName, newAddress, newPhone, id)
	if err != nil {
		fmt.Println("Error updating customer:", err)
		return
	}
	fmt.Println("Customer updated successfully.")
}

// Delete a customer by ID
func deleteCustomer(db *sql.DB, id int) {
	query := `DELETE FROM customers WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting customer:", err)
		return
	}
	fmt.Println("Customer deleted successfully.")
}

// main runs a customer management application that connects to a database and
// allows the user to perform operations such as adding, listing, editing, and
// deleting customer records. The user is prompted with a menu to select an
// operation, and the application continues running until the user chooses to
// exit. Database connection errors are handled and displayed to the user.
func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	defer db.Close()

	for {
		fmt.Println("\nCustomer Management:")
		fmt.Println("1. Add Customer")
		fmt.Println("2. List Customers")
		fmt.Println("3. Edit Customer")
		fmt.Println("4. Delete Customer")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, address, phone string
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter address: ")
			fmt.Scan(&address)
			fmt.Print("Enter phone: ")
			fmt.Scan(&phone)
			addCustomer(db, name, address, phone)
		case 2:
			listCustomers(db)
		case 3:
			var id int
			var newName, newAddress, newPhone string
			fmt.Print("Enter customer ID to edit: ")
			fmt.Scan(&id)
			fmt.Print("Enter new name: ")
			fmt.Scan(&newName)
			fmt.Print("Enter new address: ")
			fmt.Scan(&newAddress)
			fmt.Print("Enter new phone: ")
			fmt.Scan(&newPhone)
			editCustomer(db, id, newName, newAddress, newPhone)
		case 4:
			var id int
			fmt.Print("Enter customer ID to delete: ")
			fmt.Scan(&id)
			deleteCustomer(db, id)
		case 5:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
