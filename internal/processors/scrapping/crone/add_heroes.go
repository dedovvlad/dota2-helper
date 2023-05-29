package crone

import (
	"github.com/dedovvlad/dota2-helper/internal/common"
	scrappingService "github.com/dedovvlad/dota2-helper/internal/models/scrapping"
	scrappingStorage "github.com/dedovvlad/dota2-helper/internal/repositories/scrapping"
	"github.com/pkg/errors"
)

func (p *Processor) AddHeroes() error {
	heroes, err := p.service.HeroesList()
	if err != nil {
		return errors.Wrap(err, "getting heroes list")
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
