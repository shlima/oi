package dsn

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDSN_Query(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		dsn := MustNew(t, "https://example.com:443/foo?a=1&b=2")
		require.Equal(t, "1", dsn.Query("a"))
		require.Equal(t, "2", dsn.Query("b"))
		require.Equal(t, "", dsn.Query("foo"))
	})
}

func TestDSN_QueryE(t *testing.T) {
	t.Parallel()

	type args struct {
		desc string
		q    string
		want string
		err  string
	}

	dsn := MustNew(t, "https://example.com:443/foo?a=1&b=2")

	tests := []args{
		{
			desc: "when a",
			q:    "a",
			want: "1",
		},
		{
			desc: "when b",
			q:    "b",
			want: "2",
		},
		{
			desc: "when not found",
			q:    "foo",
			want: "",
			err:  "no query named <foo> found in dsn",
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			got, err := dsn.QueryE(tt.q)
			switch {
			case tt.err != "":
				require.ErrorContains(t, err, tt.err)
			default:
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDSN_MustQuery(t *testing.T) {
	t.Parallel()

	type args struct {
		desc  string
		q     string
		want  string
		panic string
	}

	dsn := MustNew(t, "https://example.com:443/foo?a=1&b=2")

	tests := []args{
		{
			desc: "when a",
			q:    "a",
			want: "1",
		},
		{
			desc: "when b",
			q:    "b",
			want: "2",
		},
		{
			desc:  "when not found",
			q:     "foo",
			panic: "no query named <foo> found in dsn",
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()

			switch tt.panic {
			case "":
				require.Equal(t, tt.want, dsn.MustQuery(tt.q))
			default:
				require.PanicsWithError(t, tt.panic, func() { dsn.MustQuery(tt.q) })
			}
		})
	}
}

func TestDSN_Format(t *testing.T) {
	t.Parallel()

	dsn := MustNew(t, "https://example.com:443/foo?q=1&b=2")

	type args struct {
		desc  string
		parts []IPart
		want  string
	}

	tests := []args{
		{
			desc:  "when empty input",
			parts: nil,
			want:  "",
		},
		{
			desc:  "when scheme",
			parts: []IPart{Scheme},
			want:  "https:",
		},
		{
			desc:  "when host",
			parts: []IPart{Scheme, Host},
			want:  "https://example.com:443",
		},
		{
			desc:  "when path",
			parts: []IPart{Scheme, Host, Path},
			want:  "https://example.com:443/foo",
		},
		{
			desc:  "when one query",
			parts: []IPart{Scheme, Host, Path, Query("q")},
			want:  "https://example.com:443/foo?q=1",
		},
		{
			desc:  "when second query",
			parts: []IPart{Scheme, Host, Path, Query("b")},
			want:  "https://example.com:443/foo?b=2",
		},
		{
			desc:  "when query without a path",
			parts: []IPart{Scheme, Host, Query("q")},
			want:  "https://example.com:443?q=1",
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			got := dsn.Format(tt.parts...)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestParse(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		t.Parallel()
		got, err := Parse("ftp://")
		require.NoError(t, err)
		require.Equal(t, "ftp", got.Scheme())
	})

	t.Run("it errors", func(t *testing.T) {
		t.Parallel()
		_, err := Parse("\n")
		require.Error(t, err)
	})
}

func TestMustParse(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		t.Parallel()
		got := MustParse("ftp://")
		require.Equal(t, "ftp", got.Scheme())
	})

	t.Run("it panics", func(t *testing.T) {
		t.Parallel()
		require.Panics(t, func() { MustParse("\n") })
	})
}

func TestDSN_String(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		url := "http://example.com/a?b=1#foo"
		require.Equal(t, url, MustNew(t, url).String())
	})
}

func MustNew(t *testing.T, input string) *DSN {
	got, err := Parse(input)
	require.NoError(t, err)
	return got
}
