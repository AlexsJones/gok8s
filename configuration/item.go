package configuration

const (
	positive = "✓"
	negative = "✗"
)

type item struct {
	Uri      string `json:"uri"`
	Executed string `json:"executed"` //✓ ✗
	Success  string `json:"success"`  //✓ ✗
	Log      string `json:"log"`
}

func (i *item) isExecuted(b bool) {

	switch b {
	case true:
		i.Executed = positive
	case false:
		i.Executed = negative
	}
}
func (i *item) isSuccess(b bool) {

	switch b {
	case true:
		i.Success = positive
	case false:

		i.Success = negative
	}
}
