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
func Merger(inputPath string, outputPath string, filename string, files[]models.Declarations){
	//var total = len(files)

	// Check if output directory exists
	_, err := os.Stat(outputPath)
	if os.IsNotExist(err){
		os.MkdirAll(outputPath, os.ModePerm)
	}else{
		utils.Check(err)
	}

	// Create new file
	outfile, err := os.Create(outputPath + filename)
	utils.Check(err)
	defer outfile.Close()

	writter := csv.NewWriter(outfile)

	// Write headers in new file
	err = writter.Write([]string{"OBSERVACIONES","INDICE","NOMBRE","DEPENDENCIA","DECLARACION","FECHA","ACUSE","TEMA","SUBTEMA","VALOR"})
	utils.Check(err)

	for _, entry := range files{
		//var entryNum = i+1

		// Check if input file exists
		_, err := os.Stat(inputPath + entry.FOLDER + "/" + entry.ARCHIVO)
		if os.IsNotExist(err){
      // fmt.Println("Merge - File not found -v4", inputPath + entry.FOLDER + "/" + entry.ARCHIVO)
			utils.Log.Print("Merge - File not found -v1", inputPath, entry.FOLDER, inputPath + entry.FOLDER + "/" + entry.ARCHIVO, "\n")
			continue
		}
		utils.Check(err)

		file, err := os.Open(inputPath + entry.FOLDER + "/" + entry.ARCHIVO)
		utils.Check(err)
		defer file.Close()

		for rec := range getRow(file){
			err = writter.Write(rec)
			utils.Check(err)
		}

		// Clean variables
		writter.Flush()
		err = writter.Error()
		utils.Check(err)

		// Closed here to avoid overload mem
		file.Close()
	}
}
