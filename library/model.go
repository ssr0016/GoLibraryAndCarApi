package library

import (
	"main/pkg/api/errors"
	"time"
)

var (
	ErrNameEmpty         = errors.New("book.name-empty", "Name is empty")
	ErrBookNameExisting  = errors.New("book.name-existing", "Name already exists")
	ErrAuthorIDInvalid   = errors.New("book.author-id-invalid", "Author is invalid")
	ErrCategoryIDInvalid = errors.New("book.category-id-invalid", "Category is invalid")
	ErrBookNotFound      = errors.New("book.book-not-found", "Book not found")
)

type Book struct {
	ID          int64  `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	AuthorID    int64  `db:"author_id" json:"author_id"`
	CategoryID  int64  `db:"category_id" json:"category_id"`
	PublishedAt string `db:"published_at" json:"pubblished_at"`
}

type BookDTO struct {
	ID          int64       `db:"id" json:"id"`
	Title       string      `db:"title" json:"title"`
	AuthorID    int64       `db:"author_id" json:"author_id"`
	CategoryID  int64       `db:"category_id" json:"category_id"`
	PublishedAt string      `db:"published_at" json:"pubblished_at"`
	Author      []*Author   `json:"author"`
	Category    []*Category `json:"category"`
}

type Category struct {
	ID   int64  `db:"category_id" json:"category_id"`
	Name string `db:"name" json:"name"`
}

type CategoryDTO struct {
	ID    int64   `db:"category_id" json:"category_id"`
	Name  string  `db:"name" json:"name"`
	Books []*Book `json:"books"`
}

type Author struct {
	ID   int64  `db:"author_id" json:"author_id"`
	Name string `db:"name" json:"name"`
}

type AuthorDTO struct {
	ID    int64   `db:"author_id" json:"author_id"`
	Name  string  `db:"name" json:"name"`
	Books []*Book `json:"books"`
}

type CreateBookCommand struct {
	Title      string `json:"name"`
	AuthorID   int64  `json:"author_id"`
	CategoryID int64  `json:"category_id"`
}

type SearchBookQuery struct {
	Title      string     `schema:"title"`
	AuthorID   int64      `schema:"author_id"`
	CategoryID int64      `schema:"category_id"`
	DateFrom   *time.Time `schema:"date_from"`
	DateTo     *time.Time `schema:"date_to"`
	Page       int        `schema:"page"`
	PerPage    int        `schema:"page_size"`
}

type SearchBookResult struct {
	TotalCount int64   `json:"total_count"`
	Books      []*Book `json:"result"`
	Page       int64   `json:"page"`
	PerPage    int64   `json:"per_page"`
}

type UpdateBookCommand struct {
	ID         int64
	Title      string `json:"name"`
	AuthorID   int64  `json:"author_id"`
	CategoryID int64  `json:"category_id"`
}

func (cmd *CreateBookCommand) Validate() error {
	if len(cmd.Title) == 0 {
		return ErrNameEmpty
	}

	if cmd.AuthorID <= 0 {
		return ErrAuthorIDInvalid
	}

	if cmd.CategoryID <= 0 {
		return ErrCategoryIDInvalid
	}

	return nil
}

func (cmd *UpdateBookCommand) Validate() error {
	if len(cmd.Title) == 0 {
		return ErrNameEmpty
	}

	if cmd.AuthorID <= 0 {
		return ErrAuthorIDInvalid
	}

	if cmd.CategoryID <= 0 {
		return ErrCategoryIDInvalid
	}

	return nil
}

//Walang validation sa search
//Ang may validation lang is update and create
