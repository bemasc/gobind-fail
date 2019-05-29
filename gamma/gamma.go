package gamma

import (
	"github.com/bemasc/gobind-fail/beta"
)

type Gamma interface {
	beta.Beta
}

func F(g Gamma) {
	beta.F(g)
}
