package jobs

func EveryMinute() string {
	return "* * * * *"
}

func EveryThreeMinutes() string {
	return "*/3 * * * *"
}

func EveryFiveMinutes() string {
	return "*/5 * * * *"
}

func EveryTenMinutes() string {
	return "*/10 * * * *"
}

func EveryFifteenMinutes() string {
	return "*/15 * * * *"
}

func EveryThirtyMinutes() string {
	return "*/30 * * * *"
}

func EveryHour() string {
	return "0 * * * *"
}

func EveryDay() string {
	return "0 0 * * *"
}

func EveryWeek() string {
	return "0 0 * * 0"
}

func EveryMonth() string {
	return "0 0 1 * *"
}
