package newsapi

import (
	"testing"
)

func TestCheckTopHeadlinesParams(t *testing.T) {
	cases := []TopHeadlinesOpts{
		TopHeadlinesOpts{
			Category: "business",
			Sources:  []string{"reuters.com"},
		},
		TopHeadlinesOpts{
			Country: "de",
			Sources: []string{"reuters.com"},
		},
		TopHeadlinesOpts{
			PageSize: 101,
		},
		TopHeadlinesOpts{},
	}

	for _, i := range cases {
		if checkTopHeadlinesParams(i) == nil {
			t.Errorf("Expected error but got nil when case=%v", i)
		}
	}
}
