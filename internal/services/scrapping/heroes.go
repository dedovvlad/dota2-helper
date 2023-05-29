package scrapping

import (
	"github.com/dedovvlad/dota2-helper/internal/common"
	"github.com/dedovvlad/dota2-helper/internal/models/scrapping"
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
	"strings"
)

func (s *Service) HeroesList() ([]*scrapping.Hero, error) {
	heroes := make([]*scrapping.Hero, 0, HeroesListCapacity)
	s.collector.OnHTML(HeroesListSelector, func(e *colly.HTMLElement) {
		heroes = append(heroes, &scrapping.Hero{
			Name: strings.TrimSpace(e.Text),
		})
	})

	err := s.collector.Visit(s.domain + HeroesListPage)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"visit page [%s]",
			HeroesListPage,
		)
	}

	if len(heroes) == 0 {
		return nil, errors.Wrapf(
			common.ErrListIsEmpty,
			"creating list for [%s]",
			HeroesListPage,
		)
	}

	return heroes, nil
}
