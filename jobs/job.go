package jobs

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
				HandleWeatherDataRetrieval()
			},
		},
	}
}
