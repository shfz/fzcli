package ui

import (
	"log"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/hi120ki/clf/model"
)

var g *widgets.Gauge
var p *widgets.Paragraph
var l *widgets.List

func Init() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	g = widgets.NewGauge()
	g.Title = "Progress"
	g.Percent = 0
	g.SetRect(0, 0, 80, 3)
	g.BarColor = ui.ColorWhite
	g.BorderStyle.Fg = ui.ColorWhite
	g.TitleStyle.Fg = ui.ColorCyan

	p = widgets.NewParagraph()
	p.Title = "Result"
	p.Text = "Progress : 0 (0%)\nSuccess : 0 (0%)\nFailure : 0 (0%)"
	p.SetRect(0, 3, 80, 8)
	p.BorderStyle.Fg = ui.ColorWhite
	p.TitleStyle.Fg = ui.ColorCyan

	l = widgets.NewList()
	l.Title = "Log"
	l.Rows = []string{}
	l.SetRect(0, 8, 80, 24)
	l.TitleStyle.Fg = ui.ColorCyan

	ui.Render(g, p, l)
}

func Close() {
	ui.Close()
}

func Update(r model.Result) {
	g.Percent = int(100 * (r.Success + r.Failure) / r.Total)
	p.Text = "Progress : " + strconv.FormatUint(r.Success+r.Failure, 10) + " (" + strconv.Itoa(int(100*(r.Success+r.Failure)/r.Total)) + "%)\nSuccess  : " + strconv.FormatUint(r.Success, 10) + " (" + strconv.Itoa(int(100*r.Success/r.Total)) + "%)\nFailure  : " + strconv.FormatUint(r.Failure, 10) + " (" + strconv.Itoa(int(100*r.Failure/r.Total)) + "%)"
	if len(r.Message) > 14 {
		l.Rows = r.Message[len(r.Message)-14:]
	} else {
		l.Rows = r.Message
	}
	ui.Render(g, p, l)
}
