package marger

import (
	"encoding/csv"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func Merger(filename string, files[]string){
	outfile, err := os.Create("./" + filename)
	check(err)
	defer outfile.Close()

	writter := csv.NewWriter(outfile)

	err = writter.Write([]string{"OBSERVACIONES","INDICE","NOMBRE","DEPENDENCIA","DECLARACION","FECHA","ACUSE","TEMA","SUBTEMA","VALOR"})
	check(err)

	
}