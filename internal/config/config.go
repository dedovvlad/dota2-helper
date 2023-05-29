package config

import (
	"github.com/pkg/errors"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	ScrappingDomain string
	StorageDSN      string

	AddHeroesSchedule string `envconfig:"default=* * * * ?"`
	AddItemsSchedule  string `envconfig:"default=* * * * ?"`
	Timezone          string `envconfig:"default=Europe/Lisbon"`
}

func InitConfig(prefix string) (*Config, error) {
	config := &Config{}
	if err := envconfig.InitWithPrefix(config, prefix); err != nil {
		return nil, errors.Wrap(err, "init config failed")
	}

	return config, nil
}
