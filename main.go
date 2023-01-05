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

	// wm, err = api.TextWatermark("test 1",
	// 	"font:Courier, points:48, col: 1 0 0, rot:0, sc:1 abs",
	// 	onTop,
	// 	false,
	// 	unit)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// api.AddWatermarksFile(inFile, outFile, []string{"1"}, wm, nil)

	// pdfcpu stamp add -mode text -- "Some multi\nline text"
	// "ma:5, bo:7 round .3 .7 .7, fillc:#3277d3, bgcol:#beded9, rot:0" in.pdf out.pdf
	// "font:Helvetica, points:24, rtl:off, sc:0.5 rel, pos:c, off:0 0, align:c, fillc:#808080, strokec:#808080, rot:0, op:1, mo:0, ma:5, bo:7 round .3 .7 .7, bgcol:#beded9"

	// wm, err = api.TextWatermark("test 2 test 2",
	// 	"ma:0, bo:0, rot:0",
	// 	onTop,
	// 	false,
	// 	unit)

	if err != nil {
		log.Fatal(err)
	}

	//api.AddWatermarksFile(inFile, outFile, []string{"1"}, wm, nil)

	wm, err = api.TextWatermark("test 2 test 2"+"\n"+"dop str",
		"font:Helvetica, points:24, rtl:off, sc:0.5 rel, pos:c, off:0 0, align:c, fillc:#808080, strokec:#808080, rot:0, op:1, mo:0, ma:5, bo:7 round .3 .7 .7, bgcol: 1.0 1.0 1.0",
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

	// pb, err := api.Box("20% 40%", unit)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//api.Crop(Readers, Writes, []string{"1"}, pb, nil)

	defer Readers.Close()
	defer Writes.Close()

}
