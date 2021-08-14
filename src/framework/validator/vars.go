package validator

type (
	FormVar map[string]interface{}
	ErrorBag map[string][]string
)

func (f FormVar) GetString(s string) string {
	return f[s].(string)
}

func (f FormVar) GetInt(s string) uint64 {
	return f[s].(uint64)
}

func (e ErrorBag) Failed() (bool) {
	return len(e) != 0
}
