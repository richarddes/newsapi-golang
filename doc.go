/*
Package newsapi provides a client library for the "NewsAPI" service.
See https://newsapi.org/ for more information.

Every project has to first initialize a client obejct with an API key like so:

	c := newsapi.Client{APIKey: "your-api-key"}

After that, one of three methods of the Client object can be called. The methods are called TopHeadlines, Everything and Sources. 
Each one of them takes in an options struct which is called like the method plus the word "Opts" at the end. Here's an example of fetching the top headlines from the UK:

opts := newsapi.TopHeadlinesOpts{
  Country: "uk",
}

	r, err := c.TopHeadlines(opts)
	if err != nil {
		log.Fatal(err)
	}

When fetching top headlines at least one of the following options must be specified: Q, Category, Country or Sources You also cannot specify the Sources option in conjunction with the Category or Country option.

When fetching everything at least one of the following options must be specified: Q, QInTitle, Sources or Domains

Since the TopHeadlines and Everything routes both return a response type of the same underlying type called "articleResp" you can cast them from one to another:

	thr := newsapi.TopHeadlinesResp{}
	er := newsapi.EverythingResp(t)

The decision to give the two routes different response types has been made to make the API more explicit but this might change in the future.

The API also has an "Article" type which represents an article in the "Articles" field of the "TopHeadlinesResp" or "EverythingResp" object. It's useful if you want to store the Articles returned by the TopHeadlines and Everything routes in a database.
For a full example of how to do this, please refer to the github page of this library (https://github.com/richarddes/newsapi-golang).
Here's a quick example of how to print the Titles of all articles returned by the TopHeadlines route:

	opts := newsapi.TopHeadlinesOpts{
		Country: "uk",
	}

	r, err := c.TopHeadlines(opts)
	for _, a in range r.Articles {
		fmt.Println(article.Title)
	}
*/

package newsapi
