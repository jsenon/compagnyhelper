package link

import (
	"context"
	"fmt"
)

// Describe retrieve application link(s) availables
// In cli compagnyhelper describe link grafana -n dev
func Describe(ctx context.Context, env string, app string, serverurl string) {
	prefix := "/describe-link/" + app

	linkstruct, err := postWithData(ctx, env, prefix, serverurl)
	if err != nil {
		fmt.Printf("Error retrieving application link: %v", err)
	}

	printresultdesc(linkstruct)
}
