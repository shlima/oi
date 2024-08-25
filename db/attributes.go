package db

// Attributes represents attributes to ise in INSERT/UPDATE
// statements
type Attributes map[string]any

// Add column and value
func (a Attributes) Add(key string, value any) Attributes {
	a[key] = value
	return a
}

// Columns returns list of columns
func (a Attributes) Columns() []string {
	out := make([]string, 0, len(a))
	for k := range a {
		out = append(out, k)
	}
	return out
}

// PickValues returns values by keys
func (a Attributes) PickValues(columns []string) []any {
	out := make([]any, len(columns))
	for i := range columns {
		out[i] = a[columns[i]]
	}
	return out
}
