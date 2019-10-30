package scrape

import (
	"github.com/globalsign/mgo/bson"
	"time"
)


//Listing comment
type Listing struct {
	ID    bson.ObjectId "_id,omitempty"
	ScrapeID bson.ObjectId "_id,omitempty"
	Date  time.Time     `json:"created"`
	StateOrg string `json:"stateOrg"`
	ListingUrl string `json:"listingUrl"`
	QueryUrl string `json:"queryUrl"`
	ContactInfoUrl string `json:"contactInfoUrl"`
	EmailResponse string `json:"emailResponse"`
	Email string `json:"email"`
}
//SocketMessage
type SocketMessage struct{
	MessageType string `json:"messageType"`
	Payload string `json:"payload"`
}