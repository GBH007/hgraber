package mock

import (
	"app/internal/domain"
	"context"
)

type Database struct{}

func (*Database) GetPage(ctx context.Context, id int, page int) (*domain.PageFullInfo, error) {
	return nil, domain.PageNotFoundError
}

func (*Database) GetTitle(ctx context.Context, id int) (domain.Book, error) {
	return domain.Book{}, domain.BookNotFoundError
}

func (*Database) GetTitles(ctx context.Context, filter domain.BookFilter) []domain.Book {
	return nil
}

func (*Database) GetUnloadedTitles(ctx context.Context) []domain.Book {
	return nil
}

func (*Database) GetUnsuccessedPages(ctx context.Context) []domain.PageFullInfo {
	return nil
}

func (*Database) NewTitle(ctx context.Context, name string, URL string, loaded bool) (int, error) {
	return 0, domain.BookAlreadyExistsError
}

func (*Database) PagesCount(ctx context.Context) int {
	return 0
}

func (*Database) TitlesCount(ctx context.Context) int {
	return 0
}

func (*Database) UnloadedPagesCount(ctx context.Context) int {
	return 0
}

func (*Database) UnloadedTitlesCount(ctx context.Context) int {
	return 0
}

func (*Database) UpdateAttributes(ctx context.Context, id int, attr domain.Attribute, data []string) error {
	return domain.UnsupportedAttributeError
}

func (*Database) UpdatePageRate(ctx context.Context, id int, page int, rate int) error {
	return domain.BookNotFoundError
}

func (*Database) UpdatePageSuccess(ctx context.Context, id int, page int, success bool) error {
	return domain.PageNotFoundError
}

func (*Database) UpdateTitleName(ctx context.Context, id int, name string) error {
	return domain.BookNotFoundError
}

func (*Database) UpdateTitlePages(ctx context.Context, id int, pages []domain.Page) error {
	return domain.BookNotFoundError
}

func (*Database) UpdateTitleRate(ctx context.Context, id int, rate int) error {
	return domain.BookNotFoundError
}
