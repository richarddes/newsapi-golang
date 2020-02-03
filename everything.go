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

// EverythingResp is what's being returned by the Everything route
// it only exists to be a more obvious than articleResp
type EverythingResp articleResp

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
func (c *Client) Everything(opts EverythingOpts) (EverythingResp, error) {
	if reflect.ValueOf(opts).Kind() != reflect.Invalid {
		err := checkEverythingParams(opts)
		if err != nil {
			return EverythingResp{}, err
		}
	}

	body, err := fetchGetRoute("https://newsapi.org/v2/everything", c.APIKey, opts)
	if err != nil {
		return EverythingResp{}, err
	}

	// we need this kind of type conversation because fetchGetRoute can only be parsed to the articleResp type
	// otherwise there will be an error
	return EverythingResp(body.(articleResp)), nil
}
