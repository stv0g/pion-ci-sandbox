// Package sandbox is a CI sandbox.
package sandbox

import (
	"errors"

	_ "github.com/pion/transport/v2/test" // nolint
)

// ErrDummy is a dummy.
var ErrDummy = errors.New("dummy")
