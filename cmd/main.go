package main

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/dedovvlad/dota2-helper/internal/config"
	scrappingProcCron "github.com/dedovvlad/dota2-helper/internal/processors/scrapping/crone"
	scrappingStg "github.com/dedovvlad/dota2-helper/internal/repositories/scrapping"
	scrappingSrv "github.com/dedovvlad/dota2-helper/internal/services/scrapping"
)

const (
	Fatal string = "[FATAL]:"
	Error string = "[ERROR]:"
)

func main() {
	err := app()
	if err != nil {
		log.Println(Fatal, "program ended with an accident [%s]", err)
	}
}

func app() error {
	cfg, err := config.InitConfig("")
	if err != nil {
		return err
	}

	// init dbs
	var db *sqlx.DB
	db, err = sqlx.Open("postgres", cfg.StorageDSN)
	if err != nil {
		return err
	}

	scrappingStorage := scrappingStg.NewStorage(db)

	// init services
	scrappingService := scrappingSrv.NewService(cfg.ScrappingDomain)

	// init processors
	scrappingProcessorCrone := scrappingProcCron.NewProcessor(
		scrappingStorage,
		scrappingService,
	)

	// init crones
	var cronSchedule *gocron.Scheduler
	{
		tz, err := time.LoadLocation(cfg.Timezone)
		if err != nil {
			log.Println(Error, err)
		}
		cronSchedule = gocron.NewScheduler(tz)
		cronSchedule.StartAsync()
		defer cronSchedule.Stop()
	}
	{
		AddHeroesTag := "AddHeroes"
		_, err = cronSchedule.Cron(cfg.AddHeroesSchedule).
			Tag(AddHeroesTag).
			SingletonMode().
			Do(func() error {
				err := scrappingProcessorCrone.AddHeroes()
				if err != nil {
					log.Println(Error, errors.Wrapf(err, "doing schedule [%s]", AddHeroesTag))
				}

				return nil
			})
		if err != nil {
			log.Println(Error, errors.Wrap(err, "configure scheduling failed"))
		}
		err = cronSchedule.RunByTag(AddHeroesTag)
		if err != nil {
			log.Println(Error, errors.Wrap(err, "first run scheduling failed"))
		}

		addItemsTag := "AddItems"
		_, err = cronSchedule.Cron(cfg.AddItemsSchedule).
			Tag(addItemsTag).
			SingletonMode().
			Do(func() error {
				err := scrappingProcessorCrone.AddItems()
				if err != nil {
					log.Println(Error, errors.Wrapf(err, "doing schedule [%s]", addItemsTag))
				}

				return nil
			})
		if err != nil {
			log.Println(Error, errors.Wrap(err, "configure scheduling failed"))
		}
		err = cronSchedule.RunByTag(addItemsTag)
		if err != nil {
			log.Println(Error, errors.Wrap(err, "first run scheduling failed"))
		}
	}

	// blocking run()
	{
		go func() {}()
		select {}
	}
}
