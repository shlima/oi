package format

import "time"

func ParseDateSoft(input string) (out time.Time) {
	if input == "" {
		return out
	}

	got, err := time.Parse(RFC3339date, input)
	if err != nil {
		return out
	}

	return got
}

func ParseTimeSoft(input string) (out time.Time) {
	if input == "" {
		return out
	}

	got, err := time.Parse(RFC3339time, input)
	if err != nil {
		return out
	}

	return got
}
