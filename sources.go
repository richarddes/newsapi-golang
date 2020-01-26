package newsapi

import (
	"errors"
	"reflect"
)

// SourcesOpts defines the options for the /sources route
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

type sourcesResp struct {
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

// Sources fetches data from the /sources route and returns the content as a struct
func (c *Client) Sources(opts SourcesOpts) (sourcesResp, error) {
	if reflect.ValueOf(opts).Kind() != reflect.Invalid {
		err := checkSourcesParams(opts)
		if err != nil {
			return sourcesResp{}, err
		}
	}

	body, err := fetchGetRoute("https://newsapi.org/v2/sources", c.APIKey, opts)
	if err != nil {
		return sourcesResp{}, err
	}

	return body.(sourcesResp), nil
}
