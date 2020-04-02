package link

import (
	"context"
	"encoding/json"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
)

// find env or name of the link inside yaml file

// Sample value for testing purpose
// will be replace by func for load the json file
func initSAmple() []byte {
	sampleJSONLinks := []byte(`{
	"applications": [
		{
		"shortname": "Kibana",
			"appdesc": {
				"longname": "Kibana Dashboard for yourr logs",
				"link": "http://kibana.com",
				"env": "dev"
			}
		},
		{
		"shortname": "Grafana",
			"appdesc": {
				"longname": "Grafana Dashboard for yourr metrics",
				"link": "http://grafana.com",
				"env": "dev"
			}
		},
		{
			"shortname": "Kibana",
				"appdesc": {
					"longname": "Kibana Dashboard for yourr logs",
					"link": "http://kibana.com",
					"env": "prod"
				}
		}
	]
	}`)

	return sampleJSONLinks
}

func find(ctx context.Context) (ObjectLink, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "(*compagnyhelper).link.findByNameEnv")
	defer span.Finish()

	var myapp ObjectLink

	myjson := initSAmple()
	err := json.Unmarshal(myjson, &myapp)

	if err != nil {
		return ObjectLink{}, err
	}

	log.Debug().Msgf("app: %v", myapp)

	return myapp, nil
}
