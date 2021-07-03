package configs

type app struct{
	Name string
}

func App() (a app) {
	a.Name = "Mixnote"
	return
}