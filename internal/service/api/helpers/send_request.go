package helpers

import (
	"bytes"
	"fmt"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

func SendRequestToOrchestrator(orchestratorAddr string, body []byte, bearerAuth string) error {
	link := fmt.Sprintf(orchestratorAddr + "/requests/")

	req, err := http.NewRequest(http.MethodPost, link, bytes.NewReader(body))
	if err != nil {
		return errors.Wrap(err, "couldn't create request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearerAuth)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "error making http request")
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return errors.New(fmt.Sprintf("error in response, status %s", res.Status))
	}

	return nil
}
