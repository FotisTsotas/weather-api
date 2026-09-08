package schedulers

import (
	"log/slog"
	"weather-api/jobs"

	"github.com/robfig/cron/v3"
)

func InitScheduler() *cron.Cron {
	c := cron.New()
	for _, j := range jobs.GetJobs() {
		_, err := c.AddFunc(j.Schedule, j.Task)
		if err != nil {
			slog.Error("Αποτυχία εγγραφής cron job", "job", j.Name, "error", err)
			continue
		}
		slog.Info("Cron job εγγράφηκε επιτυχώς", "job", j.Name, "schedule", j.Schedule)
	}

	c.Start()

	return c
}
