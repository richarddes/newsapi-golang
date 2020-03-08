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

// Client represents the client type for the apu
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

// articleSource is called "source" in the json response but it has a different values
// than the "sources" field in the /sources field.
type articleSource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// The Article type represents an article returned in the "articles" field of the /everything and
// /top-headliens routes.
type Article struct {
	Source      articleSource `json:"source"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	URL         string        `json:"url"`
	URLToImage  string        `json:"urlToImage"`
	PublishedAt string        `json:"publishedAt"`
	Content     string        `json:"content"`
}

// articleResp is the underlying type for the TopHeadlinesResp and EverythingResp types. It represents
// the  response from the /top-headlines and /everything routes.
type articleResp struct {
	Status       string    `json:"status"`
	TotalResults uint      `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func isOptOf(userOpt string, optArr []string) bool {
	// There's not really a reason to use binary search here as it doesn't improve performence
	// significantly, but hey, why not? We need less lines of code that way (one or two less probably)
	i := sort.SearchStrings(optArr, userOpt)
	if i >= len(optArr) || optArr[i] != userOpt {
		return false
	}

	return true
}

// constructURL construct a url by taking the fields of a struct as url param keys and the struct values as param values.
// It does not check if the url is a valid url; it only appends params to a base url
func constructURL(baseURL string, opt interface{}) (string, error) {
	s := reflect.ValueOf(opt)

	if s.Kind() == reflect.Ptr {
		s = reflect.Indirect(s)
	}

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

// fectchGetRoute exclusively fetches GET routes as other http methods aren't currently supported by the "NewsAPI" service
// and adding a param for the http methood seems unnecessary and just makes things more complicated
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
		var body SourcesResp

		err = json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			return nil, err
		}

		return body, nil
	}

	return nil, errors.New("The specified route doesn't exist")
}
