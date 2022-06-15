package template

type ControllerTemplate interface {
	Start()
	GetHttpEngine()
	RegisterHttpInterfaces()
	RunHttpEngine()
}
