package services

import (
	"fmt"
	"github.com/robfig/cron"
	"wim-api/repository"
)

func MergeAndInsertTraining() (string, error) {
	c := cron.New()
	c.AddFunc("@every 30m", repository.SelectMerging())
	c.Start()

}
