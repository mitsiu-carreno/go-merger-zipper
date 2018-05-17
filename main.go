package main

import(
	"flag"
	"os"
	"time"
	"sort"
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
	var logpath = flag.String("logpath", "./main.log", "Log Path")
	utils.NewLog(*logpath)

	// DATABASE
	var (
		hosts 		= os.Getenv("MAIN_DB_HOST")
		database	= os.Getenv("MAIN_DB_DB")
		username	= os.Getenv("MAIN_DB_USER")
		password	= os.Getenv("MAIN_DB_PASSWORD")
		collection	= os.Getenv("MAIN_DB_COLLECTION")
    	inputPath 	= os.Getenv("SCP_REMOTE_PATH")
    	playgroundPath  = os.Getenv("PLAYGROUND_PATH")

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

	// sort.Ints(mgoDistinctYears)
	sort.Sort(sort.Reverse(sort.IntSlice(mgoDistinctYears)))

	// Generate csv and zip by year
	for _, year := range mgoDistinctYears{

		var fileName = "Declaraciones_" + strconv.Itoa(year)
		var mergedPath = playgroundPath + "go-merge/output/csv/annual/"
		var zippedPath = playgroundPath + "go-merge/output/zip/annual/"

		var files []models.Declarations
		err = col.Find(bson.M{"ANIO" : year}).All(&files)
		utils.Check(err)

		if len(files) > 0 {
			mergeZip(fileName, inputPath, files, mergedPath, zippedPath)
		}else{
			fmt.Println("Skipping (No documents found)", fileName)
		}

		err = col.Find(bson.M{"ANIO": year}).Distinct("DEPENDENCIA", &mgoDistinctDependencies)
		utils.Check(err)

		sort.Strings(mgoDistinctDependencies)

		// Generate csv and zip by dependency/year
		for _, dependency := range mgoDistinctDependencies{
			var fileName = "Declaraciones_" + strconv.Itoa(year) + "_" + dependency
			var mergedPath = playgroundPath + "go-merge/output/csv/dependency/"+strconv.Itoa(year)+"/"
			var zippedPath = playgroundPath + "go-merge/output/zip/dependency/"+strconv.Itoa(year)+"/"

			var files []models.Declarations
			err = col.Find(bson.M{"ANIO" : year, "DEPENDENCIA": dependency}).All(&files)
			utils.Check(err)

			if len(files) > 0 {
				mergeZip(fileName, inputPath, files, mergedPath, zippedPath)
			}else{
				fmt.Println("Skipping (No documents found)", fileName)
			}

		}
	}
}

func mergeZip(fileName string, inputPath string, files []models.Declarations, mergedPath string, zippedPath string){
  fmt.Println("Working", fileName)

  merger.Merger(inputPath, mergedPath, fileName + ".csv", files)

  zipper.Zipper(mergedPath, zippedPath, fileName + ".zip", []string{fileName+".csv"})
}
