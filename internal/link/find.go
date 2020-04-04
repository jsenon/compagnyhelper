package link

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
)

func find(ctx context.Context, inputjson string) (ObjectLink, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "(*compagnyhelper).link.find")
	defer span.Finish()

	var myapp ObjectLink

	myjson, err := readFromFile(ctx, inputjson)
	if err != nil {
		return ObjectLink{}, err
	}

	err = json.Unmarshal(myjson, &myapp)

	if err != nil {
		return ObjectLink{}, err
	}

	log.Debug().Msgf("app: %v", myapp)

	return myapp, nil
}

func readFromFile(ctx context.Context, inputjson string) ([]byte, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "(*compagnyhelper).link.readFromFile")
	defer span.Finish()

	content, err := ioutil.ReadFile(inputjson) //nolint: gosec
	if err != nil {
		return nil, err
	}

	return content, nil
}
