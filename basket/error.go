package basket

import "errors"

var (
	ErrBasketNotFound           = errors.New("basket not found")
	ErrRepositoryNotInitialised = errors.New("repository not initialised")
)
