# NewsAPI Golang
[![GoDoc](https://godoc.org/github.com/richarddes/newsapi-golang?status.svg)](https://godoc.org/github.com/richarddes/newsapi-golang)

A simple to use Golang client for [NewsAPI](https://newsapi.org/).

Before you use the package you should read the NewsAPI [docs](https://newsapi.org/docs) to get familiar with the endpoints and their options.

## Install
```sh
go get github.com/richarddes/newsapi-golang
```

## Usage
Every project has to first initialize a client obejct with an API key like so:
```go
c := newsapi.Client{APIKey: "your-api-key"}
```
After that, we can call one of three methods of the Client object. The methods are called **TopHeadlines**, **Everything** and
**Sources**. Each one of them takes in an options struct which is called like the method plus the word "Opts" at the end.
Here's an example of fetching the top headlines from the US:
```go
opts := newsapi.TopHeadlinesOpts{
  Country: "uk",
}

r, err := c.TopHeadlines(opts)
if err != nil {
	log.Fatal(err)
}
```
When fetching top headlines at least one of the following options must be specified: Q, Category, Country or Sources
You also cannot specify the Sources option in conjunction with the Category or Country option.

When fetching everything at least one of the following options must be specified: Q, QInTitle, Sources or Domains

For more details about the options structs please refer to the [docs](https://godoc.org/github.com/richarddes/newsapi-golang).


## Full Example
Here's a full runnable example where we fetch the top headlines in the "business" category and print the title of each article we've received.
```go
package main

import (
	"fmt"
	"log"

	"github.com/richarddes/newsapi-golang"
)

func main() {
	c := newsapi.Client{APIKey: "your-api-key"}
	opts := newsapi.TopHeadlinesOpts{
		Country: "uk",
	}

	r, err := c.TopHeadlines(opts)
	if err != nil {
		log.Fatal(err)
	}

	for _, article := range r.Articles {
		fmt.Println(article.Title)
	}
}

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

MIT License. Click [here](https://choosealicense.com/licenses/mit/) or see the LICENSE file for details.
