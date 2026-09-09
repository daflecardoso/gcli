// Package version holds build metadata injected by GoReleaser via
// -ldflags at release time. Values stay at their defaults for
// `go build`/`go run` during local development.
package version

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

// String formats the build metadata for `gcli --version`.
func String() string {
	return Version + " (commit " + Commit + ", built " + Date + ")"
}
