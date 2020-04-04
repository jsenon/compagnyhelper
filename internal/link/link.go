package link

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
)

// Retrieve find the app with env and name provided
// Send longname value
//
// In cli compagnyhelper get link grafana -n dev, output will be
// Grafana |  http://grafana.com \n
//
// In cli compagnyhelper open link grafana -n dev, output will open a browser with url http://grafana.com
//
// In cli compagnyhelper describe link grafana -n dev
// Grafana | Your dashboard for your metrics | development | http://grafana.com \n
func Retrieve(ctx context.Context, env string, name string) (result Application, err error) {
	span, ctxChild := opentracing.StartSpanFromContext(ctx, "(*compagnyhelper).link.Retrieve")
	defer span.Finish()

	results, err := find(ctxChild)
	if err != nil {
		return Application{}, err
	}

	for i := range results.Applications {
		myobj := &results.Applications[i]
		if myobj.Shortname == name && myobj.Desc.Env == env {
			result.Shortname = myobj.Shortname
			result.Desc.Longname = myobj.Desc.Longname
			result.Desc.Link = myobj.Desc.Link
			result.Desc.Env = myobj.Desc.Env
		}
	}

	return result, nil
}

// RetrieveAll find all links with env
// Send an array of application struct
//
// In cli compagnyhelper get link -n dev, output will be
// Grafana \n Kibana \n Prometheus \n
//
// In cli compagnyhelper get link --all, output will be
// Grafana dev \n Kibana dev \n Grafana prod \n
func RetrieveAll(ctx context.Context, env string) (result ObjectLink, err error) {
	span, ctxChild := opentracing.StartSpanFromContext(ctx, "(*compagnyhelper).link.RetrieveAll")
	defer span.Finish()

	results, err := find(ctxChild)
	if err != nil {
		return ObjectLink{}, err
	}

	var myresult Application

	if env != "all" || env == "" {
		for i := range results.Applications {
			myobj := &results.Applications[i]
			if myobj.Desc.Env == env {
				myresult.Shortname = myobj.Shortname
				myresult.Desc.Longname = myobj.Desc.Longname
				myresult.Desc.Link = myobj.Desc.Link
				myresult.Desc.Env = myobj.Desc.Env
				result.Applications = append(result.Applications, myresult)
			}
		}

		log.Debug().Msgf("result: %v", result)

		return result, nil
	}

	return results, nil
}
