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

func check(e error){
	if e != nil{
		utils.Log.Println(e)
		panic(e)
	}
}

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
	)

	info := &mgo.DialInfo{
		Addrs 		: []string{hosts},
		Timeout		: 60 * time.Second,
		Database	: database,
		Username	: username,
		Password	: password,
	}

	session, err := mgo.DialWithInfo(info)
	check(err)
	defer session.Close()

	col := session.DB(database).C(collection)

	count, err := col.Count()
	check(err)
	utils.Log.Println(count, " documents found")

	var mgoResult []models.Declarations
	err = col.Find(bson.M{"ANIO":2017}).All(&mgoResult)
	check(err)
	utils.Log.Println(mgoResult)
	merger.Merger(inputPath, "merge-2017.csv", mgoResult)
}