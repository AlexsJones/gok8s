package configuration

const (
	positive = "✓"
	negative = "✗"
)

type item struct {
	uri       string
	validated string
	executed  string //✓ ✗
	success   string //✓ ✗
}

func (i *item) Executed(b bool) {

	switch b {
	case true:
		i.executed = positive
	case false:

		i.executed = negative
	}
}
func (i *item) Success(b bool) {

	switch b {
	case true:
		i.success = positive
	case false:

		i.success = negative
	}
}
func (i *item) Validated(b bool) {

	switch b {
	case true:
		i.validated = positive
	case false:

		i.validated = negative
	}
}
