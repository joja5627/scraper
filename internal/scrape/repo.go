package scrape

import (
	"github.com/globalsign/mgo"
)

//Repo comments are meaningless
type Repo interface {
	// FindAll() ([]*User, error)
	// FindByUsername(username string) (*User, error)
	Add(listing *Listing) error
	// Remove(username string) error
}

// MongoRepository represents a mongodb repository
type MongoRepository struct {
	session *mgo.Session
}

// NewMongoRepository creates a mongo API definition repo
func NewMongoRepository(session *mgo.Session) (*MongoRepository, error) {
	return &MongoRepository{session}, nil
}

// FindAll fetches all the API definitions available
// func (r *MongoRepository) FindAll() ([]*User, error) {
// 	var result []*User
// 	session, coll := r.getSession()
// 	defer session.Close()

// 	err := coll.Find(nil).Sort("username").All(&result)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }

// FindByUsername find an user by username
// func (r *MongoRepository) FindByUsername(username string) (*User, error) {
// 	return r.findOneByQuery(bson.M{"username": username})
// }

// func (r *MongoRepository) findOneByQuery(query interface{}) (*User, error) {
// 	var result User
// 	session, coll := r.getSession()
// 	defer session.Close()

// 	err := coll.Find(query).One(&result)
// 	if err != nil {
// 		if err == mgo.ErrNotFound {
// 			return nil, ErrUserNotFound
// 		}
// 		return nil, err
// 	}

// 	return &result, nil
// }

// Add adds an user to the repository
//func (r *MongoRepository) Add(listing *Listing) error {
//	session, coll := r.getSession()
//	defer session.Close()
//	listing.Date = time.Now()
//	coll.Insert(listing)
//
//	// _, err := coll.Upsert(bson.M{"username": user.Username}, user)
//	// if err != nil {
//	// 	log.WithField("username", user.Username).Error("There was an error adding the user")
//	// 	return err
//	// }
//
//	// log.WithField("username", user.Username).Debug("User added")
//	return nil
//}
func (r *MongoRepository) getSession() (*mgo.Session, *mgo.Collection) {
	session := r.session.Copy()
	coll := session.DB("scrape_db").C("scrape_collection")

	return session, coll
}

// Remove an user from the repository
// func (r *MongoRepository) Remove(username string) error {
// 	session, coll := r.getSession()
// 	defer session.Close()

// 	err := coll.Remove(bson.M{"username": username})
// 	if err != nil {
// 		if err == mgo.ErrNotFound {
// 			return ErrUserNotFound
// 		}
// 		log.WithField("username", username).Error("There was an error removing the user")
// 		return err
// 	}

// 	log.WithField("username", username).Debug("User removed")
// 	return nil
// }

func (r *MongoRepository) GetAll()  []map[string]string {
	var listings []map[string]string // see bson.M
	session, coll := r.getSession()
	coll.Find(nil).All(&listings)
	defer session.Close()
	return listings
}
