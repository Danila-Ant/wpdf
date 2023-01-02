package main

import (
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {

	if len(os.Args) < 3 {
		log.Printf("\nUsage:\n\t%s <inputfile> <output-dir>\n\n", os.Args[0])
	}

	inputFile := os.Args[1]
	outputDir := os.Args[2]

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Unable to open file '%s': %v", inputFile, err)
	}

	err = api.ExtractImages(f, outputDir, nil, nil)
	if err != nil {
		log.Fatalf("Unable to extract images from '%s' into '%s': %v", inputFile, outputDir, err)
	}
}
