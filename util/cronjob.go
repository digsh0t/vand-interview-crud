package util

import (
	"time"

	"github.com/robfig/cron/v3"
	"github.com/wintltr/vand-interview-crud-project/model"
)

func StartExportCSVCron(crontime string) {
	c := cron.New()
	c.AddFunc(crontime, func() {
		model.ExportAllProductToCSV("csv/products_" + time.Now().String())
	})
	c.Start()
}