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
			Schedule: EveryFifteenMinutes(),
			Task: func() {
				HandleWeatherDataRetrieval()
			},
		},
	}
}
