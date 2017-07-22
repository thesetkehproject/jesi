package functions

import (
	"github.com/thesetkehproject/jesi/config"
	"github.com/thesetkehproject/jesi/model"
	"gopkg.in/mgo.v2"
)

//MongoConnect Connect to MongoDB
func mongoConnect() (*mgo.Session, error) {
	session, err := mgo.Dial(config.DatabaseServer)
	if err != nil {
		return nil, err
	}
	return session, nil
}

//InsertScreenshot Insert Screenshot into DB
func InsertScreenshot(screenshot *model.Screenshot) error {
	db, err := mongoConnect()
	if err != nil {
		return err
	}

	c := db.DB(config.ScreenshotDB).C("screenshots")
	err = c.Insert(screenshot)
	return err //Should Return nil if all is well
}

func closeMongoConnect(session *mgo.Session) {
	session.Close()
}
