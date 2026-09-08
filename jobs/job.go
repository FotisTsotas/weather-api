package jobs

import (
	"log/slog"
)

type Job struct {
	Name     string
	Schedule string
	Task     func()
}

func GetJobs() []Job {
	return []Job{
		{
			Name:     "Weather Update",
			Schedule: EveryMinute(),
			Task: func() {
				slog.Info("Cron Job: Ενημέρωση δεδομένων καιρού...")
			},
		},
		{
			Name:     "Daily Cleanup",
			Schedule: EveryThreeMinutes(), // Κάθε 3 λεπτά
			Task: func() {
				slog.Info("Cron Job: Εκκαθάριση προσωρινών δεδομένων...")
			},
		},
	}
}
