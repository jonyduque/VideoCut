package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// CommandRunner defines a function signature for running commands
type CommandRunner func(name string, arg ...string) ([]byte, error)

// App struct
type App struct {
	ctx           context.Context
	executableDir string
	runCommand    CommandRunner
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		runCommand: func(name string, arg ...string) ([]byte, error) {
			return exec.Command(name, arg...).Output()
		},
	}
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

// GetVideoDuration returns the duration of the video file in seconds
func (a *App) GetVideoDuration(filePath string) (float64, error) {
	ffprobePath := filepath.Join(a.executableDir, "ffprobe.exe")
	
	// If in development mode and binaries aren't next to exe, try to find them in PATH or current dir
	if _, err := os.Stat(ffprobePath); err != nil {
		ffprobePath = "ffprobe" // Fallback to system path
	}

	args := []string{
		"-v", "error",
		"-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1",
		filePath,
	}

	out, err := a.runCommand(ffprobePath, args...)
	if err != nil {
		return 0, err
	}

	durationStr := strings.TrimSpace(string(out))
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse duration '%s': %v", durationStr, err)
	}

	return duration, nil
}

