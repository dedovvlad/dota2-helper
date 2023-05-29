package crone

import (
	"github.com/dedovvlad/dota2-helper/internal/common"
	scrappingService "github.com/dedovvlad/dota2-helper/internal/models/scrapping"
	scrappingStorage "github.com/dedovvlad/dota2-helper/internal/repositories/scrapping"
	"github.com/pkg/errors"
)

func (p *Processor) AddItems() error {
	items, err := p.service.ItemsList()
	if err != nil {
		return errors.Wrap(err, "getting items list")
	}

	err = p.storage.AddItems(
		common.TransformSlice(
			items,
			func(hero *scrappingService.Item) scrappingStorage.Item {
				return scrappingStorage.Item{
					ItemName: hero.Name,
				}
			},
		),
	)
	if err != nil {
		return errors.Wrap(err, "adding items list")
	}

	return nil
}
