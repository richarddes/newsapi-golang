package newsapi

import (
	"testing"
)

func TestCheckTopHeadlinesParams(t *testing.T) {
	cases := []struct {
		c     TopHeadlinesOpts
		valid bool
	}{
		{TopHeadlinesOpts{}, false},
		{TopHeadlinesOpts{
			PageSize: 101,
		}, false},
		{TopHeadlinesOpts{
			Category: "business",
			Sources:  []string{"reuters.com"},
		}, false},
		{TopHeadlinesOpts{
			Country: "de",
			Sources: []string{"reuters.com"},
		}, false},
		{TopHeadlinesOpts{
			Category: "health",
		}, true},
		{TopHeadlinesOpts{
			Country:  "us",
			PageSize: 100,
		}, true},
	}

	for _, i := range cases {
		err := checkTopHeadlinesParams(i.c)
		if !i.valid {
			if err == nil {
				t.Errorf("Expected error but got nil when case=%v", i)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error %v whenc case=%v", err, i.c)
			}
		}
	}
}
