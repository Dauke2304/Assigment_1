package library

import (
	"errors"
	"fmt"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	books map[string]Book
}

func NewLibrary() *Library {
	return &Library{
		books: make(map[string]Book),
	}
}

func (l *Library) AddBook(book Book) error {
	if _, exists := l.books[book.ID]; exists {
		return errors.New("book with this ID already exists")
	}
	l.books[book.ID] = book
	return nil
}

func (l *Library) BorrowBook(id string) error {
	book, exists := l.books[id]
	if !exists {
		return errors.New("book not found")
	}
	if book.IsBorrowed {
		return errors.New("book is already borrowed")
	}
	book.IsBorrowed = true
	l.books[id] = book
	return nil
}

func (l *Library) ReturnBook(id string) error {
	book, exists := l.books[id]
	if !exists {
		return errors.New("book not found")
	}
	if !book.IsBorrowed {
		return errors.New("book is not borrowed")
	}
	book.IsBorrowed = false
	l.books[id] = book
	return nil
}

func (l *Library) ListBooks() {
	fmt.Println("Available books:")
	for _, book := range l.books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}
