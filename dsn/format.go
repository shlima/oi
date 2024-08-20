package dsn

// IPart interface to use with Format() parts
type IPart interface {
	isPart()
}

// Query formats query argument
type Query string
type scheme struct{}
type host struct{}
type path struct{}

var (
	// Scheme formats scheme
	Scheme = new(scheme)
	// Host formats host (with port)
	Host = new(host)
	// Path formats paths
	Path = new(path)
)

func (_ Query) isPart()  {}
func (_ scheme) isPart() {}
func (_ host) isPart()   {}
func (_ path) isPart()   {}
