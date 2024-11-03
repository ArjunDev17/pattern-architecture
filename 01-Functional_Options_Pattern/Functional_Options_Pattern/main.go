package main

import "fmt"

// Database struct with configuration fields
type Database struct {
	Host     string
	Port     int
	User     string
	Password string
}

// DatabaseOption Option type
type DatabaseOption func(*Database)

// WithHost Option function for Host
func WithHost(host string) DatabaseOption {
	return func(db *Database) {
		db.Host = host
	}
}

// WithPort Option function for Port
func WithPort(port int) DatabaseOption {
	return func(db *Database) {
		db.Port = port
	}
}

// WithUser Option function for User
func WithUser(user string) DatabaseOption {
	return func(db *Database) {
		db.User = user
	}
}

// WithPassword Option function for Password
func WithPassword(password string) DatabaseOption {
	return func(db *Database) {
		db.Password = password
	}
}

// NewDatabase Single Constructor with variadic options
func NewDatabase(options ...DatabaseOption) *Database {
	// Default configuration
	db := &Database{
		Host:     "localhost",
		Port:     5432,
		User:     "default_user",
		Password: "",
	}
	// Apply each option
	for _, option := range options {
		option(db)
	}
	return db
}

func main() {
	// Configure the database with only selected options
	db := NewDatabase(WithHost("localhost"), WithUser("admin"))
	fmt.Printf("Database connection: %+v\n", db)
}
