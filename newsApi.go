package newsapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Client represents the client
type Client struct {
	APIKey string
}

var (
	categoryOpts = []string{
		"business",
		"entertainment",
		"general",
		"health",
		"science",
		"sports",
		"technology",
	}

	langOpts = []string{
		"ar", "de", "en", "es", "fr", "he", "it",
		"nl", "no", "pt", "ru", "se", "ud", "zh",
	}

	countryOpts = []string{
		"ae", "ar", "at", "au", "be", "bg", "br", "ca", "ch", "cn", "co", "cu", "cz", "de", "eg", "fr", "gb", "gr",
		"hk", "hu", "id", "ie", "il", "in", "it", "jp", "kr", "lt", "lv", "ma", "mx", "my", "ng", "nl", "no", "nz",
		"ph", "pl", "pt", "ro", "rs", "ru", "sa", "se", "sg", "si", "sk", "th", "tr", "tw", "ua", "us", "ve", "za",
	}

	sortByOpts = []string{
		"popularity",
		"publishedAt",
		"relevancy",
	}
)

// also called source in the json response but has different vlaues than
// the sources field from the sources route
type articleSource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type article struct {
	Source      articleSource `json:"source"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	URL         string        `json:"url"`
	URLToImage  string        `json:"urlToImage"`
	PublishedAt string        `json:"publishedAt"`
	Content     string        `json:"content"`
}

// articleResp represents the response object that is returned by the /everything and /top-headlines routes
type articleResp struct {
	Status       string    `json:"status"`
	TotalResults uint      `json:"totalResults"`
	Articles     []article `json:"articles"`
}

func isOptOf(userOpt string, optArr []string) bool {
	i := sort.SearchStrings(optArr, userOpt)
	if i >= len(optArr) || optArr[i] != userOpt {
		return false
	}

	return true
}

// constructURL construct a url by taking the fields of a struct as url param keys and the struct values as param values
// it does not check if the url is a valid url; it only appends params to a base url
func constructURL(baseURL string, opt interface{}) (string, error) {
	s := reflect.ValueOf(opt)

	if s.Kind() == reflect.Ptr {
		s = reflect.Indirect(s)
	}

	// check that opt is of type struct
	if s.Kind() != reflect.Struct {
		return "", errors.New("Expected opt param to be of type struct but got something different")
	}

	buf := strings.Builder{}
	buf.WriteString(baseURL + "?")

	for i := 0; i < s.NumField(); i++ {
		var (
			fieldName  = strings.ToLower(s.Type().Field(i).Name)
			fieldValue = s.Field(i)
		)

		// rewrite to switch
		if fieldValue.IsValid() {
			switch kind := fieldValue.Kind(); {
			case kind == reflect.String && fieldValue.String() != "":
				buf.WriteString(fieldName + "=" + fieldValue.String() + "&")

			case kind == reflect.Int && fieldValue.Int() != 0:
				buf.WriteString(fieldName + "=" + strconv.Itoa(int(fieldValue.Int())) + "&")

			case kind == reflect.Uint && fieldValue.Uint() != 0:
				buf.WriteString(fieldName + "=" + strconv.Itoa(int(fieldValue.Uint())) + "&")

			case kind == reflect.TypeOf(time.Time{}).Kind() && !fieldValue.Interface().(time.Time).IsZero():
				buf.WriteString(fieldName + "=" + fieldValue.Interface().(time.Time).Format(time.RFC3339) + "&")

			case kind == reflect.Slice && fieldValue.Len() > 0:
				b := strings.Builder{}
				for i := 0; i < fieldValue.Len(); i++ {
					b.WriteString(fieldValue.Index(i).String() + ",")
				}
				buf.WriteString(fieldName + "=" + b.String() + "&")
			}
		}
	}

	result := buf.String()
	result = strings.TrimSuffix(result, "&")

	return result, nil
}

// fectchGetRoute exclusively fetches GET routes as other http methods aren't currently supported by the News API service
// and adding a param for the http methood seems unnecessary
func fetchGetRoute(baseURL, apiKey string, opt interface{}) (interface{}, error) {
	if apiKey == "" {
		return nil, errors.New("The API key cannot be nil")
	}

	url, err := constructURL(baseURL, opt)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Api-Key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// parse the json into the specific return type based on the baseURL route
	if strings.HasSuffix(baseURL, "/top-headlines") || strings.HasSuffix(baseURL, "/everything") {
		var body articleResp

		err = json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			return nil, err
		}

		return body, nil
	} else if strings.HasSuffix(baseURL, "/sources") {
		var body sourcesResp

		err = json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			return nil, err
		}

		return body, nil
	}

	return nil, errors.New("The specified route doesn't exist")
}
