package twitter

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostStatusUpdate(client *TwitterClient, statusText string) error {
	httpClient := getHttpClient(client)
	route := TwitterBaseURL + TwitterAPIVersion + "/statuses/update.json"

	req, _ := http.NewRequest(http.MethodPost, route, nil)
	url := req.URL.Query()
	url.Add("status", statusText)
	req.URL.RawQuery = url.Encode()

	res, err := httpClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("%s", string(body))
	}
	return nil
}
