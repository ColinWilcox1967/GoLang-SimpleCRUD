package main

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
      
)
var DB *gorm.DB

// InitializeDatabase opens a database connection and migrates the schema.
func initializeDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    // Automatically migrate the schema (only does this once).
    err = DB.AutoMigrate(&User{})
    if err != nil {
        log.Fatal("Failed to migrate schema:", err)
    }
    fmt.Println("Database connected!")
}

func main() {

	// Initialize database connection

	initializeDatabase()

	// Create a new user
	user, err := crud.CreateUser("John Doe", "john.doe@example.com")

	if err != nil {
		log.Fatal("Error creating user:", err)
	}

	fmt.Println("User created:", user)

	// Get all users
	users, err := crud.GetAllUsers()

	if err != nil {
		log.Fatal("Error fetching users:", err)
	}

	fmt.Println("All users:", users)

	// Update user information
	updatedUser, err := crud.UpdateUser(user.ID, "Johnathan Doe", "john.doe@newdomain.com")

	if err != nil {
		log.Fatal("Error updating user:", err)
	}

	fmt.Println("User updated:", updatedUser)

	// Delete a user
	err = crud.DeleteUser(user.ID)
	if err != nil {
		log.Fatal("Error deleting user:", err)
	}

	fmt.Println("User deleted successfully")
}