package main

import(
	"flag"
	"os"
	"time"
	"fmt"
	"strconv"
	"github.com/mitsiu-carreno/go-merger-zipper/utils"
	"github.com/mitsiu-carreno/go-merger-zipper/merger"
	"github.com/mitsiu-carreno/go-merger-zipper/zipper"
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

		var fileName = "Declaraciones_" + strconv.Itoa(year)
		var mergedPath = "./output/csv/annual/"
		var zippedPath = "./output/zip/annual/"

		var files []models.Declarations
		err = col.Find(bson.M{"ANIO" : year}).All(&files)
		utils.Check(err)

		mergeZip(fileName, inputPath, files, mergedPath, zippedPath)

		// Generate csv and zip by dependency/year
		for _, dependency := range mgoDistinctDependencies{
			var fileName = "Declaraciones_" + strconv.Itoa(year) + "_" + dependency
			var mergedPath = "./output/csv/dependency/"+strconv.Itoa(year)+"/"
			var zippedPath = "./output/zip/dependency/"+strconv.Itoa(year)+"/"

			var files []models.Declarations
			err = col.Find(bson.M{"ANIO" : year, "DEPENDENCIA": dependency}).All(&files)
			utils.Check(err)

			mergeZip(fileName, inputPath, files, mergedPath, zippedPath)
		}
	}
}

func mergeZip(fileName string, inputPath string, files []models.Declarations, mergedPath string, zippedPath string){
	fmt.Println("Merging", fileName)
	utils.Log.Println("Merging", fileName)

	utils.Log.Println(len(files), " documents to be merged")
	merger.Merger(inputPath, mergedPath, fileName + ".csv", files)

	utils.Log.Println("----------------------------------------")
	fmt.Println("Zipping", fileName)
	utils.Log.Println("Zipping", fileName)
	zipper.Zipper(mergedPath, zippedPath, fileName + ".zip", []string{fileName+".csv"})
}