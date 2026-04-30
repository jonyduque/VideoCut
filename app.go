package main

import (
	"context"
	"os"
	"path/filepath"
)

// App struct
type App struct {
	ctx           context.Context
	executableDir string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	dir, err := a.getExecutableDir()
	if err == nil {
		a.executableDir = dir
	}
}

// executableFunc allows mocking os.Executable in tests
var executableFunc = os.Executable

// getExecutableDir returns the directory of the executable
func (a *App) getExecutableDir() (string, error) {
	ex, err := executableFunc()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}

// IsEnvironmentReady checks if ffmpeg and ffprobe are in the executable directory
func (a *App) IsEnvironmentReady() bool {
	ffmpegPath := filepath.Join(a.executableDir, "ffmpeg.exe")
	ffprobePath := filepath.Join(a.executableDir, "ffprobe.exe")

	if _, err := os.Stat(ffmpegPath); err != nil {
		return false
	}
	if _, err := os.Stat(ffprobePath); err != nil {
		return false
	}
	return true
}
