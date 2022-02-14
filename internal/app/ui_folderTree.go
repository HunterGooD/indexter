package app

import (
	"github.com/rivo/tview"
)

func (ui *UI) folderTree(nextSlide func()) (string, *tview.Primitive) {
	input := tview.NewInputField().
		SetLabel("Путь к директории: ") //.
		// SetDoneFunc(func(key tcell.Key) {
	input.SetText("") // обновлять директории по кликам в дереве
	// })
	// tree = tview.NewTreeView()
	return "", nil
}
