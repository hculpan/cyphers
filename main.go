package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/atotto/clipboard"

	cypher "github.com/hculpan/cyphers/cyphers"
	"github.com/hculpan/cyphers/types"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

//go:generate go run scripts/includecyphers.go

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

var unique bool = false
var fullOutput bool = false
var onClipboard bool = true
var useSubtleCyphers bool = false
var useNumeneraCyphers bool = true
var outputRolls bool = false

var selectedList int
var cypherCount int = 1

var outputEntry *widget.Label

type labelLayout struct {
}

func main() {
	a := app.New()
	w := a.NewWindow("Cyphers")

	radio := widget.NewRadioGroup([]string{"Numenera", "Subtle"}, func(c string) {
		switch c {
		case "Numenera":
			selectedList = cypher.NUMENERA_CYPHERS
		case "Subtle":
			selectedList = cypher.SUBTLE_CYPHERS
		}
	})
	radio.SetSelected("Numenera")
	radio.Horizontal = true

	countLabel := widget.NewLabel("How many cyphers: 1")
	countSlider := widget.NewSlider(1.0, 10.0)
	countSlider.OnChanged = func(v float64) {
		countLabel.Text = fmt.Sprintf("How many cyphers: %d", int(v))
		cypherCount = int(v)
	}

	outputEntry = widget.NewLabelWithStyle("", fyne.TextAlignLeading, fyne.TextStyle{})
	outputEntry.Wrapping = fyne.TextWrapWord

	vbox := widget.NewVBox(
		widget.NewHBox(layout.NewSpacer(), radio, layout.NewSpacer()),
		widget.NewHBox(layout.NewSpacer(), countLabel, layout.NewSpacer()),
		fyne.NewContainerWithLayout(layout.NewMaxLayout(), countSlider),
		widget.NewHBox(layout.NewSpacer(), widget.NewButton("Generate", func() {
			generateCyphers()
		}), layout.NewSpacer(), widget.NewButton("->", func() {})),
		fyne.NewContainerWithLayout(&labelLayout{}, outputEntry),
	)
	w.SetContent(vbox)

	w.Resize(fyne.NewSize(400, 600))
	w.CenterOnScreen()
	w.SetFixedSize(true)

	w.ShowAndRun()
}

func (l *labelLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 0)
}

func (l *labelLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	objects[0].Resize(fyne.NewSize(350, 200))
	objects[0].Move(fyne.NewPos(0, 0))
}

func generateCyphers() {
	clipText := ""
	for i := 0; i < cypherCount; i++ {
		if v, err := cypher.GetCyphers(selectedList, outputRolls); err == nil {
			if fullOutput {
				clipText += buildCypherOutputForConsole(v, v.ActualLevel, v.Effect)
			} else {
				clipText += buildCypherOutputForClipboard(v, v.ActualLevel, v.Effect)
			}
		} else {
			fmt.Printf("Error: %s\n", err.Error())
			break
		}
	}

	if clipText != "" {
		clipboard.WriteAll(clipText)
	}

	outputEntry.SetText(clipText)
}

func buildCypherOutputForClipboard(c *types.Cypher, lvl int, effect string) string {
	return buildCypherOutput(c, lvl, strings.ReplaceAll(effect, "\n", " "), true)
}

func buildCypherOutputForConsole(c *types.Cypher, lvl int, effect string) string {
	return buildCypherOutput(c, lvl, effect, false)
}

func buildCypherOutput(c *types.Cypher, lvl int, effect string, roll20Formatting bool) string {
	var result string
	if roll20Formatting {
		result = fmt.Sprintf("**%s : Level %d**\n", c.Name, lvl)
	} else {
		result = fmt.Sprintf("%s : Level %d\n", c.Name, lvl)
	}

	if len(c.Type) > 0 {
		result += fmt.Sprintf("%s\n", c.Type[rand.Int()%len(c.Type)])
	}

	result += effect + "\n"

	return result
}
