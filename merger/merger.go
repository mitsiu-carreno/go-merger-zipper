package merger

import (
	"io"
	"fmt"
	"encoding/csv"
	"os"
	models "github.com/mitsiu-carreno/go-merger-zipper/declarations"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func getRow(file io.Reader) (ch chan[]string){
	ch = make(chan []string, 10)
	go func(){
		reader := csv.NewReader(file)
		_, err := reader.Read()
		check(err)
		defer close(ch)

		for{
			record, err := reader.Read()
			if err == io.EOF{
				break
			}
			check(err)
			ch <- record
		}
	}()
	return
}

// Merger receives a list of csv files to merge into a single file
func Merger(inputPath string, filename string, files[]models.Declarations){
	outfile, err := os.Create("./" + filename)
	check(err)
	defer outfile.Close()

	writter := csv.NewWriter(outfile)

	err = writter.Write([]string{"OBSERVACIONES","INDICE","NOMBRE","DEPENDENCIA","DECLARACION","FECHA","ACUSE","TEMA","SUBTEMA","VALOR"})
	check(err)

	for _, entry := range files{
		fmt.Println(entry.ARCHIVO)

		file, err := os.Open(inputPath + entry.ARCHIVO)
		check(err)
		defer file.Close()

		for rec := range getRow(file){
			err = writter.Write(rec)
			check(err)
		}

		writter.Flush()
		err = writter.Error()
		check(err)
	}
}