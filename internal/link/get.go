package link

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Get retrieve application link(s) available
// In cli compagnyhelper get link -n all
func Get(ctx context.Context, env string, serverurl string) {
	var linkstruct ObjectLink

	var resp *http.Response

	var requestbody []byte

	var req *http.Request

	var err error

	prefix := "/retrieve-links"

	client := &http.Client{}

	if env == "all" || env == "" {
		requestbody, err = json.Marshal(map[string]string{
			"env":  "all",
			"name": "",
		})
	} else {
		requestbody, err = json.Marshal(map[string]string{
			"env":  env,
			"name": "",
		})
	}

	if err != nil {
		fmt.Printf("Error retrieving application link: %v", err)
	}

	req, err = http.NewRequestWithContext(ctx, "POST", serverurl+prefix, bytes.NewBuffer(requestbody))
	if err != nil {
		fmt.Printf("Error retrieving application link: %v", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Error retrieving application link: %v", err)
	}
	defer resp.Body.Close() //nolint: errcheck

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&linkstruct)

	if err != nil {
		fmt.Println("Error reading body response:", err)
	}

	printresultGet(linkstruct)
}

// GetDetail retrieve application link(s) available
// In cli compagnyhelper get link -n devs
func GetDetail(ctx context.Context, env string, app string, serverurl string) {
	prefix := "/retrieve-link/" + app

	linkstruct, err := postWithData(ctx, env, prefix, serverurl)
	if err != nil {
		fmt.Printf("Error retrieving application link: %v", err)
	}

	printresultGetDetails(linkstruct)
}
