package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFullWorkflow_Integration(t *testing.T) {
	app := NewApp()
	
	// Create a temporary workspace
	tempDir, err := os.MkdirTemp("", "videocut_integration")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Mock executable dir
	app.executableDir = tempDir
	
	// Create mock binaries
	os.WriteFile(filepath.Join(tempDir, "ffmpeg.exe"), []byte("dummy"), 0755)
	os.WriteFile(filepath.Join(tempDir, "ffprobe.exe"), []byte("dummy"), 0755)

	// Mock command runner to simulate success
	app.runCommand = func(name string, arg ...string) ([]byte, error) {
		// If ffprobe duration request
		for _, a := range arg {
			if a == "format=duration" {
				return []byte("100.0\n"), nil
			}
		}
		
		// If ffmpeg split request, create dummy output files
		for _, a := range arg {
			if filepath.Ext(a) == ".mp4" && a != "input.mp4" {
				os.WriteFile(a, []byte("output content"), 0644)
			}
		}
		
		return []byte("success"), nil
	}

	// 1. Test Duration Extraction
	duration, err := app.GetVideoDuration("input.mp4")
	if err != nil {
		t.Errorf("Duration extraction failed: %v", err)
	}
	if duration != 100.0 {
		t.Errorf("Expected 100.0, got %f", duration)
	}

	// 2. Test Video Splitting
	err = app.SplitVideo("input.mp4", 50.0, 5.0, tempDir, "part1", "part2")
	if err != nil {
		t.Errorf("Video splitting failed: %v", err)
	}

	// 3. Verify output files exist
	if _, err := os.Stat(filepath.Join(tempDir, "part1.mp4")); os.IsNotExist(err) {
		t.Error("Part 1 output file not found")
	}
	if _, err := os.Stat(filepath.Join(tempDir, "part2.mp4")); os.IsNotExist(err) {
		t.Error("Part 2 output file not found")
	}
}
