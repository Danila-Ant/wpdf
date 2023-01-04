package main

import (
	"log"

	api "github.com/pdfcpu/pdfcpu/pkg/api"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func main() {

	var err error
	var inFile = "in.pdf"
	var outFile = "out.pdf"

	wm := pdfcpu.DefaultWatermarkConfig()
	unit := pdfcpu.POINTS

	onTop := true // we are testing stamps

	wm, err = api.TextWatermark("test texst",
		"font:Courier, points:48, col: 1 0 0, rot:0, sc:1 abs",
		onTop,
		false,
		unit)

	if err != nil {
		log.Fatal(err)
	}

	api.AddWatermarksFile(inFile, outFile, []string{"1"}, wm, nil)

}
