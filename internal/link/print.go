package link

import (
	"fmt"
)

const tabName string = "APPLICATION(S)"
const envName string = "ENV"
const linkName string = "URL"
const descName string = "DESCRIPTION"

//printresult will print the result to stdout for get command
func printresultGet(linkstruct ObjectLink) {
	fmt.Printf("|%-15s|%-15s|\n", tabName, envName)

	if linkstruct.Applications == nil {
		fmt.Printf("No application links found \n")
	}

	for i := range linkstruct.Applications {
		myobj := &linkstruct.Applications[i]
		fmt.Printf("|%-15s|%-15s|\n", myobj.Shortname, myobj.Desc.Env)
	}
}

//printresult will print the result to stdout for get command
func printresultGetDetails(linkstruct Application) {
	fmt.Printf("|%-15s|%-15s|%-30s|\n", tabName, envName, linkName)

	if linkstruct.Shortname == "" {
		fmt.Printf("No application links found \n")
		return
	}

	fmt.Printf("|%-15s|%-15s|%-30s|\n", linkstruct.Shortname, linkstruct.Desc.Env, linkstruct.Desc.Link)
}

//printresultdesc will print the result to stdout for a describe command
func printresultdesc(linkstruct Application) {
	fmt.Printf("|%-15s|%-50s|%-15s|%-30s|\n", tabName, descName, envName, linkName)

	if linkstruct.Shortname == "" {
		fmt.Printf("No application links found \n")
		return
	}

	fmt.Printf("|%-15s|%-50s|%-15s|%-30s|\n", linkstruct.Shortname,
		linkstruct.Desc.Longname, linkstruct.Desc.Env, linkstruct.Desc.Link)
}
