package scrape

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//Scrape comment
type Scrape struct {
	ID    bson.ObjectId "_id,omitempty"
	Links string        `json:"links"`
	Date  time.Time     `json:"created"`
}
