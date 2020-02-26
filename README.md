# NewsAPI Golang
[![GoDoc](https://godoc.org/github.com/richarddes/newsapi-golang?status.svg)](https://godoc.org/github.com/richarddes/newsapi-golang)

A simple to use Golang client for the [NewsAPI](https://newsapi.org/) service.

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
After that, one of three methods of the Client object can be called. The methods are called **TopHeadlines**, **Everything** and **Sources**. Each one of them takes in an options struct which is called like the method plus the word "Opts" at the end.
Here's an example of fetching the top headlines from the UK:
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

Since the **TopHeadlines** and **Everything** routes both return a response type of the same underlying type called "articleResp" you can cast them from one to another: 
```go
thr := newsapi.TopHeadlinesResp{}
er := newsapi.EverythingResp(t)
```
The decision to give the two routes different response types has been made to make the API more explicit but this might change in the future.

The API also has an "Article" type which represents an article in the "Articles" field of the "TopHeadlinesResp" or "EverythingResp" object. It's useful if you want to store the Articles returned by the **TopHeadlines** and **Everything** routes in a database.   
Here's a quick example of how to print the Titles of all articles returned by the **TopHeadlines** route:
```go
opts := newsapi.TopHeadlinesOpts{
  Country: "uk",
}

r, err := c.TopHeadlines(opts)
for _, a in range r.Articles {
	fmt.Println(article.Title)	
}
```


## Full Example
Here's a full runnable example on how to fetch the top headlines in the "business" category and save the recieved articles in a PostgreSQL database.The articles are being saved in a table with following schema:   
news(url TEXT PRIMARY KEY, author TEXT, title TEXT, source TEXT)
```go
package main

import (
	"database/sql"
	"log"

	"github.com/richarddes/newsapi-golang"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=user password=password host=localhost port=5432")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO news VALUES($1,$2,$3,$4);")
	if err != nil {
		log.Fatal(err)
	}

	c := newsapi.Client{APIKey: "your-api-key"}
	opts := newsapi.TopHeadlinesOpts{
		Country: "uk",
	}

	r, err := c.TopHeadlines(opts)
	if err != nil {
		log.Fatal(err)
	}

	for _, article := range r.Articles {
		_, err := stmt.Exec(article.URL, article.Author, article.Title, article.Source)
		if err != nil {
			log.Fatal(err)
		}
	}
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

MIT License. Click [here](https://choosealicense.com/licenses/mit/) or see the LICENSE file for details.
