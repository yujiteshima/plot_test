package plot

import (
	"encoding/csv"
	"image/color"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func plotPointsSepal(x string) plotter.XYs {
	pts := make(plotter.XYs, 150)

	file, err := os.Open("./iris.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var line []string

	for i := range pts {
		line, err = reader.Read()
		if err != nil {
			panic(err)
		}
		if line[4] == x {
			pts[i].X, _ = strconv.ParseFloat(line[0], 64)
			pts[i].Y, _ = strconv.ParseFloat(line[1], 64)
		}
	}
	return pts
}

func PlotSepal() {

	// 図の生成
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	//label
	p.Title.Text = "Sepal"
	p.X.Label.Text = "length"
	p.Y.Label.Text = "width"
	// 補助線
	p.Add(plotter.NewGrid())

	x1 := "Setosa"
	x2 := "Versicolor"
	x3 := "Virginica"

	// 散布図の作成
	plot1, err := plotter.NewScatter(plotPointsSepal(x1))
	if err != nil {
		panic(err)
	}

	plot2, err := plotter.NewScatter(plotPointsSepal(x2))
	if err != nil {
		panic(err)
	}
	plot3, err := plotter.NewScatter(plotPointsSepal(x3))
	if err != nil {
		panic(err)
	}

	//色を指定する．
	plot1.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 55}
	plot2.GlyphStyle.Color = color.RGBA{R: 155, B: 128, A: 255}
	plot3.GlyphStyle.Color = color.RGBA{R: 55, B: 255, A: 128}
	//plot1,plot2をplot
	p.Add(plot1)
	p.Add(plot2)
	p.Add(plot3)

	//label
	p.Legend.Add("Seotsa", plot1)
	p.Legend.Add("Versicolor", plot2)
	p.Legend.Add("Virginica", plot3)

	// 座標範囲
	p.X.Min = 0
	p.X.Max = 10
	p.Y.Min = 0
	p.Y.Max = 10

	// plot.pngに保存
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "plotSepal.png"); err != nil {
		panic(err)
	}
}
