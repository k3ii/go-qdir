package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	nestedDepth int
	useNames    bool
	usePets     bool
	tmpDir      bool
	hexLength   int
)

var pets = []string{
	"dog", "cat", "mouse", "cow", "rabbit", "hamster", "parrot", "goldfish",
	"turtle", "guinea pig", "horse", "donkey", "goat", "sheep", "pig", "chicken",
	"duck", "goose", "ferret", "gerbil", "canary", "parakeet", "iguana", "gecko",
}

func generateRandomHex(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("Failed to generate random bytes: %v", err)
	}
	return hex.EncodeToString(bytes), nil
}

func getRandomPet() (string, error) {
	max := big.NewInt(int64(len(pets)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", fmt.Errorf("failed to generate random number: %v", err)
	}
	return pets[n.Int64()], nil
}

func createNestedTempDirs(tmpDir bool, depth int, useNames bool, usePets bool, hexLength int) (string, error) {
	var path string
	if tmpDir {
		path = os.TempDir()
	} else {
		path = "."
	}
	for i := 0; i <= depth; i++ {
		var name string
		var err error
		switch {
		case useNames:
			name, err = getRandomName()
		case usePets:
			name, err = getRandomPet()
		default:
			name, err = generateRandomHex(hexLength / 2) // divide by 2 because each byte becomes 2 hex characters
		}
		if err != nil {
			return "", err
		}
		path = filepath.Join(path, name)
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		return "", fmt.Errorf("failed to create nested directories: %w", err)
	}

	return path, nil
}

func tmp(cmd *cobra.Command, args []string) error {
	dir, err := createNestedTempDirs(tmpDir, nestedDepth, useNames, usePets, hexLength)
	if err != nil {
		return err
	}

	fmt.Printf("Created directory: %s\n", dir)

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "qdir",
	Short: "Quick Directory Generator",
	Long: `Qdir (Quick Directory) is a versatile tool for creating directories with customizable naming schemes.

Key features:
  - Generate directories with random hexadecimal names, names of notable scientists/technologists, or pet types
  - Create nested directory structures with controllable depth
  - Option to use system's temporary directory or current working directory
  - Adjustable length for hexadecimal names

Qdir streamlines the process of creating uniquely named directories for various purposes such as 
temporary workspaces, project organization, or test environments.`,
	RunE: tmp,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&nestedDepth, "nested", "n", 0, "Depth of nested directories to create")
	rootCmd.Flags().BoolVarP(&useNames, "use-names", "u", false, "Use scientist/technologist names instead of random hex")
	rootCmd.Flags().BoolVarP(&usePets, "use-pets", "p", false, "Use pet types instead of random hex")
	rootCmd.Flags().BoolVarP(&tmpDir, "tmp", "t", false, "Use the system's temporary directory")
	rootCmd.Flags().IntVarP(&hexLength, "hex-length", "l", 16, "Length of the random hexadecimal name")
}

