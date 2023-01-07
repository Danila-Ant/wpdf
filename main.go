package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func main() {

	var err error
	var inFile = "in.pdf"
	var outFile = "out.pdf"

	wm := pdfcpu.DefaultWatermarkConfig()
	unit := pdfcpu.POINTS

	onTop := true

	if err != nil {
		log.Fatal(err)
	}

	wm, err = api.TextWatermark("test 2 test 2"+"\n"+"dop str",
		"font:Helvetica, points:24, rtl:off, sc:0.5 rel, pos:br, off:0 0, align:c, fillc: 0 0 0, strokec: 0 0 0, rot:0, op:1, mo:0, ma:5, bo:7 round .0 .0 .0, bgcol: 1.0 1.0 1.0",
		onTop,
		false,
		unit)

	if err != nil {
		log.Fatal(err)
	}

	Readers, err := os.Open(inFile)
	if err != nil {
		fmt.Println(err)
	}

	Writes, err := os.Create(outFile)
	if err != nil {
		fmt.Println(err)
	}

	api.AddWatermarks(Readers, Writes, []string{"1"}, wm, nil)

	defer Readers.Close()
	defer Writes.Close()

}
