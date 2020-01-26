package newsapi

import (
	"testing"
)

func TestCheckEverythingParams(t *testing.T) {
	cases := []EverythingOpts{
		EverythingOpts{
			Sources: []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"},
		},
		EverythingOpts{
			PageSize: 101,
		},
		EverythingOpts{},
	}

	for _, i := range cases {
		if checkEverythingParams(i) == nil {
			t.Errorf("Expected error but got nil when case=%v", i)
		}
	}
}
