package link

// find env or name of the link inside yaml file

// Sample value for testing purpose
func initSAmple() []byte {
	sampleJSONLinks := []byte(`{
	"applications": [
		{
		"shortname": "Kibana",
			"appdesc": {
				"longname": "Kibana Dashboard for yourr logs",
				"link": "http://toto.com",
				"env": "dev"
			}
		},
		{
		"shortname": "Grafana",
			"appdesc": {
				"longname": "Grafana Dashboard for yourr metrics",
				"link": "http://tata.com",
				"env": "dev"
			}
		},
		{
			"shortname": "Kibana",
				"appdesc": {
					"longname": "Kibana Dashboard for yourr logs",
					"link": "http://titi.com",
					"env": "prod"
				}
		}
	]
}`)
	return sampleJSONLinks

}

func findAll() {

}

func findByNameEnv() {

}
