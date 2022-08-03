package bug

type IBug interface {
	// Ensure if error is nil or performing an action otherwise,
	// example: IBug.ensue(file.Close())
	Ensure(error)
}
