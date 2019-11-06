package async

import (
	"fmt"
	"github.com/joja5627/github.com/joja5627/scraper/internal/scrape"
	"io/ioutil"
)

// FetchAll grabs a list of urls
func FetchAll(urls []string, c *Client) {
	for _, url := range urls {
		go c.AsyncGet(url)
	}
}
// FetchAll grabs a list of urls
func FetchAllHigherOrder(listings []scrape.Listing, c *Client,parseResponse func(response string)) {
	for _, listing := range listings {
		if listing.ContactInfoUrl != "" {
			go c.AsyncGet(listing.ContactInfoUrl)
		}

	}
	for i := 0; i < len(listings); i++ {
		select {
		case resp := <-c.Resp:
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil{
				fmt.Printf("error")
			}
			stringBody := string(body)
			fmt.Printf("Status received for %s", stringBody)
			listings[i].EmailResponse = stringBody
		case err := <-c.Err:
			fmt.Printf("Error received: %s\n", err)
			listings[i].EmailResponse = fmt.Sprintf("%s",err)
		}
	}
}
