package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
			cmd := exec.Command(name, arg...)
			out, err := cmd.CombinedOutput()
			return out, err
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

	if a.ctx != nil {
		runtime.LogInfof(a.ctx, "Running ffprobe: %s %v", ffprobePath, args)
	}

	out, err := a.runCommand(ffprobePath, args...)
	if err != nil {
		if a.ctx != nil {
			runtime.LogErrorf(a.ctx, "ffprobe failed (out: %s): %v", string(out), err)
		}
		return 0, fmt.Errorf("ffprobe failed (output: %s): %v", string(out), err)
	}

	durationStr := strings.TrimSpace(string(out))
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse duration '%s': %v", durationStr, err)
	}

	return duration, nil
}

// SplitVideo splits the video into two parts with overlap
func (a *App) SplitVideo(filePath string, cutPoint float64, overlap float64, outputDir string, part1Name string, part2Name string) error {
	ffmpegPath := filepath.Join(a.executableDir, "ffmpeg.exe")
	if _, err := os.Stat(ffmpegPath); err != nil {
		ffmpegPath = "ffmpeg"
	}

	if outputDir == "" {
		outputDir = filepath.Dir(filePath)
	}

	// Ensure output names have extension
	if !strings.HasSuffix(strings.ToLower(part1Name), ".mp4") {
		part1Name += ".mp4"
	}
	if !strings.HasSuffix(strings.ToLower(part2Name), ".mp4") {
		part2Name += ".mp4"
	}

	out1 := filepath.Join(outputDir, part1Name)
	out2 := filepath.Join(outputDir, part2Name)

	// Part 1: 0 to cutPoint + overlap
	args1 := []string{
		"-i", filePath,
		"-ss", "0",
		"-to", fmt.Sprintf("%f", cutPoint+overlap),
		"-c:v", "libx264", "-preset", "fast", "-crf", "22", "-c:a", "aac",
		"-y", out1,
	}

	out1_res, err := a.runCommand(ffmpegPath, args1...)
	if err != nil {
		if a.ctx != nil {
			runtime.LogErrorf(a.ctx, "Part 1 failed (out: %s): %v", string(out1_res), err)
		}
		return fmt.Errorf("part 1 failed (output: %s): %v", string(out1_res), err)
	}

	// Part 2: cutPoint - overlap to end
	start2 := cutPoint - overlap
	if start2 < 0 {
		start2 = 0
	}
	args2 := []string{
		"-i", filePath,
		"-ss", fmt.Sprintf("%f", start2),
		"-c:v", "libx264", "-preset", "fast", "-crf", "22", "-c:a", "aac",
		"-y", out2,
	}

	out2_res, err := a.runCommand(ffmpegPath, args2...)
	if err != nil {
		if a.ctx != nil {
			runtime.LogErrorf(a.ctx, "Part 2 failed (out: %s): %v", string(out2_res), err)
		}
		return fmt.Errorf("part 2 failed (output: %s): %v", string(out2_res), err)
	}

	return nil
}

// GetFFmpegVersion returns the version of FFmpeg/ffprobe
func (a *App) GetFFmpegVersion() (string, error) {
	ffprobePath := filepath.Join(a.executableDir, "ffprobe.exe")
	if _, err := os.Stat(ffprobePath); err != nil {
		ffprobePath = "ffprobe"
	}

	out, err := a.runCommand(ffprobePath, "-version")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// SelectFile opens a dialog to select a video file
func (a *App) SelectFile() (string, error) {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Selecionar Vídeo",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Vídeos",
				Pattern:     "*.mp4;*.mkv;*.avi;*.mov;*.wmv;*.flv",
			},
		},
	})
	if err != nil {
		return "", err
	}
	return selection, nil
}

// SelectDirectory opens a dialog to select a directory
func (a *App) SelectDirectory() (string, error) {
	selection, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Selecionar Pasta de Destino",
	})
	if err != nil {
		return "", err
	}
	return selection, nil
}


