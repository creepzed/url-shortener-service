package repository

import "github.com/creepzed/url-shortener-service/app/key-generation-service/domain"

type RangeKeyRepository interface {
	GetRange() (*domain.RangeKey, error)
}

//go:generate mockery --case=snake --outpkg=rangekeymocks --output=../mocks/rangekeymocks --name=RangeKeyRepository
