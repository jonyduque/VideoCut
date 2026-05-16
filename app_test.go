package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetExecutableDir(t *testing.T) {
	app := NewApp()
	dir, err := app.getExecutableDir()
	if err != nil {
		t.Fatalf("Failed to get executable dir: %v", err)
	}

	// In test mode, it should be a valid directory
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Errorf("Directory does not exist: %s", dir)
	}

	// Test error path
	oldFunc := executableFunc
	defer func() { executableFunc = oldFunc }()
	executableFunc = func() (string, error) {
		return "", os.ErrPermission
	}

	_, err = app.getExecutableDir()
	if err == nil {
		t.Error("Expected error from getExecutableDir, got nil")
	}
}

func TestLocateBinaries(t *testing.T) {
	app := NewApp()
	// Create a temp dir for testing
	tempDir, err := os.MkdirTemp("", "videocut_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	app.executableDir = tempDir

	// Should fail initially
	if app.IsEnvironmentReady() {
		t.Error("IsEnvironmentReady should be false when binaries are missing")
	}

	// Create dummy binaries
	ffmpegPath := filepath.Join(tempDir, "ffmpeg.exe")
	ffprobePath := filepath.Join(tempDir, "ffprobe.exe")
	
	os.WriteFile(ffmpegPath, []byte("dummy"), 0755)
	if app.IsEnvironmentReady() {
		t.Error("IsEnvironmentReady should be false when only ffmpeg exists")
	}

	os.Remove(ffmpegPath)
	os.WriteFile(ffprobePath, []byte("dummy"), 0755)
	if app.IsEnvironmentReady() {
		t.Error("IsEnvironmentReady should be false when only ffprobe exists")
	}

	os.WriteFile(ffmpegPath, []byte("dummy"), 0755)
	if !app.IsEnvironmentReady() {
		t.Error("IsEnvironmentReady should be true when both binaries exist")
	}
}

func TestGetVideoDuration(t *testing.T) {
	app := NewApp()
	app.runCommand = func(name string, arg ...string) ([]byte, error) {
		return []byte("123.456\n"), nil
	}

	duration, err := app.GetVideoDuration("test.mp4")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if duration != 123.456 {
		t.Errorf("Expected duration 123.456, got %f", duration)
	}

	// Test error case
	app.runCommand = func(name string, arg ...string) ([]byte, error) {
		return nil, os.ErrNotExist
	}
	_, err = app.GetVideoDuration("nonexistent.mp4")
	if err == nil {
		t.Error("Expected error for nonexistent file, got nil")
	}
}

func TestGetVideoDuration_SpecialChars(t *testing.T) {
	app := NewApp()
	// Create a temp file with spaces and special characters
	tempDir, err := os.MkdirTemp("", "video cut test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	specialPath := filepath.Join(tempDir, "Gravao de Tela.mp4")
	err = os.WriteFile(specialPath, []byte("dummy media"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Mock runCommand to verify it receives the path exactly
	app.runCommand = func(name string, arg ...string) ([]byte, error) {
		lastArg := arg[len(arg)-1]
		if lastArg != specialPath {
			t.Errorf("Expected path '%s', got '%s'", specialPath, lastArg)
		}
		return []byte("10.0\n"), nil
	}

	_, err = app.GetVideoDuration(specialPath)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestSplitVideo(t *testing.T) {
	app := NewApp()
	commands := [][]string{}
	app.runCommand = func(name string, arg ...string) ([]byte, error) {
		cmd := append([]string{name}, arg...)
		commands = append(commands, cmd)
		return []byte("done"), nil
	}

	err := app.SplitVideo("input.mp4", 100.0, 5.0, "output_dir", "p1", "p2")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(commands) != 2 {
		t.Errorf("Expected 2 commands, got %d", len(commands))
	}

	// Verify Part 1 (0 to 105)
	// args: -i input.mp4 -ss 0 -to 105 ... output_dir/p1.mp4
	foundP1 := false
	for _, cmd := range commands {
		fullCmd := strings.Join(cmd, " ")
		if strings.Contains(fullCmd, "p1.mp4") {
			foundP1 = true
			if !strings.Contains(fullCmd, "-ss 0") || !strings.Contains(fullCmd, "-to 105") {
				t.Errorf("Part 1 command arguments incorrect: %v", cmd)
			}
		}
	}
	if !foundP1 {
		t.Error("Part 1 command not found")
	}

	// Verify Part 2 (95 to end)
	foundP2 := false
	for _, cmd := range commands {
		fullCmd := strings.Join(cmd, " ")
		if strings.Contains(fullCmd, "p2.mp4") {
			foundP2 = true
			if !strings.Contains(fullCmd, "-ss 95") {
				t.Errorf("Part 2 command arguments incorrect: %v", cmd)
			}
		}
	}
	if !foundP2 {
		t.Error("Part 2 command not found")
	}
}

func TestSplitVideo_PermissionError(t *testing.T) {
	app := NewApp()
	// Mock command runner to simulate permission error
	app.runCommand = func(name string, arg ...string) ([]byte, error) {
		return nil, os.ErrPermission
	}

	err := app.SplitVideo("input.mp4", 50.0, 5.0, "restricted_dir", "p1", "p2")
	if err == nil {
		t.Error("Expected error for restricted directory, got nil")
	}
	if !strings.Contains(err.Error(), "part 1 failed") {
		t.Errorf("Expected 'part 1 failed' in error message, got: %v", err)
	}
}

func TestGetFFmpegVersion_Missing(t *testing.T) {
	app := NewApp()
	app.executableDir = "/nonexistent"
	app.runCommand = func(name string, arg ...string) ([]byte, error) {
		return nil, os.ErrNotExist
	}

	_, err := app.GetFFmpegVersion()
	if err == nil {
		t.Error("Expected error for missing FFmpeg, got nil")
	}
}
