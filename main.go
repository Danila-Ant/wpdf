package main

import (
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {


	inputFile := "hello.pdf"
	outputDir := ""

	f := os.Open(inputFile)


	var outputDirs []string
	outputDirs[0] = outputDir

	err = api.AddTextWatermarksFile(f, outputDir, , false, "schtamp", "description")
	

}
