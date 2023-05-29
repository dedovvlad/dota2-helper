package scrapping

const BaseProtocol string = "https://"

// Heroes
const (
	HeroesListSelector string = "div.hero"
	HeroesListPage     string = "/heroes"
	HeroesListCapacity int    = 200
)

// Items
const (
	ItemsListSelector string = "td.cell-xlarge"
	ItemsListPage     string = "/items"
	ItemsListCapacity int    = 300
)
