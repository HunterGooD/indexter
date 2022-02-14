package app

import "os"

type App struct {
	homePath string
	ui       *UI
}

func NewApp() *App {
	ui := NewUI()
	ui.Init()
	ui.StartApp()

	dir, err := os.UserHomeDir()
	if err != nil {
		// TODO: Modal window error get home dir
		ui.ShowModal("Не удается получить домашнею директорию", []string{"Закрыть"})
	}
	return &App{dir, ui}
}
