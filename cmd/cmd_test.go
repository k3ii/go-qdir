package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestGenerateRandomHex(t *testing.T) {
	length := 8
	hex, err := generateRandomHex(length)
	if err != nil {
		t.Fatalf("generateRandomHex failed: %v", err)
	}
	if len(hex) != length*2 { // each byte becomes two hex characters
		t.Errorf("Expected length %d, got %d", length*2, len(hex))
	}
}

func TestGetRandomPet(t *testing.T) {
	pet, err := getRandomPet()
	if err != nil {
		t.Fatalf("getRandomPet failed: %v", err)
	}
	found := false
	for _, p := range pets {
		if pet == p {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Generated pet %s not found in pets slice", pet)
	}
}

func TestCreateNestedTempDirs(t *testing.T) {
	tests := []struct {
		name      string
		tmpDir    bool
		depth     int
		useNames  bool
		usePets   bool
		hexLength int
	}{
		{"Default", false, 0, false, false, 16},
		{"Nested", false, 2, false, false, 16},
		{"TmpDir", true, 0, false, false, 16},
		{"UseNames", false, 1, true, false, 16},
		{"UsePets", false, 1, false, true, 16},
		{"CustomHexLength", false, 0, false, false, 32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := createNestedTempDirs(tt.tmpDir, tt.depth, tt.useNames, tt.usePets, tt.hexLength)
			if err != nil {
				t.Fatalf("createNestedTempDirs failed: %v", err)
			}

			// Check if the directory was created
			if _, err := os.Stat(path); os.IsNotExist(err) {
				t.Errorf("Directory was not created: %s", path)
			}

			// Check the depth
			depth := strings.Count(path, string(os.PathSeparator))
			var basePath string
			if tt.tmpDir {
				basePath = os.TempDir()
			} else {
				basePath = "."
			}
			baseDepth := strings.Count(basePath, string(os.PathSeparator))
			expectedDepth := baseDepth + tt.depth

			if depth != expectedDepth {
				t.Errorf("Expected depth %d, got %d", expectedDepth, depth)
			}

			// Check naming scheme
			dirs := strings.Split(path, string(os.PathSeparator))
			lastDir := dirs[len(dirs)-1]
			switch {
			case tt.useNames:
				if _, err := getRandomName(); err != nil {
					t.Errorf("Expected scientist name, got: %s", lastDir)
				}
			case tt.usePets:
				found := false
				for _, p := range pets {
					if lastDir == p {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected pet name, got: %s", lastDir)
				}
			default:
				if len(lastDir) != tt.hexLength {
					t.Errorf("Expected hex length %d, got %d", tt.hexLength, len(lastDir))
				}
			}

			// Clean up
			if err := os.RemoveAll(path); err != nil {
				t.Errorf("Failed to clean up test directory: %v", err)
			}
		})
	}
}

func TestTmp(t *testing.T) {
	// Mock cobra.Command
	cmd := &cobra.Command{}

	// Set some flags
	nestedDepth = 2
	useNames = true
	tmpDir = true

	err := tmp(cmd, []string{})
	if err != nil {
		t.Fatalf("tmp failed: %v", err)
	}

	// Check if a directory was created (we can't know the exact path)
	dirs, err := filepath.Glob(filepath.Join(os.TempDir(), "*", "*", "*"))
	if err != nil {
		t.Fatalf("Failed to glob directories: %v", err)
	}
	if len(dirs) == 0 {
		t.Errorf("No directories were created")
	}

	// Clean up
	for _, dir := range dirs {
		if err := os.RemoveAll(dir); err != nil {
			t.Errorf("Failed to clean up test directory: %v", err)
		}
	}
}
