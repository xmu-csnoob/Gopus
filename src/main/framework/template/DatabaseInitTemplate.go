package template

type DatabaseInitTemplate interface {
	Initialize()
	ReadConfigs()
	ConnectToDatabase()
}
