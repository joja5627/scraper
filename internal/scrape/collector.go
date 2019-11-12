package scrape

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var errors []string
var links []string
var requests []string
var contactInfos []string
var listUA []string
var uaList []string
var proxyList []string

const proxyListAPI = "https://www.proxy-list.download/api/v1/get?type=http"
const userAgentAPI = "http://www.ua-tracker.com/user_agents.txt"

func getUserAgentsList(collector *colly.Collector) ([]string, error) {

	collector.OnResponse(func(response *colly.Response) {
		if response.StatusCode == 200 {
			for _, ua := range strings.Split(string(response.Body), "\n") {
				uaList = append(uaList, ua)
			}
		}
	})
	visitWithRetry(collector, userAgentAPI, 30)
	return uaList, nil
}
func getProxyList(collector *colly.Collector) ([]string, error) {

	collector.OnResponse(func(response *colly.Response) {
		if response.StatusCode == 200 {
			for _, pr := range strings.Split(string(response.Body), "\r\n") {
				proxyList = append(proxyList, fmt.Sprintf(`http://%s`, pr))
			}
		}
	})

	visitWithRetry(collector, proxyListAPI, 30)
	return proxyList, nil
}
func visitWithRetry(collector *colly.Collector, URL string, retryCount int) {
	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("errors: ", err.Error())
		fmt.Println("errors: ", len(errors))
		if retryCount > 0 {
			r.Request.Retry()
		}
	})
	collector.Visit(URL)
	collector.Wait()
}

func getProxyFunc(collector *colly.Collector) (colly.ProxyFunc, error) {
	proxyList, err := getProxyList(collector)
	if err != nil {
		return nil, err
	}
	proxyFunc, err := proxy.RoundRobinProxySwitcher(proxyList...)
	if err != nil {
		return nil, err
	}
	return proxyFunc, nil
}
func getRandomItem(len int) int {
	return int(math.Abs(float64(rand.Intn(len - 1))))
}

func ChangeUAWithTimeout(changingTimeout time.Duration, collector *colly.Collector) {
	rand.Seed(time.Now().Unix())
	for range time.NewTicker(time.Duration(changingTimeout * time.Minute)).C {
		collector.UserAgent = listUA[getRandomItem(len(listUA))]
	}
}

func VisitWithRetry(collector *colly.Collector, URL string, retryCount int) {
	if err := collector.Visit(URL); err != nil {
		count := 1
		fmt.Println("Try", count, err)
		for count <= retryCount {
			count++
			if err := collector.Visit(URL); err != nil {
				fmt.Println("Try", count, err)
			} else {
				break
			}
		}
	}
}

func BuildCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowURLRevisit(),
		colly.Async(true))

	proxyFunc, err := getProxyFunc(collector)
	if err != nil {
		fmt.Println("can't set proxy")
	}
	list, err := getUserAgentsList(collector)
	if err != nil {
		fmt.Println("can't build user agent list")
	}
	listUA = list

	collector.SetProxyFunc(proxyFunc)
	collector.OnHTML("a.result-title.hdrlnk", func(e *colly.HTMLElement) {

		listingURL := e.Attr("href")
		links = append(links, listingURL)
		fmt.Println("links: ", len(links))
		fmt.Println(listingURL)
		collector.Visit(listingURL)

	})
	collector.OnHTML("button.reply-button.js-only", func(e *colly.HTMLElement) {
		info := e.Attr("data-href")
		r, _ := regexp.Compile("https://([a-z]+).craigslist.org")
		contactInfoTop := fmt.Sprintf("%s/contactinfo/", r.FindString(e.Request.URL.String()))
		info = strings.Replace(info, "/__SERVICE_ID__/", contactInfoTop, 1)
		contactInfos = append(contactInfos, info)
		fmt.Println("contactInfos: ", len(contactInfos))

	})

	collector.Limit(&colly.LimitRule{
		DomainGlob:  "*.craigslist.org.*",
		Parallelism: 2,
		RandomDelay: 60 * time.Second,
	})

	collector.OnRequest(func(r *colly.Request) {
		requests = append(requests, r.URL.String())
		fmt.Println("errors: ", len(requests))
	})

	//collector.OnError(func(r *colly.Response, err error) {
	//	errors = append(errors,r.Request.URL.String())
	//	fmt.Println("errors: ", len(errors))
	//	//fmt.Println(r.Request.Retry())
	//	//if(err.TimeOut())
	//	//retryURL := r.Request.URL.String()
	//	//r.Request.Retry()
	//
	//	//collector.Visit(retryURL)
	//})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("visited", r.Request.URL)
	})
	go ChangeUAWithTimeout(1, collector)

	return collector
}
