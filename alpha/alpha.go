package alpha

type A struct {
	v int
}

func (a *A) Inc() {
	a.v++
}

type Alpha interface {
	F(a A)
}
