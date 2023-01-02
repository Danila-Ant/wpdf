package main

import (
	"github.com/jung-kurt/gofpdf"
)


func main() {

	///var pts []gofpdf.PointType

	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0
	pdf.AddPage()
	//pdf.SetDrawColor(50, 50, 50)
	//pdf.SetFont("Arial", "B", 16)

	pts := []gofpdf.PointType{
		{X: 0, Y: 0},
		{X: 0, Y: 100},
		{X: 100, Y: 100},
		{X: 100, Y: 0},
	}

	pdf.Polygon(pts, "D")

	fileStr := "hello.pdf"
	//pdf.OpenLayerPane()
	//pdf.OutputFileAndClose(fileStr)
	//example.Summary(err, fileStr)
	pdf.Output()
}
