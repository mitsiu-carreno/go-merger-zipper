package merger

import (
	"io"
	"encoding/csv"
	"os"
	"github.com/mitsiu-carreno/go-merger-zipper/utils"
	models "github.com/mitsiu-carreno/go-merger-zipper/declarations"
)

func getRow(file io.Reader) (ch chan[]string){
	ch = make(chan []string, 10)
	go func(){
		reader := csv.NewReader(file)
		_, err := reader.Read()
		utils.Check(err)
		defer close(ch)

		for{
			record, err := reader.Read()
			if err == io.EOF{
				break
			}
			utils.Check(err)
			ch <- record
		}
	}()
	return
}

// Merger receives a list of csv files to merge into a single file
func Merger(inputPath string, filename string, files[]models.Declarations){
	var total = len(files)

	outfile, err := os.Create("./" + filename)
	utils.Check(err)
	defer outfile.Close()
	utils.Log.Println("Merge file: " + filename + " created")

	writter := csv.NewWriter(outfile)

	err = writter.Write([]string{"OBSERVACIONES","INDICE","NOMBRE","DEPENDENCIA","DECLARACION","FECHA","ACUSE","TEMA","SUBTEMA","VALOR"})
	utils.Check(err)

	for i, entry := range files{
		var entryNum = i+1
		_, err := os.Stat(inputPath + entry.ARCHIVO)
		if os.IsNotExist(err){
			utils.Log.Print(entryNum, "/", total, ": ", entry.ARCHIVO, " file not found\n")
			continue
		}
		utils.Check(err)

		file, err := os.Open(inputPath + entry.ARCHIVO)
		utils.Check(err)
		defer file.Close()

		for rec := range getRow(file){
			err = writter.Write(rec)
			utils.Check(err)
		}

		writter.Flush()
		err = writter.Error()
		utils.Check(err)

		utils.Log.Print(entryNum, "/", total, ": ", "finished")

		file.Close()
	}
}