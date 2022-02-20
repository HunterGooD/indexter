package app

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UI struct {
	app    *tview.Application
	layout *tview.Flex     // расположение основных деталей интерфейса
	pages  *tview.Pages    // основные страницы с объктами
	info   *tview.TextView // оглавление страниц

	directoryTree *tview.TreeView
}

type Slide func(nextSlide func()) (title string, content tview.Primitive)

func NewUI() *UI {
	return &UI{}
}

func (ui *UI) Init() {

	var slides = []Slide{
		ui.MainPage,
	}

	ui.app = tview.NewApplication()
	ui.pages = tview.NewPages()

	ui.info = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			ui.pages.SwitchToPage(added[0])
		})
	ui.info.SetTextColor(tcell.Color101)
	previousSlide := func() {
		slide, _ := strconv.Atoi(ui.info.GetHighlights()[0])
		slide = (slide - 1 + len(slides)) % len(slides)
		ui.info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(ui.info.GetHighlights()[0])
		slide = (slide + 1) % len(slides)
		ui.info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	for index, slide := range slides {
		title, primitive := slide(nextSlide)
		ui.pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
		fmt.Fprintf(ui.info, `%d ["%d"][cyan]%s[white][""]  `, index+1, index, title)
	}
	ui.info.Highlight("0")
	ui.DefaultLayaout()

	ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextSlide()
			return nil
		} else if event.Key() == tcell.KeyCtrlP {
			previousSlide()
			return nil
		}
		return event
	})
}

func (ui *UI) StartApp() error {
	return ui.app.SetRoot(ui.layout, true).EnableMouse(true).Run()
}

func (ui *UI) ShowModal(text string, buttons []string) {
	ui.layout.Clear()

	var modal *tview.Modal
	modal = tview.NewModal().
		SetText(text).
		AddButtons(buttons).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Close" || buttonLabel == "Закрыть" || buttonLabel == "Ок" || buttonLabel == "" {
				ui.layout.Clear()
				ui.DefaultLayaout()
			}
		})
	ui.layout.AddItem(modal, 1, 1, true)
}

func (ui *UI) Center(width, height int, p tview.Primitive) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}

func (ui *UI) DefaultLayaout() {
	ui.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(ui.pages, 0, 1, false).
		AddItem(ui.info, 1, 1, false)
}
