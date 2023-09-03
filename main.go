package main

import (
	"math/rand"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("PassWord MINE")
	var pw string

	input := widget.NewEntry()
	input.SetPlaceHolder("Ex )  my pass word is lucky - 777. Generate !")
	out := widget.NewLabel(`Input charactors with SPACE, then Press ENTER .`)
	input.OnSubmitted = func(text string) {
		pw = generator(text)
		out.SetText(pw)
	}
	copybtn := widget.NewButtonWithIcon("Copy", theme.ContentCopyIcon(), func() {
		w.Clipboard().SetContent(out.Text)
	})

	w.SetContent(container.NewVBox(out, input, copybtn))

	// desktop menu
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("PassWord MINE",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}
	w.SetCloseIntercept(func() {
		w.Hide()
	})
	// Do not quit by closing window ↑

	w.Resize(fyne.NewSize(300, 100))
	w.ShowAndRun()
}

func generator(input string) string {
	charactors := strings.Fields(input)
	rand.Shuffle(len(charactors), func(i, j int) {
		charactors[i], charactors[j] = charactors[j], charactors[i]
	})
	pw := strings.Join(charactors, "")
	return pw
}
