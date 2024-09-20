// main.go
package main

import (
	"assignment_2/ex_1"
	"assignment_2/ex_2"
	"assignment_2/ex_6"
	"log"
)

func main() {
	// Exercise 1
	//db, err := ex_1.ConnectDB()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//ex_1.CreateTable(db)
	//
	//ex_1.InsertUser(db, "Tima", 25)
	//ex_1.InsertUser(db, "Bota", 30)
	//
	//ex_1.GetUsers(db)
	//
	//err = db.Close()
	//if err != nil {
	//	return
	//}

	// Exercise 2
	//db2 := ex_2.ConnectDB()
	//
	//ex_2.AutoMigrate(db2)
	//
	//ex_2.InsertData(db2)
	//
	//ex_2.QueryData(db2)

	// ############# Exercise 3 ####################
	//r := ex_3.SetupRouter()
	//r.Run(":8080")

	// ############# Exercise 4 ####################
	//db := ex_2.ConnectDB()
	//
	//if err := ex_4.AutoMigrate(db); err != nil {
	//	log.Fatalf("Failed to migrate database: %v", err)
	//}
	//
	//// Insert users
	//users := []model.User2{
	//	{Name: "Name1", Age: 30},
	//	{Name: "Name2", Age: 25},
	//	{Name: "Name3", Age: 35},
	//}
	//if err := ex_4.InsertUsersWithTransaction(db, users); err != nil {
	//	log.Fatalf("Failed to insert users: %v", err)
	//}
	//
	//// Query users with pagination
	//page := 1
	//pageSize := 2
	//filteredUsers, err := ex_4.QueryUsersWithPagination(db, 0, page, pageSize)
	//if err != nil {
	//	log.Fatalf("Failed to query users: %v", err)
	//}
	//log.Printf("Users (Page %d): %+v\n", page, filteredUsers)
	//
	//// Update user
	//if err := ex_4.UpdateUser(db, 1, "Name Updated", 31); err != nil {
	//	log.Fatalf("Failed to update user: %v", err)
	//}
	//
	//// Delete user
	//if err := ex_4.DeleteUser(db, 2); err != nil {
	//	log.Fatalf("Failed to delete user: %v", err)
	//}

	// ############# Exercise 5 ####################
	//db := ex_5.ConnectDB()
	//
	//if err := model.AutoMigrate(db); err != nil {
	//	log.Fatalf("Failed to migrate database: %v", err)
	//}
	//
	//// Create and insert a customer with profile
	//customer := &model.Customer{
	//	Name: "Hello",
	//	Age:  25,
	//}
	//profile := &model.Profile{
	//	Bio:               "Software Engineer",
	//	ProfilePictureURL: "https://example.com/picture.jpg",
	//}
	//if err := ex_5.InsertCustomerWithProfile(db, customer, profile); err != nil {
	//	log.Fatalf("Failed to insert customer and profile: %v", err)
	//}
	//
	//// Query the customer with profile
	//queriedCustomer, err := ex_5.QueryCustomerWithProfile(db, customer.ID)
	//if err != nil {
	//	log.Fatalf("Failed to query customer: %v", err)
	//}
	//log.Printf("Queried Customer: %+v", queriedCustomer)
	//
	//// Update the customer's profile
	//if err := ex_5.UpdateCustomerProfile(db, customer.ID, "Lead Software Engineer", "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png"); err != nil {
	//	log.Fatalf("Failed to update profile: %v", err)
	//}
	//
	//// Delete the customer and profile
	//if err := ex_5.DeleteCustomerWithProfile(db, customer.ID); err != nil {
	//	log.Fatalf("Failed to delete customer and profile: %v", err)
	//}

	// ############# Exercise 6 ####################
	dbSQL, err := ex_1.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to SQL: %v", err)
	}

	dbGORM := ex_2.ConnectDB()

	// Set up routes for both SQL and GORM-based operations
	r := ex_6.SetupRoutesSQL(dbSQL)
	r.Run(":8080") // For SQL

	// For GORM
	r2 := ex_6.SetupRoutesGORM(dbGORM)
	r2.Run(":8081") // For GORM

}
