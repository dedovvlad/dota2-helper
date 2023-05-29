package scrapping

import (
	"strings"

	"github.com/gocolly/colly"
	"github.com/pkg/errors"

	"github.com/dedovvlad/dota2-helper/internal/common"
	"github.com/dedovvlad/dota2-helper/internal/models/scrapping"
)

type Service struct {
	domain    string
	collector *colly.Collector
}

func NewService(domain string) *Service {
	newCollector := colly.NewCollector(
		colly.AllowedDomains(domain))
	newCollector.AllowURLRevisit = true

	return &Service{
		domain:    BaseProtocol + domain,
		collector: newCollector,
	}
}

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

func (s *Service) ItemsList() ([]*scrapping.Item, error) {
	items := make([]*scrapping.Item, 0, ItemsListCapacity)
	s.collector.OnHTML(ItemsListSelector, func(e *colly.HTMLElement) {
		items = append(items, &scrapping.Item{
			Name: strings.TrimSpace(e.Text),
		})
	})

	err := s.collector.Visit(s.domain + ItemsListPage)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"visiting page [%s]",
			ItemsListPage,
		)
	}

	if len(items) == 0 {
		return nil, errors.Wrapf(
			common.ErrListIsEmpty,
			"creating list for [%s]",
			ItemsListPage,
		)
	}

	return items, nil
}
