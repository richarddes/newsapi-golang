package newsapi

import (
	"testing"
	"time"
)

// use isCountryOpts instead of areOpts because it saves code in the test and
// tests the areOpts function with a relatively large dataset
// tests all isXxxOpt apart from isSortByOpt as well
func TestAreOpts(t *testing.T) {
	cases := []struct {
		c     string
		isOpt bool
	}{
		{"se", true},
		{"bg", true},
		{"il", true},
		{"za", true},
		{"ba", false},
		{"12", false},
		{"", false},
		{"hugh", false},
	}

	for _, i := range cases {
		var isCountry bool = isOptOf(i.c, countryOpts)
		if isCountry != i.isOpt {
			t.Errorf("Expected %v but got %v when case=%v", i.isOpt, isCountry, i.c)
		}
	}
}

func TestConstructURL(t *testing.T) {
	type mockStruct struct {
		// randomly named fields without any context
		Name string
		Age  int
		T    time.Time
		Pets []string
	}

	mockTm := time.Now()
	expectedkTm := mockTm.Format(time.RFC3339)

	cases := []struct {
		baseURL     string
		ms          mockStruct
		expectedURL string
	}{
		{"http://localhost:3000", mockStruct{Name: "steve", Age: 32}, "http://localhost:3000?name=steve&age=32"},
		{"https://localhost:3000", mockStruct{Name: "steve"}, "https://localhost:3000?name=steve"},
		{"http://google.com", mockStruct{Age: 3}, "http://google.com?age=3"},
		{"james.com", mockStruct{}, "james.com?"},
		{"steve.com", mockStruct{T: mockTm}, "steve.com?t=" + expectedkTm},
		{"steve.com", mockStruct{Pets: []string{"dog", "cat"}}, "steve.com?pets=dog,cat,"},
		{"a.com", mockStruct{Name: "steve", Age: 32, T: mockTm, Pets: []string{"dog", "cat"}}, "a.com?name=steve&age=32&t=" + expectedkTm + "&pets=dog,cat,"},
	}

	for _, i := range cases {
		u, err := constructURL(i.baseURL, i.ms)
		if err != nil {
			t.Fatal(err)
		}

		if u != i.expectedURL {
			t.Errorf("Expected %s but got %s when case=%v", i.expectedURL, u, i.ms)
		}
	}
}
