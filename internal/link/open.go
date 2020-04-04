package link

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// OpenLink open a link in a browser
// In cli compagnyhelper open link Kibana -n dev
func OpenLink(ctx context.Context, env string, app string, serverurl string) {
	prefix := "/retrieve-link/" + app

	linkstruct, err := postWithData(ctx, env, prefix, serverurl)

	if err != nil {
		fmt.Printf("Error retrieving application link: %v", err)
	}

	if linkstruct.Desc.Link != "" {
		openbrowser(linkstruct.Desc.Link)
		return
	}

	fmt.Printf("Error retrieving application link\n")
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start() //nolint: gosec
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start() //nolint: gosec
	case "darwin":
		err = exec.Command("open", url).Start() //nolint: gosec
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}
}
