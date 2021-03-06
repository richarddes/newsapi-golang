package newsapi

import (
	"context"
	"errors"
	"reflect"
)

// SourcesOpts defines the options for the /sources route.
type SourcesOpts struct {
	Category, Country, Language string
}

type source struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

// SourcesResp represents what's being returned by the /source route.
type SourcesResp struct {
	Status  string   `json:"status"`
	Sources []source `json:"sources"`
}

func checkSourcesParams(opts SourcesOpts) error {
	if opts.Category != "" && !isOptOf(opts.Category, categoryOpts) {
		return errors.New("A specified category isn't a valid category")
	}

	if opts.Country != "" && !isOptOf(opts.Country, countryOpts) {
		return errors.New("A specified country isn't a valid country")
	}

	if opts.Language != "" && !isOptOf(opts.Language, langOpts) {
		return errors.New("A specified language isn't a valid language")
	}

	return nil
}

// Sources fetches data from the /sources route and returns the content as a SourcesResp object.
func (c *Client) Sources(ctx context.Context, opts SourcesOpts) (SourcesResp, error) {
	if reflect.ValueOf(opts).Kind() != reflect.Invalid {
		err := checkSourcesParams(opts)
		if err != nil {
			return SourcesResp{}, err
		}
	}

	body, err := fetchGetRoute(ctx, "https://newsapi.org/v2/sources", c.APIKey, opts)
	if err != nil {
		return SourcesResp{}, err
	}

	return body.(SourcesResp), nil
}
