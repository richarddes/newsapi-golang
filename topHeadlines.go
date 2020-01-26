package newsapi

import (
	"errors"
	"reflect"
)

// TopHeadlinesOpts defines the options for the /top-headlines route
type TopHeadlinesOpts struct {
	Sources              []string
	Q, Category, Country string
	PageSize, Page       int
}

func checkTopHeadlinesParams(opts TopHeadlinesOpts) error {
	if opts.Q == "" && opts.Category == "" && opts.Country == "" && len(opts.Sources) < 1 {
		return errors.New("At least one of the following options must be specified: Q, Category, Country, Sources")
	}

	if opts.Category != "" && !isOptOf(opts.Category, categoryOpts) {
		return errors.New("A specified category isn't a valid category")
	}

	if opts.Country != "" && !isOptOf(opts.Country, countryOpts) {
		return errors.New("A specified country isn't a valid country")
	}

	if len(opts.Sources) > 0 && len(opts.Category) > 0 || len(opts.Sources) > 0 && len(opts.Country) > 0 {
		return errors.New("The category and country options cannopt be used in conjunction with the sources option")
	}

	if opts.PageSize > 100 {
		return errors.New("The specified pageSize options is largen than the maximum of 100")
	}

	return nil
}

// TopHeadlines fetches data from the /top-headlines route and returns the content as a struct
func (c *Client) TopHeadlines(opts TopHeadlinesOpts) (articleResp, error) {
	if reflect.ValueOf(opts).Kind() != reflect.Invalid {
		err := checkTopHeadlinesParams(opts)
		if err != nil {
			return articleResp{}, err
		}
	}

	body, err := fetchGetRoute("https://newsapi.org/v2/top-headlines", c.APIKey, opts)
	if err != nil {
		return articleResp{}, err
	}

	return body.(articleResp), nil
}
