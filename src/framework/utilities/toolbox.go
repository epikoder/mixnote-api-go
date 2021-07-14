package utilities

type toolbox struct {}

var Toolbox *toolbox

func init() {
	once.Do(func() {
		Toolbox = new(toolbox)
	})
}

func (*toolbox) Isset(a map[string]interface{}, key string) bool {
	x := a[key];return x == nil
}
