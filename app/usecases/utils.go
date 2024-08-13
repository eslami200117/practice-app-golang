package usecases

import (
	"fmt"
	"time"
)

func timeAgo(t time.Time) string {
	duration := time.Since(t)

	
	if hours := duration.Hours(); hours >= 100 {
		return "last status in log time ago"
	} else if hours := duration.Hours(); hours >= 1 {
		minutes := int(duration.Minutes()) % 60
		return fmt.Sprintf("%d hour(s) and %d min(s) before", int(hours), minutes)
	} else if minutes := duration.Minutes(); minutes >= 1 {
		seconds := int(duration.Seconds()) % 60
		return fmt.Sprintf("%d min(s) and %d sec(s) before", int(minutes), seconds)
	} else {
		seconds := int(duration.Seconds())
		return fmt.Sprintf("%d second(s) before", seconds)
	}
}