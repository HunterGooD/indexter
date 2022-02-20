package app

import "os"

type App struct {
	homePath string
	ui       *UI
}

func NewApp() *App {
	ui := NewUI()

	return &App{ui: ui}
}

func (a *App) Run() error {
	a.ui.Init()
	err := a.ui.StartApp()
	if err != nil {
		return err
	}
	dir, err := os.UserHomeDir()
	if err != nil {
		// TODO: Modal window error get home dir
		a.ui.ShowModal("Не удается получить домашнею директорию", []string{"Закрыть"})
	}
	a.homePath = dir
	//TODO: a.ui.setDir
	return nil
}
