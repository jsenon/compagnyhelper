package link

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

// DescribeLink find the app with env and name provided
// Send longname value
//
// In cli compagnyhelper describe link grafana -n dev, output will be
// Grafana | Your dashboard for your metrics \n
//
// In cli compagnyhelper open link grafana -n dev, output will open a browser with url myobj.Desc.Link
func DescribeLink(ctx context.Context, env string, name string) (result Application, err error) {
	span, ctxChild := opentracing.StartSpanFromContext(ctx, "(*compagnyhelper).link.DescribeLink")
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
		}
	}

	return result, nil
}

// RetrieveAll find all links with env
// Send an array of application struct
//
// In cli compagnyhelper get link grafana -n dev, output will be
// Grafana \n Kibana \n Prometheus \n
//
// In cli compagnyhelper get link grafana --all, output will be
// Grafana dev \n Kibana dev \n Grafana prod \n
func RetrieveAll(ctx context.Context, env string) (result ObjectLink, err error) {
	span, ctxChild := opentracing.StartSpanFromContext(ctx, "(*compagnyhelper).link.RetrieveAll")
	defer span.Finish()

	result, err = find(ctxChild)
	if err != nil {
		return ObjectLink{}, err
	}

	return result, nil
}

// Retrieve Find the app with env and name provided
// Send application struct
//
// In cli output will be
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
			result.Desc.Env = myobj.Desc.Env
			result.Desc.Link = myobj.Desc.Link
		}
	}

	return result, nil
}
