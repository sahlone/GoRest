// Package dal Data Access Layer
package dbhandler

import (
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/config"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/logger"
	_ "github.com/lib/pq"
	mgo "gopkg.in/mgo.v2"
	"sync"
	"time"
)

type MongoHandler struct {
	*mgo.Session
	*config.Config
}

var once sync.Once

func (h *MongoHandler) RecipeAccess() (*mgo.Collection, *mgo.Session) {
	logger.Info("New mongo session created ")
	session := h.Copy()
	return session.DB(h.DBName).C("recipe"), session
}

func (h *MongoHandler) RecipeRateAccess() (*mgo.Collection, *mgo.Session) {
	logger.Info("New mongo session created")
	session := h.Copy()
	return session.DB(h.DBName).C("reciperate"), session
}

/*
 * This is used with fatal error as it should be used for fail fast while application startup
 */
func StartConnection(config *config.DBConfig) *mgo.Session {
	dialInfo := mgo.DialInfo{Addrs: config.Server, Timeout: time.Duration(config.Timeout) * time.Second}
	s, err := mgo.DialWithInfo(&dialInfo)
	if err != nil {
		logger.Error("connecting mongo %+v", dialInfo)
		logger.Fatal(err)
	}
	s.SetMode(mgo.Strong, true)
	return s
}
