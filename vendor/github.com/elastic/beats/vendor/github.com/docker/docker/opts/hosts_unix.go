// +build !windows

package opts // import "github.com/docker/docker/opts"

import "fmt"

// DefaultHost constant defines the default host string used by docker on other hosts than Windows
var DefaultHost = fmt.Sprintf("unix://%s", DefaultUnixSocket)
