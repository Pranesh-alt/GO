package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

// User with many Roles (many-to-many)
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"unique"`
	Roles []Role `gorm:"many2many:user_roles;"`
}

// Role can belong to many Users
type Role struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:user_roles;"`
}

// Product with soft delete example
type Product struct {
	gorm.Model
	Name      string
	Price     float64
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func main() {
	dsn := "root:62145090@tcp(127.0.0.1:3306)/gorm_example?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Auto migrate
	if err := db.AutoMigrate(&User{}, &Role{}, &Product{}); err != nil {
		log.Fatal("auto migrate failed:", err)
	}

	// Create roles
	adminRole := Role{Name: "Admin"}
	userRole := Role{Name: "User"}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&adminRole)
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&userRole)

	// Create a user with roles (many-to-many)
	user := User{
		Name:  "Pranesh",
		Email: "pranesh@example.com",
		Roles: []Role{adminRole, userRole},
	}
	db.Create(&user)

	// Query user with roles (preload many-to-many)
	var fetchedUser User
	db.Preload("Roles").First(&fetchedUser, "email = ?", "pranesh@example.com")
	fmt.Println("User:", fetchedUser.Name)
	for _, r := range fetchedUser.Roles {
		fmt.Println("Role:", r.Name)
	}

	// Soft delete example with product
	product := Product{Name: "Tablet", Price: 299.99}
	db.Create(&product)

	// Soft delete product
	db.Delete(&product)
	fmt.Printf("Product %s soft deleted at %v\n", product.Name, time.Now())

	// Query only non-deleted products
	var products []Product
	db.Find(&products)
	fmt.Printf("Products (excluding soft deleted): %d\n", len(products))

	// Query including soft deleted
	var allProducts []Product
	db.Unscoped().Find(&allProducts)
	fmt.Printf("Products (including soft deleted): %d\n", len(allProducts))

	// Filtering, sorting, pagination example:
	var users []User
	pageSize := 2
	page := 1
	db.Preload("Roles").
		Where("name LIKE ?", "%a%").
		Order("created_at desc").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&users)

	fmt.Println("Filtered users:")
	for _, u := range users {
		fmt.Println("User:", u.Name)
	}

	// Hook example: before create callback (defined below)
	newUser := User{Name: "Saravanan", Email: "saravanan@example.com"}
	db.Create(&newUser)
}

// Example of GORM hook (callback) on User model
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate hook triggered for user:", u.Name)
	// You can add validations or modify fields here
	return nil
}
