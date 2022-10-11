package utils

import (
	"math"
	"sort"
	"time"

	"image/color"

	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

type TimeTester func(n int, tries int) []time.Duration

func Median(data []int) float64 {
	sort.Ints(data)

	var length = len(data)
	var midpoint = float64(length) / 2

	var floored = int(math.Floor(midpoint))
	var ceiled = int(math.Ceil(midpoint))

	if floored == ceiled {
		return float64(data[floored])
	} else {
		return float64(data[floored]+data[floored]) / 2
	}
}

func Average(data []int) float64 {
	var sum = 0
	for _, elem := range data {
		sum += elem
	}

	return float64(sum) / float64(len(data))
}

func Outliers(data []int) (minimum float64, maximum float64) {
	if len(data) == 0 {
		panic("expected data to contain elements")
	}

	var min = data[0]
	var max = data[0]

	for _, elem := range data {
		if elem > max {
			max = elem
		}

		if elem < min {
			min = elem
		}
	}

	return float64(min), float64(max)
}

// type timeUnit int

// const (
// 	Nanoseconds timeUnit = iota
// 	Milliseconds
// )

type Result struct {
	n       int
	tries   int
	timeMin float64
	timeMax float64
	timeAvg float64
	timeMed float64
}

type Results []Result

func GetResults(TimeTester_func TimeTester, tries int, n_samples []int) Results {
	results := make(Results, len(n_samples))

	for index, n := range n_samples {
		var nanoseconds []int
		for _, elem := range TimeTester_func(n, tries) {
			nanoseconds = append(nanoseconds, int(elem.Nanoseconds()))
		}

		avg := Average(nanoseconds)
		med := Median(nanoseconds)
		min, max := Outliers(nanoseconds)

		results[index] = Result{
			n,
			tries,
			min,
			max,
			avg,
			med,
		}
	}

	return results
}

// type FunctionFit int

// const (
// 	None FunctionFit = iota
// 	N
// 	LogN
// 	NLogN
// 	NSquared
// 	NCubed
// )

// type ComplexityPlot struct {
// 	data []triesResult
// }

// //?

// func NewComplexityPlot(data []triesResult) ComplexityPlot {
// 	return ComplexityPlot{
// 		data,
// 	}
// }

// func MakeComplexityPlot(data []triesResult) *ComplexityPlot {
// 	a := NewComplexityPlot(data)
// 	return &a
// }

// func (cp *ComplexityPlot) IntoScatter(){

// 	for index, element := range cp.data {
// 		toPlot["average"].data[index] = plotter.XY{X: float64(element.n), Y: element.timeAvg}
// 	}
// }

// func (cp *ComplexityPlot) Plot(fileName string, title string) {

// 	p := plot.New()

// 	type dataAndStyle struct {
// 		data            plotter.XYs
// 		styleShape      draw.GlyphDrawer
// 		styleBrightness int
// 	}

// 	toPlot := make(map[string]dataAndStyle)

// 	if cp.average {
// 		toPlot["average"] = dataAndStyle{
// 			data:            make(plotter.XYs, len(cp.data)),
// 			styleShape:      draw.CircleGlyph{},
// 			styleBrightness: 50,
// 		}

// 		for index, element := range cp.data {
// 			toPlot["average"].data[index] = plotter.XY{X: float64(element.n), Y: element.timeAvg}
// 		}
// 	}

// 	if cp.median {
// 		toPlot["median"] = dataAndStyle{
// 			data:            make(plotter.XYs, len(cp.data)),
// 			styleShape:      draw.PlusGlyph{},
// 			styleBrightness: 50,
// 		}

// 		for index, element := range cp.data {
// 			toPlot["median"].data[index] = plotter.XY{X: float64(element.n), Y: element.timeMed}
// 		}
// 	}

// 	if cp.maximum {
// 		toPlot["maximum"] = dataAndStyle{
// 			data:            make(plotter.XYs, len(cp.data)),
// 			styleShape:      draw.PyramidGlyph{},
// 			styleBrightness: 50,
// 		}

// 		for index, element := range cp.data {
// 			toPlot["maximum"].data[index] = plotter.XY{X: float64(element.n), Y: element.timeMax}
// 		}
// 	}

// 	if cp.minimum {
// 		toPlot["minimum"] = dataAndStyle{
// 			data:            make(plotter.XYs, len(cp.data)),
// 			styleShape:      draw.TriangleGlyph{},
// 			styleBrightness: 50,
// 		}

// 		for index, element := range cp.data {
// 			toPlot["minimum"].data[index] = plotter.XY{X: float64(element.n), Y: element.timeMin}
// 		}
// 	}

// 	for name, dataAndStyle := range toPlot {
// 		scatter, err := plotter.NewScatter(dataAndStyle.data)
// 		scatter.GlyphStyle.Shape = dataAndStyle.styleShape
// 		scatter.GlyphStyle.Color = color.RGBA{R: 255, A: 255}
// 		if err != nil {
// 			panic("Failed to create scatter from data")
// 		}
// 		p.Add(scatter)
// 		p.Legend.Add(name, scatter)
// 	}

// 	p.Title.Text = title
// 	p.X.Label.Text = "n"
// 	p.Y.Label.Text = "runtime (ns)"
// 	p.Save(5*vg.Inch, 5*vg.Inch, fileName)
// }

type ScatterType string

const (
	Min ScatterType = "min"
	Max ScatterType = "max"
	Avg ScatterType = "average"
	Med ScatterType = "median"
)

type ScatterColor int

const (
	Red ScatterColor = iota
	Green
	Blue
)

func (r Results) Scatter(scatterType ScatterType, glyph draw.GlyphDrawer, scatterColor ScatterColor) *plotter.Scatter {
	scatterData := make(plotter.XYs, len(r))

	for index, element := range r {
		scatterData[index] = plotter.XY{X: float64(element.n), Y: element.timeAvg}

		switch scatterType {
		case Min:
			scatterData[index] = plotter.XY{X: float64(element.n), Y: element.timeMin}
		case Max:
			scatterData[index] = plotter.XY{X: float64(element.n), Y: element.timeMax}
		case Avg:
			scatterData[index] = plotter.XY{X: float64(element.n), Y: element.timeAvg}
		case Med:
			scatterData[index] = plotter.XY{X: float64(element.n), Y: element.timeMed}
		}
	}

	p, err := plotter.NewScatter(scatterData)

	if err != nil {
		panic("faild to create scatter")
	}

	p.GlyphStyle.Shape = glyph

	switch scatterColor {
	case Red:
		p.GlyphStyle.Color = color.RGBA{R: 255, A: 255}
	case Green:
		p.GlyphStyle.Color = color.RGBA{G: 255, A: 255}
	case Blue:
		p.GlyphStyle.Color = color.RGBA{B: 255, A: 255}
	}

	return p
}

func IntLinspace(n int, min int, max int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = ((i * max) / n) + 1
	}

	return arr
}

// type FitType int

// const (
// 	N FitType = iota
// 	LogN
// 	NLogN
// 	NSquared
// )

// func getFit(data []float64, fitType FitType) func(n float64){
// 	scoreFunction := func(){
// 		//
// 	}
// } IS THIS NEEDED???? +M will be tiny???
