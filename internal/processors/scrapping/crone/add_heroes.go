package crone

import (
	"context"

	"github.com/pkg/errors"

	"github.com/dedovvlad/dota2-helper/internal/common"
	scrappingService "github.com/dedovvlad/dota2-helper/internal/models/scrapping"
	scrappingStorage "github.com/dedovvlad/dota2-helper/internal/repositories/scrapping"
)

func (p *Processor) AddHeroesList(ctx context.Context) error {
	total, err := p.storage.CountHeroes(ctx)
	if err != nil {
		return errors.Wrap(err, "getting total heroes")
	}

	/*
		TODO: Add logic for AddHeroes:
			проверка списка, при условии, если количество в бд и в пришедшем списке одинаковое
		    - отсутствующие деактивировать
			- новые добавлять
	*/
	switch total {
	case 0:
		err = p.addHeroes()
		if err != nil {
			return errors.Wrap(err, "colling saving heroes list")
		}
	default:
		return nil
	}

	return nil
}

func (p *Processor) addHeroes() error {
	heroes, err := p.service.HeroesList()
	if err != nil {
		return errors.Wrap(err, "getting heroes list from scrapping")
	}

	err = p.storage.AddHeroes(
		common.TransformSlice(
			heroes,
			func(hero *scrappingService.Hero) scrappingStorage.Hero {
				return scrappingStorage.Hero{
					HeroName: hero.Name,
				}
			},
		),
	)
	if err != nil {
		return errors.Wrap(err, "adding heroes list")
	}

	return nil
}
