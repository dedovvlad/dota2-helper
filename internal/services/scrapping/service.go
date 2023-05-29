package scrapping

import (
	"github.com/gocolly/colly"
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
