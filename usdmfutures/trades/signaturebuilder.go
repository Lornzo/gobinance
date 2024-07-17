package trades

type signatureBuilder interface {
	SetAPIKeyAndSecret(key string, secret string)
	BuildRequestBodyWithSignature(body map[string]interface{}) (map[string]interface{}, error)
}
