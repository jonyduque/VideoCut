package main

import (
	"os"
	"path/filepath"
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
