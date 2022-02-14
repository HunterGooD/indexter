package app

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const logo = `
	  _          _               __
	 (_)____    / /__  _     __ / /__ __ ______
	/ / __  \__/ / _ \  \   /  /  __/ _ \  ___/
   / / /  / / _ /  __/   \_/  / /_ /  __/  /
  /_/_/  / /___/\___/ /_/  \__\__/ \___/__/
  
`
const (
	subtitle   = `indexter Индексация файлов в указанной директории`
	navigation = `Ctrl-N: Следующая страница Ctrl-P: Предыдущая страница Ctrl-C: Выход`
	mouse      = `(или используйте мышку)`
)

func (ui *UI) MainPage(nextSlide func()) (title string, content tview.Primitive) {

	lines := strings.Split(logo, "\n")
	logoWidth := 0
	logoHeight := len(lines)
	for _, line := range lines {
		if len(line) > logoWidth {
			logoWidth = len(line)
		}
	}
	logoBox := tview.NewTextView().
		SetTextColor(tcell.ColorGreen) //.
		// SetDoneFunc(func(key tcell.Key) {
		// 	nextSlide()
		// })
	fmt.Fprint(logoBox, logo)

	// Create a frame for the subtitle and navigation infos.
	frame := tview.NewFrame(tview.NewBox()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText(subtitle, true, tview.AlignCenter, tcell.ColorWhite).
		AddText("", true, tview.AlignCenter, tcell.ColorWhite).
		AddText(navigation, true, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText(mouse, true, tview.AlignCenter, tcell.ColorDarkMagenta)

	// Create a Flex layout that centers the logo and subtitle.
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 7, false).
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(logoBox, logoWidth+5, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), logoHeight, 1, true).
		AddItem(frame, 0, 10, false)

	return "Начало", flex
}
