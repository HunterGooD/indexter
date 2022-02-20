package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type node struct {
	text     string
	expand   bool
	selected func()
	children []*node
}

func (ui *UI) folderTree(nextSlide func()) (string, *tview.Primitive) {
	input := tview.NewInputField().
		SetLabel("Путь к директории: ") //.
		// SetDoneFunc(func(key tcell.Key) {
	input.SetText("") // обновлять директории по кликам в дереве
	// })
	// tree = tview.NewTreeView()
	return "Директория", nil
}

func add(target *node) *tview.TreeNode {
	node := tview.NewTreeNode(target.text).
		SetSelectable(target.expand || target.selected != nil).
		SetExpanded(true).
		SetReference(target)
	if target.expand {
		node.SetColor(tcell.ColorGreen)
	} else if target.selected != nil {
		node.SetColor(tcell.ColorRed)
	}
	for _, child := range target.children {
		node.AddChild(add(child))
	}
	return node
}

func (ui *UI) setTreeDir(path string) {

}
