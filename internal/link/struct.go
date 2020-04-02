package link

type objectLink struct {
	Applications []application `json:"applications"`
}

type application struct {
	Shortname string    `json:"shortname"`
	Desc      []appdesc `json:"appdesc"`
}

type appdesc struct {
	Longname string `json:"longname"`
	Link     string `json:"link"`
	Env      string `json:"env"`
}
