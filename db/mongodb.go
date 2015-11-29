package db

import (
	"os"

	"github.com/SDkie/metric_collector/logger"

	"gopkg.in/mgo.v2"
)

var mongoSession *mgo.Session
var dbName = os.Getenv("MG_DB_NAME")

// this initializes the db connection
func InitMongo() {
	var err error
	mongoSession, err = mgo.Dial(os.Getenv("MG_URI"))
	logger.PanicfIfError(err, "Error while initalizing MongoDB, %s", err)

	// Reads may not be entirely up-to-date, but they will always see the
	// history of changes moving forward, the data read will be consistent
	// across sequential queries in the same session, and modifications made
	// within the session will be observed in following queries (read-your-writes).
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	mongoSession.SetMode(mgo.Monotonic, true)

	logger.Info("MongoDB Successfully Initialize")
}

// this returns the mongo db session copy
// every of the routine that is requesting this db session is responsible for closing this session
// New creates a new session with the same parameters as the original session, // including consistency, batch size, prefetching, safety mode, etc. The
// returned session will use sockets from the pool, so there's a chance that
// writes just performed in another session may not yet be visible.
// Login information from the original session will not be copied over into the
// new session unless it was provided through the initial URL for the Dial
// function.
// Copy is similar to new but it just maitains the auth information
func GetMongoSession() *mgo.Session {
	return mongoSession.Copy()
}

func CloseMongo() {
	mongoSession.Close()
}

//MgCreate creates a new record of given interface
func MgCreate(c string, data interface{}) error {
	session := GetMongoSession()
	defer session.Close()
	return session.DB(dbName).C(c).Insert(data)
}
