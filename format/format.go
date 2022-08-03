package format

import "time"

func Date(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format(RFC3339date)
}

func Time(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format(RFC3339time)
}
