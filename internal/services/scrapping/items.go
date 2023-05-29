package scrapping

import (
	"github.com/dedovvlad/dota2-helper/internal/common"
	"github.com/dedovvlad/dota2-helper/internal/models/scrapping"
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
	"strings"
)

func (s *Service) ItemsList() ([]*scrapping.Item, error) {
	items := make([]*scrapping.Item, 0, ItemsListCapacity)
	s.collector.OnHTML(ItemsListSelector, func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", Href)
		items = append(items, &scrapping.Item{
			Name: strings.TrimSpace(e.Text),
			Link: link,
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
