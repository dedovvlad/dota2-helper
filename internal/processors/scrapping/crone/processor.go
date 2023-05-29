package crone

import (
	"context"

	scrappingService "github.com/dedovvlad/dota2-helper/internal/models/scrapping"
	scrappingStorage "github.com/dedovvlad/dota2-helper/internal/repositories/scrapping"
)

type (
	storage interface {
		AddHeroes(heroes []scrappingStorage.Hero) error
		AddItems(heroes []scrappingStorage.Item) error
		CountHeroes(ctx context.Context) (int64, error)
		CountItems(ctx context.Context) (int64, error)
	}
	service interface {
		HeroesList() ([]*scrappingService.Hero, error)
		ItemsList() ([]*scrappingService.Item, error)
	}
)

type Processor struct {
	storage storage
	service service
}

func NewProcessor(
	storage storage,
	service service,
) *Processor {
	return &Processor{
		storage: storage,
		service: service,
	}
}
