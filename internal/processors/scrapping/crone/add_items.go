package crone

import (
	"context"

	"github.com/pkg/errors"

	"database/sql"
	"github.com/dedovvlad/dota2-helper/internal/common"
	scrappingService "github.com/dedovvlad/dota2-helper/internal/models/scrapping"
	scrappingStorage "github.com/dedovvlad/dota2-helper/internal/repositories/scrapping"
)

func (p *Processor) AddItemsList(ctx context.Context) error {
	total, err := p.storage.CountItems(ctx)
	if err != nil {
		return errors.Wrap(err, "getting total items")
	}

	/*
		TODO: Add logic for AddItems:
			проверка списка, при условии, если количество в бд и в пришедшем списке одинаковое
		  	- отсутствующие деактивировать
		  	- новые добавлять
	*/
	switch total {
	case 0:
		err = p.addItems()
		if err != nil {
			return errors.Wrap(err, "colling saving items list")
		}
	default:
		return nil
	}

	return nil
}

func (p *Processor) addItems() error {
	heroes, err := p.service.ItemsList()
	if err != nil {
		return errors.Wrap(err, "getting items list from scrapping")
	}

	err = p.storage.AddItems(
		common.TransformSlice(
			heroes,
			func(item *scrappingService.Item) scrappingStorage.Item {
				return scrappingStorage.Item{
					ItemName: item.Name,
					Link: sql.NullString{
						String: item.Link,
						Valid:  item.Link != "",
					},
				}
			},
		),
	)
	if err != nil {
		return errors.Wrap(err, "adding items list")
	}

	return nil
}
