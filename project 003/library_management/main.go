package main

import (
	"fmt"
	"library_management/controllers"
	"library_management/services"
	"os"
	"strconv"
)

func main() {
	libraryService := services.NewLibrary()
	controller := controllers.NewLibraryController(libraryService)

	// Add some initial data for testing
	libraryService.AddBook(models.Book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"})
	libraryService.AddBook(models.Book{Title: "1984", Author: "George Orwell"})
	libraryService.AddBook(models.Book{Title: "To Kill a Mockingbird", Author: "Harper Lee"})

	member1 := libraryService.AddMember("Alice Smith")
	member2 := libraryService.AddMember("Bob Johnson")

	fmt.Printf("Initial Library Data: Book IDs %v, Member IDs %v\n", libraryService.Books, libraryService.Members)
	fmt.Printf("Member 1 (ID: %d) Name: %s\n", member1.ID, member1.Name)
	fmt.Printf("Member 2 (ID: %d) Name: %s\n", member2.ID, member2.Name)

	for {
		fmt.Println("\n--- Library Management System ---")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Add Member (for testing/setup)")
		fmt.Println("8. Exit")
		fmt.Print("Enter your choice: ")

		var choiceStr string
		fmt.Scanln(&choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid choice. Please enter a number between 1 and 8.")
			continue
		}

		switch choice {
		case 1:
			controller.AddBook()
		case 2:
			controller.RemoveBook()
		case 3:
			controller.BorrowBook()
		case 4:
			controller.ReturnBook()
		case 5:
			controller.ListAvailableBooks()
		case 6:
			controller.ListBorrowedBooks()
		case 7:
			controller.AddMember()
		case 8:
			fmt.Println("Exiting Library Management System. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 8.")
		}
	}
}