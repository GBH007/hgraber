package agentserver

import (
	"app/internal/domain/hgraber"
	"context"
	"io"
)

type logger interface {
	Info(ctx context.Context, args ...any)
}

type storage interface {
	GetUnloadedBooks(ctx context.Context) []hgraber.Book

	UpdateBookPages(ctx context.Context, id int, pages []hgraber.Page) error
	UpdateBookName(ctx context.Context, id int, name string) error
	UpdateAttributes(ctx context.Context, id int, attr hgraber.Attribute, data []string) error

	GetUnsuccessPages(ctx context.Context) []hgraber.Page

	UpdatePageSuccess(ctx context.Context, id int, page int, success bool) error
	UpdatePage(ctx context.Context, id int, page int, success bool, url string) error

	GetBook(ctx context.Context, id int) (hgraber.Book, error)
}

type tempStorage interface {
	TryLockBookHandle(ctx context.Context, bookID int) bool
	UnLockBookHandle(ctx context.Context, bookID int)
	HasLockBookHandle(ctx context.Context, bookID int) bool

	TryLockPageHandle(ctx context.Context, bookID int, pageNumber int) bool
	UnLockPageHandle(ctx context.Context, bookID int, pageNumber int)
	HasLockPageHandle(ctx context.Context, bookID int, pageNumber int) bool
}

type files interface {
	CreatePageFile(ctx context.Context, id, page int, ext string, body io.Reader) error
}

type UseCase struct {
	logger      logger
	storage     storage
	tempStorage tempStorage
	files       files
}

func New(logger logger, storage storage, tempStorage tempStorage, files files) *UseCase {
	return &UseCase{
		logger:      logger,
		storage:     storage,
		tempStorage: tempStorage,
		files:       files,
	}
}
