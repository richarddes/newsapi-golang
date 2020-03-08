package newsapi

import "testing"

func TestCheckSourcesParams(t *testing.T) {
	cases := []struct {
		c     SourcesOpts
		valid bool
	}{
		{SourcesOpts{
			Category: "wrong-category",
		}, false},
		{SourcesOpts{
			Country: "wrong-country",
		}, false},
		{SourcesOpts{
			Language: "wrong-language",
		}, false},
		{SourcesOpts{}, true},
		{SourcesOpts{
			Category: "business",
			Language: "en",
		}, true},
		{SourcesOpts{
			Country: "us",
		}, true},
	}

	for _, i := range cases {
		err := checkSourcesParams(i.c)
		if !i.valid {
			if err == nil {
				t.Errorf("Expected error but got nil when case=%v", i.c)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error %v whenc case=%v", err, i.c)
			}
		}
	}
}
