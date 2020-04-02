package link

import (
	"context"

	"github.com/opentracing/opentracing-go"
)

func Open(ctx context.Context, env string, name string) (result objectLink, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "(*compagnyhelper).link.Open")
	defer span.Finish()

	// Find the link with env and name provided
	// Send url value
	// In cli compagnyhelper open link grafana -n dev, open link to the browser http://grafana.com

	return result, nil

}

func DescribeLink(ctx context.Context, env string, name string) (result objectLink, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "(*compagnyhelper).link.DescribeLink")
	defer span.Finish()

	// Find the app with env and name provided
	// Send longname value
	// In cli compagnyhelper describe link grafana -n dev, ouput will be
	// Grafana | Your dashboard for your metrics \n
	return result, nil

}

func RetrieveAll(ctx context.Context, env string) (result objectLink, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "(*compagnyhelper).link.RetrieveAll")
	defer span.Finish()

	// Find all links with env
	// Send an array of application struct
	// In cli compagnyhelper get link grafana -n dev, output will be
	// Grafana \n Kibana \n Prometheus \n
	//
	// In cli compagnyhelper get link grafana --all, output will be
	// Grafana dev \n Kibana dev \n Grafana prod \n

	return result, nil

}

func Retrieve(ctx context.Context, env string, name string) (result objectLink, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "(*compagnyhelper).link.Retrieve")
	defer span.Finish()

	// Find the app with env and name provided
	// Send application struct
	// In cli ouput will be
	// Grafana | Your dashboard for your metrics | development | http://grafana.com \n

	return result, nil
}
