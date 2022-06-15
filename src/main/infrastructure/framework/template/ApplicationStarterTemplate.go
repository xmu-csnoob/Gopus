package template

type ApplicationStarterTemplate interface {
	Initialize()
	InitDatabase()
	InitMiddleware()
	InitHttpInterfaces()
}
