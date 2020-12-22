package dates

// DaysInWeek 一周有7天
const DaysInWeek = 7

// DaysToWeeks 日期在第几周
func DaysToWeeks(days int) float64 {
	return float64(days) / float64(DaysInWeek)
}

// WeeksToDays 周转日期
func WeeksToDays(week int) int {
	return week * DaysInWeek
}
