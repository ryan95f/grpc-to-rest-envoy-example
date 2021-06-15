package v1

import (
	context "context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Using an array here so a database is not required.
// However, multiple instances of the book service will not
// share the same data.
var (
	books []*Book = []*Book{}
)

// BooksEndpoint represents an implementation of the gRPC books service.
type BooksEndpoint struct {
	UnsafeBookServiceServer
}

// ListBooks returns all books that are stored in the service.
func (be *BooksEndpoint) ListBooks(ctx context.Context, request *ListBooksRequest) (*ListBooksResponse, error) {
	return &ListBooksResponse{
		Books: books,
	}, nil
}

// GetBook returns a book with the given title.
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

// CreateBook creates a new book and adds it to the service.
func (be *BooksEndpoint) CreateBook(ctx context.Context, request *CreateBookRequest) (*CreateBookResponse, error) {
	newBook := new(Book)
	newBook.Title = request.Title

	books = append(books, newBook)
	return &CreateBookResponse{
		Book: newBook,
	}, nil
}

// UpdateBook will update an existing book with the given title.
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

// DeleteBook removes a book from the service with a given name.
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
