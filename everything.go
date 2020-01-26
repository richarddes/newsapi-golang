package newsapi

import (
	"errors"
	"reflect"
	"time"
)

// EverythingOpts defines the options for the /everything route
type EverythingOpts struct {
	Sources, Domains, ExcludeDomains []string
	Q, QInTitle, Language, SortBy    string
	From, To                         time.Time
	PageSize                         uint8 // cannot be larger than 100 and smaller than 0 so uint8 is sufficient
	Page                             int
}

func checkEverythingParams(opts EverythingOpts) error {
	if opts.Q == "" && opts.QInTitle == "" && len(opts.Sources) < 1 && len(opts.Domains) < 1 {
		return errors.New("At least one of the following options must be specified: Q, QInTitle, Sources, Domains")
	}

	if opts.Language != "" && !isOptOf(opts.Language, langOpts) {
		return errors.New("A specified language isn't a valid language")
	}

	if opts.SortBy != "" && !isOptOf(opts.SortBy, sortByOpts) {
		return errors.New("The specified sortBy option isn't valid")
	}

	if len(opts.Sources) > 20 {
		return errors.New("You cannot specify more than 20 sources")
	}

	if opts.PageSize > 100 {
		return errors.New("The specified pageSize options is largen than the maximum of 100")
	}

	return nil
}

// Everything fetches data from the /everything route and returns the content as a struct
func (c *Client) Everything(opts EverythingOpts) (articleResp, error) {
	if reflect.ValueOf(opts).Kind() != reflect.Invalid {
		err := checkEverythingParams(opts)
		if err != nil {
			return articleResp{}, err
		}
	}

	body, err := fetchGetRoute("https://newsapi.org/v2/everything", c.APIKey, opts)
	if err != nil {
		return articleResp{}, err
	}

	return body.(articleResp), nil
}
