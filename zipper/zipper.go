package zipper

import (
	"archive/zip"
	"github.com/mitsiu-carreno/go-merger-zipper/utils"

)

// Zipper function recieves a list of files to zip and a string to name the zip file created
func Zipper(filename string, files[]string) {
	newFile, err := os.Create(filename)
	utils.Check(err)
	defer newFile.Close()

	zipWritter := zip.NewWriter(newFile)
	defer zipWritter.Close()

	// Add files to zip
	
}