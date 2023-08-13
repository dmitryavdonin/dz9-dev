package book

import (
	"context"
)

type BookApi interface {
	GetBookPrice(ctx context.Context, bookId int) (bookPrice int, err error)
}
