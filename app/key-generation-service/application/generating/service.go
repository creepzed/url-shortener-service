package generating

import (
	"github.com/creepzed/url-shortener-service/app/key-generation-service/domain"
	"github.com/creepzed/url-shortener-service/app/key-generation-service/domain/repository"
)

type KeyGenerateService interface {
	GetKey() (string, error)
}

//go:generate mockery --case=snake --outpkg=servicemocks --output=../mocks/servicemocks --name=KeyGenerateService

type keyGenerateService struct {
	repository repository.RangeKeyRepository
	rangeKey   *domain.RangeKey
}

func NewKeyGenerateService(repository repository.RangeKeyRepository) (*keyGenerateService, error) {
	rangeKey, err := repository.GetRange()
	if err != nil {
		return nil, err
	}
	keyGenerateService := &keyGenerateService{
		repository: repository,
		rangeKey:   rangeKey,
	}
	return keyGenerateService, nil
}

func (k *keyGenerateService) GetKey() (string, error) {
	return k.rangeKey.GetKey()
}
