package main

import (
	"fmt"
	"github.com/joja5627/scraper/internal/scrape"
	internalq "github.com/joja5627/scraper/internal/scrape"
)

func TestCase1() {
	q, _ := internalq.New(
		2, // Number of consumer threads
		&internalq.InMemoryQueueStorage{MaxSize: 10000},
	)
	c := scrape.BuildCollector(q)

	for _, state := range stateCodes {
		stateOrg := fmt.Sprintf("https://%s.craigslist.org", state)
		q.AddURL(fmt.Sprintf("%s/d/software-qa-dba-etc/search/sof", stateOrg))
		q.AddURL(fmt.Sprintf("%s/search/sof?employment_type=3", stateOrg))

	}

	q.Run(c)
}
