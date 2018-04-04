package zipper

import (
	"io"
	"os"
	"archive/zip"
	"github.com/mitsiu-carreno/go-merger-zipper/utils"

)

// Zipper function recieves a list of files to zip and a string to name the zip file created
func Zipper(inputPath string, outputPath string, filename string, files[]string) {

	// Check if output directory exists
	_, err := os.Stat(outputPath)
	if os.IsNotExist(err){
		os.MkdirAll(outputPath, os.ModePerm)
	}else{
		utils.Check(err)
	}

	// Create new file
	newFile, err := os.Create(outputPath + filename)
	utils.Check(err)
	defer newFile.Close()
	utils.Log.Println("Zip file: " + filename + " created")

	zipWritter := zip.NewWriter(newFile)
	defer zipWritter.Close()

	// Add files to zip
	for _, entry := range files{

		// Open zip file
		zipfile, err := os.Open(inputPath + entry)
		utils.Check(err)
		defer zipfile.Close()

		// Check if input file exists and get info
		info, err := os.Stat(inputPath + entry)
		if os.IsNotExist(err){
			utils.Log.Println("Zip - File not found ", entry)
			continue
		}
		utils.Check(err)

		header, err := zip.FileInfoHeader(info)
		utils.Check(err)

		// Deflate offers a better compression
		header.Method = zip.Deflate

		// Set headers to writter
		writter, err := zipWritter.CreateHeader(header)
		utils.Check(err)

		// Zip file
		_, err = io.Copy(writter, zipfile)
		utils.Check(err)

		zipfile.Close()
	}
}