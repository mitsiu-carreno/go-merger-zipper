package main

import(
	"flag"
	"os"
	"time"
	"fmt"
	"strconv"
	"github.com/mitsiu-carreno/go-merger-zipper/utils"
	"github.com/mitsiu-carreno/go-merger-zipper/merger"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	models "github.com/mitsiu-carreno/go-merger-zipper/declarations"
)

func main(){
	// CONFIG LOG
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
	utils.Check(err)
	defer session.Close()

	col := session.DB(database).C(collection)

	var mgoDistinctYears []int
	var mgoDistinctDependencies []string

	// Get all years and dependencies
	err = col.Find(nil).Distinct("ANIO", &mgoDistinctYears)
	utils.Check(err)

	err = col.Find(nil).Distinct("DEPENDENCIA", &mgoDistinctDependencies)
	utils.Check(err)

	// Generate csv and zip by year
	for _, year := range mgoDistinctYears{

		var newFileName = "Decl-" + strconv.Itoa(year)
		var outputPath = "./output/csv/annual/"

		fmt.Println("Merging", newFileName)
		utils.Log.Println("Merging", newFileName)

		var mgoCsvsResult []models.Declarations
		err = col.Find(bson.M{"ANIO" : year}).All(&mgoCsvsResult)
		utils.Check(err)

		utils.Log.Println(len(mgoCsvsResult), " documents to be merged")
		merger.Merger(inputPath, outputPath, newFileName + ".csv", mgoCsvsResult)
		
	}
	/*
	var mgoCsvsResult []models.Declarations
	err = col.Find(bson.M{"ANIO":2017}).All(&mgoResult)
	utils.Check(err)
	utils.Log.Println(len(mgoCsvsResult), " documents to be merged")
	merger.Merger(inputPath, newFileName + ".csv", mgoCsvsResult)

	zipper.Zipper(newFile + ".zip", )
	*/
}