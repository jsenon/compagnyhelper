package link

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func postWithData(ctx context.Context, env string, prefix string, serverurl string) (Application, error) {
	var linkstruct Application

	var resp *http.Response

	var req *http.Request

	client := &http.Client{}

	switch env {
	case "":
		requestbody, err := json.Marshal(map[string]string{
			"env":  "all",
			"name": "",
		})
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

	default:
		requestbody, err := json.Marshal(map[string]string{
			"env":  env,
			"name": "",
		})
		if err != nil {
			return linkstruct, err
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
	}

	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&linkstruct)

	if err != nil {
		return linkstruct, err
	}

	return linkstruct, nil
}
