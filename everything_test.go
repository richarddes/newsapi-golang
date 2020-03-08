package newsapi

import (
	"testing"
)

func TestCheckEverythingParams(t *testing.T) {
	cases := []struct {
		c     EverythingOpts
		valid bool
	}{
		{EverythingOpts{}, false},
		{EverythingOpts{
			PageSize: 101,
		}, false},
		{EverythingOpts{
			Sources: []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"},
		}, false},
		{EverythingOpts{
			Language: " ",
		}, false},
		{EverythingOpts{
			Domains: []string{"reuters.com"},
		}, true},
	}

	for _, i := range cases {
		err := checkEverythingParams(i.c)
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
