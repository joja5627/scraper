package scrape

import (
	"github.com/google/uuid"
)

//Listing comment
type Listing struct {

	//ScrapeID       bson.ObjectId "_id,omitempty"
	//Date           time.Time     `json:"created"`
	ID                  uuid.UUID `json:"id"`
	StateCode           string    `json:"stateOrg"`
	Title               string    `json:"title"`
	Url                 string    `json:"url"`
	Query               string    `json:"query"`
	ContactInfoUrl      string    `json:"contactInfoUrl"`
	ListingInfoResponse string    `json:"listingInfoResponse"`
	EmailResponse       string    `json:"emailResponse"`
	Email               string    `json:"email"`
}

//SocketMessage
type SocketMessage struct {
	MessageType string `json:"messageType"`
	Payload     string `json:"payload"`
}
