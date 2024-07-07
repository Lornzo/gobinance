package commons

import "context"

type RestfulBase interface {
	SetAPIKeyAndSecret(key string, secret string)
	SetBaseURL(url string)
	SetPathes(pathes ...string)
	AddPathes(pathes ...string)
	SetQuery(name string, value string)
	DelQuery(name string)
	UseSignature(isUse bool)
	UseApiKeyHeader(isUse bool)
	GET(ctx context.Context, result interface{}) (int, error)
	POST(ctx context.Context, body interface{}, result interface{}) (int, error)
	PUT(ctx context.Context, body interface{}, result interface{}) (int, error)
	DELETE(ctx context.Context, result interface{}) (int, error)
	Copy() RestfulBase
}
