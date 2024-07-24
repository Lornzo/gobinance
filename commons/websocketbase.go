package commons

type WebsocketBase interface {
	SetBaseURL(url string) error
	SetPathes(pathes ...string) error
	GetWebsocketURL() string
	SetRunning(isRunning bool)
	IsRunning() bool
	SetDebug(isDebug bool)
	IsDebug() bool
}
