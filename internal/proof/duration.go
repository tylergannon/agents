package proof

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParsePeriod accepts Go durations plus day and week suffixes such as 3d and 2w.
func ParsePeriod(value string) (time.Duration, error) {
	value = strings.TrimSpace(value)
	if len(value) > 1 {
		unit := value[len(value)-1]
		if unit == 'd' || unit == 'w' {
			amount, err := strconv.ParseFloat(value[:len(value)-1], 64)
			if err != nil || amount <= 0 {
				return 0, fmt.Errorf("invalid period %q", value)
			}
			multiplier := 24 * time.Hour
			if unit == 'w' {
				multiplier *= 7
			}
			return time.Duration(amount * float64(multiplier)), nil
		}
	}
	duration, err := time.ParseDuration(value)
	if err != nil || duration <= 0 {
		return 0, fmt.Errorf("invalid period %q", value)
	}
	return duration, nil
}
