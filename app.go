package main

import (
	"context"
	"log"

	"CodeX/backend/utils"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// onStartUpHanlder is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) onStartUpHanlder(ctx context.Context) {
	a.ctx = ctx

	// Ensure Folders and Files are exist
	err := utils.EnsurePaths()
	if err != nil {
		log.Println("Error:", err)
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Error",
			Message: err.Error(),
		})
	}

	// Validate
	err = utils.ValidateFiles()
	if err != nil {
		log.Println("Error:", err)
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Title:   "Error",
			Message: err.Error(),
		})
	}
}
