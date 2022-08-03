package format

import (
	"time"

	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseDateSoft(t *testing.T) {
	type args struct {
		desc  string
		input string
		want  time.Time
	}

	tests := []args{
		{
			desc:  "when blank",
			input: "",
			want:  time.Time{},
		}, {
			desc:  "when invalid",
			input: "foo",
			want:  time.Time{},
		}, {
			desc:  "when valid",
			input: "2000-01-02",
			want:  time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := ParseDateSoft(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestParseTimeSoft(t *testing.T) {
	type args struct {
		desc  string
		input string
		want  time.Time
	}

	tests := []args{
		{
			desc:  "when blank",
			input: "",
			want:  time.Time{},
		}, {
			desc:  "when invalid",
			input: "foo",
			want:  time.Time{},
		}, {
			desc:  "when valid",
			input: "2000-01-02T03:04:05Z",
			want:  time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := ParseTimeSoft(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
