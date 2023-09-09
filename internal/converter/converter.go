package converter

import (
	"app/internal/domain"
	"app/system"
	"context"
	"errors"
	"fmt"
)

type storageFrom interface {
	GetBooks(ctx context.Context, filter domain.BookFilter) []domain.Book
	BooksCount(ctx context.Context) int
}

type storageTo interface {
	NewBook(ctx context.Context, name string, URL string, loaded bool) (int, error)
	UpdateBookRate(ctx context.Context, id int, rate int) error
	UpdateBookPages(ctx context.Context, id int, pages []domain.Page) error
	UpdateAttributes(ctx context.Context, id int, attr domain.Attribute, data []string) error
}

type Builder struct {
	src storageFrom
	dst storageTo
}

func (b *Builder) WithFrom(src storageFrom) *Builder {
	b.src = src

	return b
}

func (b *Builder) WithTo(dst storageTo) *Builder {
	b.dst = dst

	return b
}

func (b *Builder) Convert(ctx context.Context, offset int, notUniqWorkaround bool) {
	books := b.src.GetBooks(ctx, domain.BookFilter{
		Limit:  b.src.BooksCount(ctx),
		Offset: offset,
	})

	for _, book := range books {
		system.Info(ctx, "Начат", book.ID)

		id, err := b.dst.NewBook(ctx, book.Data.Name, book.URL, book.Data.Parsed.Name)
		if err != nil {
			system.Error(ctx, err)

			system.Debug(ctx, book)

			if !notUniqWorkaround || !errors.Is(err, domain.BookAlreadyExistsError) {
				return
			}

			id, err = b.dst.NewBook(ctx, book.Data.Name, fmt.Sprintf("err (%d): %s", book.ID, book.URL), book.Data.Parsed.Name)
			if err != nil {
				system.Error(ctx, err)

				return
			}
		}

		if id != book.ID {
			system.Warning(ctx, fmt.Sprintf("ID %d изменился на %d", book.ID, id))
		}

		err = b.dst.UpdateBookPages(ctx, id, book.Pages)
		if err != nil {
			system.Error(ctx, err)

			return
		}

		err = b.dst.UpdateBookRate(ctx, id, book.Data.Rate)
		if err != nil {
			system.Error(ctx, err)

			return
		}

		for attr := range book.Data.Parsed.Attributes {
			err = b.dst.UpdateAttributes(ctx, id, attr, book.Data.Attributes[attr])
			if err != nil {
				system.Error(ctx, err)

				return
			}
		}
	}
}
