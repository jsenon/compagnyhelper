// Package link will find and return information for link from json file
package link

// ObjectLink describe multiple applications
type ObjectLink struct {
	Applications []Application `json:"applications"`
}

// Application describe an application
type Application struct {
	Shortname string  `json:"shortname"`
	Desc      appdesc `json:"appdesc"`
}

type appdesc struct {
	Longname string `json:"longname"`
	Link     string `json:"link"`
	Env      string `json:"env"`
}
