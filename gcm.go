package gcm

type GCMClient interface {
	// Sends GCM message to GCM server.
	Send(apiKey string, data map[string]interface{}, regIds ...string) (Response, error)
}
