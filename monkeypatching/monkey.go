package monkeypatching

type Person struct {
	Name         string
	tweetFetcher func(p *Person) string
}

func NewPerson(name string) *Person {
	return &Person{name, (*Person).fetchTweet}
}

// the hidden mockable thing
func (p *Person) fetchTweet() string {
	// This would actually use the net/http package to fetch the last
	// tweet.
	return "a tweet from the internetz"
}

func (p *Person) LastTweet() string {
	return p.Name + " tweeted " + p.tweetFetcher(p)
}
