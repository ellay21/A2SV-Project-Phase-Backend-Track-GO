package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
}

// Library implements the LibraryManager interface.
type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
	nextBookID int
	nextMemberID int
}

// NewLibrary creates and returns a new Library instance.
func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
		nextBookID: 1,
		nextMemberID: 1,
	}
}

// AddBook adds a new book to the library.
func (l *Library) AddBook(book models.Book) {
	book.ID = l.nextBookID
	book.Status = "Available"
	l.Books[book.ID] = book
	l.nextBookID++
}

func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.Books[bookID]; !exists {
		return errors.New("book not found")
	}
	if l.Books[bookID].Status == "Borrowed" {
		return errors.New("cannot remove a borrowed book, please return it first")
	}
	delete(l.Books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, bookExists := l.Books[bookID]
	if !bookExists {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book 

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, bookExists := l.Books[bookID]
	if !bookExists {
		return errors.New("book not found")
	}
	if book.Status == "Available" {
		return errors.New("book is not borrowed")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return errors.New("member not found")
	}

	foundAndRemoved := false
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			foundAndRemoved = true
			break
		}
	}

	if !foundAndRemoved {
		return errors.New("member did not borrow this book")
	}

	book.Status = "Available"
	l.Books[bookID] = book 
	l.Members[memberID] = member 

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member, exists := l.Members[memberID]
	if !exists {
		return nil, errors.New("member not found")
	}
	return member.BorrowedBooks, nil
}

func (l *Library) AddMember(name string) models.Member {
	member := models.Member{
		ID:   l.nextMemberID,
		Name: name,
		BorrowedBooks: []models.Book{},
	}
	l.Members[member.ID] = member
	l.nextMemberID++
	return member
}