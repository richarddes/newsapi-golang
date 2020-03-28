package newsapi

import (
	"context"
	"errors"
	"reflect"
)

// TopHeadlinesOpts defines the options for the /top-headlines route.
type TopHeadlinesOpts struct {
	PageSize             uint8  // cannot be larger than 100 and smaller than 0 so uint8 is sufficient
	Page                 uint16 // unlikely to be larger than ~65k
	Q, Category, Country string
	Sources              []string
}

// TopHeadlinesResp represents what's being returned by the /everything route. It relys on the same
// underlying type (called articleResp) that the EverythingResp type relys on. That means that
// it can easily be casted to the EverythingResp type.
type TopHeadlinesResp articleResp

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

// TopHeadlines fetches the data from the /top-headlines route and returns the response as a TopHeadlinesResp object.
func (c *Client) TopHeadlines(ctx context.Context, opts TopHeadlinesOpts) (TopHeadlinesResp, error) {
	if reflect.ValueOf(opts).Kind() != reflect.Invalid {
		err := checkTopHeadlinesParams(opts)
		if err != nil {
			return TopHeadlinesResp{}, err
		}
	}

	body, err := fetchGetRoute(ctx, "https://newsapi.org/v2/top-headlines", c.APIKey, opts)
	if err != nil {
		return TopHeadlinesResp{}, err
	}

	// we need this kind of type conversation because fetchGetRoute can only be parsed to the articleResp type
	// otherwise there will be an error
	return TopHeadlinesResp(body.(articleResp)), nil
}
