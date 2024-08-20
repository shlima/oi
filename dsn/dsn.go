package dsn

import (
	"fmt"
	"net/url"
)

// DSN object representing parsed DSN string
type DSN struct {
	u     *url.URL
	query url.Values
}

// Parse returns DSN
func Parse(input string) (*DSN, error) {
	u, err := url.Parse(input)
	if err != nil {
		return nil, fmt.Errorf("failed to parse: %w", err)
	}

	return &DSN{u: u, query: u.Query()}, nil
}

// MustParse parses DSN string or panics
func MustParse(input string) *DSN {
	dsn, err := Parse(input)
	if err != nil {
		panic(fmt.Errorf("failed to parse DSN: %w", err))
	}

	return dsn
}

// Query returns query param
func (d *DSN) Query(name string) string {
	return d.query.Get(name)
}

// Scheme return scheme
func (d *DSN) Scheme() string {
	return d.u.Scheme
}

// QueryE returns query param or error
func (d *DSN) QueryE(name string) (string, error) {
	got := d.Query(name)
	switch got {
	case "":
		return "", fmt.Errorf("no query named <%s> found in dsn", name)
	default:
		return got, nil
	}
}

// MustQuery returns query param or panics
func (d *DSN) MustQuery(name string) string {
	got, err := d.QueryE(name)
	switch {
	case err != nil:
		panic(err)
	default:
		return got
	}
}

// String returns originally passed string
func (d *DSN) String() string {
	d.u.RawQuery = d.query.Encode()
	return d.u.String()
}

// Format formats DSN bby the specified format
func (d *DSN) Format(parts ...IPart) string {
	out := new(url.URL)
	query := out.Query()

	for i := range parts {
		switch v := parts[i].(type) {
		case *scheme:
			out.Scheme = d.u.Scheme
		case *host:
			out.Host = d.u.Host
		case *path:
			out.Path = d.u.Path
		case Query:
			query.Set(string(v), d.query.Get(string(v)))
		}
	}

	out.RawQuery = query.Encode()
	return out.String()
}
