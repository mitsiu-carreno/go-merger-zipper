package main

import(
	"flag"
	"os"
	"time"
	"github.com/mitsiu-carreno/go-merger-zipper/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type declarations struct{
	_id 			string `bson:"_id,omitempty"`
	ANIO			string `bson:"ANIO"`
	INDICE 			string `bson:"INDICE"`
	ACUSE			string `bson:"ACUSE"`
	FECHA			string `bson:"FECHA"`
	DEPENDENCIA		string `bson:"DEPENDENCIA"`
	DIA				string `bson:"DIA"`
	DECLARACION		string `bson:"DECLARACION"`
	SOURCE			string `bson:"SOURCE"`
	MES				string `bson:"MES"`
	ARCHIVO 		string `bson:"ARCHIVO"`
	NOMBRE			string `bson:"NOMBRE"`
	FOLDER			string `bson:"FOLDER"`
}

func check(e error){
	if e != nil{
		utils.Log.Println(e)
		panic(e)
	}
}

func main(){
	var logpath = flag.String("logpath", os.Getenv("LOG_FILE"), "Log Path")
	utils.NewLog(*logpath)

	var (
		hosts 		= os.Getenv("MAIN_DB_HOST")
		database	= os.Getenv("MAIN_DB_DB")
		username	= os.Getenv("MAIN_DB_USER")
		password	= os.Getenv("MAIN_DB_PASSWORD")
		collection	= os.Getenv("MAIN_DB_COLLECTION")
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

	var mgoResult []declarations
	err = col.Find(bson.M{"ANIO":2017}).All(&mgoResult)
	check(err)
	utils.Log.Println(mgoResult)
}