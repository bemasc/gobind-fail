package beta

import (
	"github.com/bemasc/gobind-fail/alpha"
)

type Beta interface {
	alpha.Alpha
}

func F(b Beta) {
}
