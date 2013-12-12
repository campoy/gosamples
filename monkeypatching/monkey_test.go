package monkeypatching

import (
	"strings"
	"testing"
)

// This method won't be available in the non/test code.
func (p *Person) fakeFetchTweet() string {
	return "a fake tweet"
}

func TestPerson(t *testing.T) {
	p := NewPerson("John")
	// that we mock at some point
	p.tweetFetcher = (*Person).fakeFetchTweet
	lt := p.LastTweet()
	if !strings.Contains(lt, "a fake tweet") {
		t.Errorf("couldn't find the fake tweet in %q", lt)
	}
}
