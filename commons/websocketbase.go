package commons

type WebsocketBase interface {
	SetAPIKeyAndSecret(key string, secret string)
	SetBaseURL(url string)
	SetPathes(pathes ...string)
	GetWebsocketURL() string
	SetRunning(isRunning bool)
	IsRunning() bool
}
