package impl

import (
	"testing"
)

const (
	test_api_key = "INPUT_YOUR_API_KEY"
	test_reg_id  = "INPUT_YOUR_REG_ID"
)

func Test_SendGCM(t *testing.T) {
	client := NewClient()
	data := map[string]interface{}{}
	_, err := client.Send(test_api_key, data, test_reg_id)
	if err != nil {
		t.Errorf("Failed to send GCM message : %s", err)
	}
}
