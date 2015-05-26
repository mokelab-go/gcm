package impl

import (
	"testing"
)

const (
	test_api_key = "AIzaSyAW3a1RrqMCnDFQry22ymV86O9IDJtxr_Y"
	test_reg_id  = "APA91bHfqbxReVo4DriNkV_rDxOMTmCp0h4WLldpzIoPgU9QAeLJQLpkgfUdGdsxBFSQSXLvS2JVg7oKFJ5PJvMKuOzRuB9Bd9QVTJr48tV02ncz0Phj9f8gEzHBFL1RzS97-LQR5-xmZu7h6x-k64taLKGlyq-2YPAxTrgNxmL6iCjuig_NoSA"
)

func Test_SendGCM(t *testing.T) {
	client := NewClient()
	data := map[string]interface{}{}
	_, err := client.Send(test_api_key, data, test_reg_id)
	if err != nil {
		t.Errorf("Failed to send GCM message : %s", err)
	}
}
