package fooda

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

var hipchatToken string

func init() {
	hipchatToken = os.Getenv("HIPCHAT_TOKEN")
}

type Hipchat struct {
}

func (Hipchat) Notify(userOrRoom string, message string) error {
	url := fmt.Sprintf("http://api.hipchat.com/v2/room/%s/notification?auth_token=%s", userOrRoom, hipchatToken)

	// use json.Marshal(data)
	var json = []byte(fmt.Sprintf(`{"color": "purple", "message_format": "text", "message":"%s"}`, message))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)

	if err != nil {
		return err
	}

	return nil
}
