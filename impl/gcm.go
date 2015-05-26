package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mokelab-go/gcm"
	"net/http"
)

const (
	gcm_url = "https://android.googleapis.com/gcm/send"
)

type gcmClient struct {
}

func NewClient() *gcmClient {
	return &gcmClient{}
}

func (c *gcmClient) Send(apiKey string, data map[string]interface{}, regIds ...string) (gcm.Response, error) {
	msg := map[string]interface{}{
		"registration_ids": regIds,
		"data":             data,
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return gcm.Response{}, err
	}
	req, err := http.NewRequest("POST", gcm_url, bytes.NewBuffer(body))
	if err != nil {
		return gcm.Response{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("key=%s", apiKey))
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return gcm.Response{}, err
	}
	defer resp.Body.Close()

	return gcm.Response{}, nil
}
