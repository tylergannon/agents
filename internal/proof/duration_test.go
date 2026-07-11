package proof

import (
	"testing"
	"time"
)

func TestParsePeriod(t *testing.T) {
	t.Parallel()
	for input, want := range map[string]time.Duration{
		"30s": 30 * time.Second,
		"5m":  5 * time.Minute,
		"2d":  48 * time.Hour,
		"2w":  14 * 24 * time.Hour,
	} {
		got, err := ParsePeriod(input)
		if err != nil {
			t.Fatalf("ParsePeriod(%q): %v", input, err)
		}
		if got != want {
			t.Errorf("ParsePeriod(%q) = %v, want %v", input, got, want)
		}
	}
}

func TestParsePeriodRejectsInvalidValues(t *testing.T) {
	t.Parallel()
	for _, input := range []string{"", "0s", "-1m", "soon", "0w"} {
		if _, err := ParsePeriod(input); err == nil {
			t.Errorf("ParsePeriod(%q) succeeded", input)
		}
	}
}
