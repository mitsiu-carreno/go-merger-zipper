package main

import(
	"flag"
	"os"
	"time"
	"github.com/mitsiu-carreno/go-merger-zipper/utils"
	"github.com/mitsiu-carreno/go-merger-zipper/merger"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	models "github.com/mitsiu-carreno/go-merger-zipper/declarations"
)

func main(){
	var logpath = flag.String("logpath", os.Getenv("LOG_FILE"), "Log Path")
	utils.NewLog(*logpath)
	// DATABASE
	var (
		hosts 		= os.Getenv("MAIN_DB_HOST")
		database	= os.Getenv("MAIN_DB_DB")
		username	= os.Getenv("MAIN_DB_USER")
		password	= os.Getenv("MAIN_DB_PASSWORD")
		collection	= os.Getenv("MAIN_DB_COLLECTION")
		inputPath 	= os.Getenv("FILE_INPUT")

		newFileName = "test-merge-2017"
	)

	info := &mgo.DialInfo{
		Addrs 		: []string{hosts},
		Timeout		: 60 * time.Second,
		Database	: database,
		Username	: username,
		Password	: password,
	}

	session, err := mgo.DialWithInfo(info)
	utils.Check(err)
	defer session.Close()

	col := session.DB(database).C(collection)

	var mgoResult []models.Declarations
	err = col.Find(bson.M{"ANIO":2017}).All(&mgoResult)
	utils.Check(err)
	utils.Log.Println(len(mgoResult), " documents to be merged")
	merger.Merger(inputPath, newFileName + ".csv", mgoResult)

	zipper.Zipper(newFile + ".zip", )
}