package v1

import (
	context "context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	books []*Book = []*Book{}
)

type BooksEndpoint struct {
	UnsafeBookServiceServer
}

func (be *BooksEndpoint) ListBooks(ctx context.Context, request *ListBooksRequest) (*ListBooksResponse, error) {
	return &ListBooksResponse{
		Books: books,
	}, nil
}

// Returns a book with the given title.
func (be *BooksEndpoint) GetBook(ctx context.Context, request *GetBookRequest) (*GetBookResponse, error) {
	title := request.Title
	for _, book := range books {
		if book.Title == title {
			return &GetBookResponse{
				Book: book,
			}, nil
		}
	}
	return &GetBookResponse{},
		status.Errorf(codes.NotFound, "book not found with title %s", title)
}

// Creates a new book. If the book already exists then it will return an 'ALREADY_EXISTS' error.
func (be *BooksEndpoint) CreateBook(ctx context.Context, request *CreateBookRequest) (*CreateBookResponse, error) {
	newBook := new(Book)
	newBook.Title = request.Title

	books = append(books, newBook)
	return &CreateBookResponse{
		Book: newBook,
	}, nil
}

// Updates an existing book with the given title.
func (be *BooksEndpoint) UpdateBook(ctx context.Context, request *UpdateBookRequest) (*UpdateBookResponse, error) {
	title := request.Title
	updatedBook := request.Book

	for _, book := range books {
		if book.Title == title {
			book.Title = updatedBook.Title
			return &UpdateBookResponse{
				Book: updatedBook,
			}, nil
		}
	}

	return &UpdateBookResponse{},
		status.Errorf(codes.NotFound, "book not found with title %s", title)
}

// Deletes a book with the given name.
func (be *BooksEndpoint) DeleteBook(ctx context.Context, reqest *DeleteBookRequest) (*DeleteBookResponse, error) {
	title := reqest.Title
	for idx, book := range books {
		if book.Title == title {
			books = append(books[:idx], books[idx+1:]...)
			return &DeleteBookResponse{
				Done: true,
			}, nil
		}
	}
	return &DeleteBookResponse{},
		status.Errorf(codes.NotFound, "book not found with title %s", title)
}
