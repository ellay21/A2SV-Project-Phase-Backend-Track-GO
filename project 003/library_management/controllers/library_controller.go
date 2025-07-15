package controllers

import (
	"fmt"
	"library_management/models"
	"library_management/services"
	"strconv"
)

type LibraryController struct {
	Service services.LibraryManager
}

func NewLibraryController(service services.LibraryManager) *LibraryController {
	return &LibraryController{
		Service: service,
	}
}

func (lc *LibraryController) AddBook() {
	fmt.Print("Enter book title: ")
	title := ""
	fmt.Scanln(&title)
	fmt.Print("Enter book author: ")
	author := ""
	fmt.Scanln(&author)

	book := models.Book{Title: title, Author: author}
	lc.Service.AddBook(book)
	fmt.Println("Book added successfully!")
}

func (lc *LibraryController) RemoveBook() {
	fmt.Print("Enter book ID to remove: ")
	var idStr string
	fmt.Scanln(&idStr)
	bookID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid book ID. Please enter a number.")
		return
	}

	err = lc.Service.RemoveBook(bookID)
	if err != nil {
		fmt.Printf("Error removing book: %v\n", err)
		return
	}
	fmt.Println("Book removed successfully!")
}

func (lc *LibraryController) BorrowBook() {
	fmt.Print("Enter book ID to borrow: ")
	var bookIDStr string
	fmt.Scanln(&bookIDStr)
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Println("Invalid book ID. Please enter a number.")
		return
	}

	fmt.Print("Enter member ID: ")
	var memberIDStr string
	fmt.Scanln(&memberIDStr)
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid member ID. Please enter a number.")
		return
	}

	err = lc.Service.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error borrowing book: %v\n", err)
		return
	}
	fmt.Println("Book borrowed successfully!")
}

func (lc *LibraryController) ReturnBook() {
	fmt.Print("Enter book ID to return: ")
	var bookIDStr string
	fmt.Scanln(&bookIDStr)
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		fmt.Println("Invalid book ID. Please enter a number.")
		return
	}

	fmt.Print("Enter member ID: ")
	var memberIDStr string
	fmt.Scanln(&memberIDStr)
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid member ID. Please enter a number.")
		return
	}

	err = lc.Service.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error returning book: %v\n", err)
		return
	}
	fmt.Println("Book returned successfully!")
}

func (lc *LibraryController) ListAvailableBooks() {
	books := lc.Service.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No books available in the library.")
		return
	}
	fmt.Println("\n--- Available Books ---")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
	fmt.Println("-----------------------")
}

func (lc *LibraryController) ListBorrowedBooks() {
	fmt.Print("Enter member ID to list borrowed books: ")
	var memberIDStr string
	fmt.Scanln(&memberIDStr)
	memberID, err := strconv.Atoi(memberIDStr)
	if err != nil {
		fmt.Println("Invalid member ID. Please enter a number.")
		return
	}

	books, err := lc.Service.ListBorrowedBooks(memberID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(books) == 0 {
		fmt.Printf("Member %d has not borrowed any books.\n", memberID)
		return
	}

	fmt.Printf("\n--- Books Borrowed by Member %d ---\n", memberID)
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
	fmt.Println("---------------------------------")
}

func (lc *LibraryController) AddMember() {
	fmt.Print("Enter member name: ")
	name := ""
	fmt.Scanln(&name)
	libraryService, ok := lc.Service.(*services.Library)
	if !ok {
		fmt.Println("Error: Cannot add member. Underlying service is not a Library.")
		return
	}
	member := libraryService.AddMember(name)
	fmt.Printf("Member added successfully! ID: %d, Name: %s\n", member.ID, member.Name)
}