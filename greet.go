// Package shared is a standalone module imported by project1 and project2.
// It is not part of a monorepo: each project has its own go.mod.
package shared

// Info is the payload both Gin services return from the shared module.
type Info struct {
	Service string `json:"service"`
	Message string `json:"message"`
}

// Describe builds a response both projects can return without copying the logic.
func Describe(service string) Info {
	return Info{
		Service: service,
		Message: "hello from the shared module",
	}
}
