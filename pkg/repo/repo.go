package repo

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	config "github.com/rwxrob/config/pkg"
)

const (
	DefaultFileName = "README.md"
)

// exists returns an error if the path specified exists
func exists(path string) (bool, error) {
	rootDir := getRootDir()
	check := fmt.Sprintf("%v/%v", rootDir, path)
	if _, err := os.Stat(check); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	// It exists, throw nothing
	return true, nil
}

// Init initializes a new Zettelkasten at the configured directory root.
func Init() error {
	rootDir := getRootDir()
	exist, err := exists(rootDir)
	if err != nil {
		return err
	} else if exist {
		return fmt.Errorf("directory %v already exists", rootDir)
	}

	if err := os.MkdirAll(rootDir, 0750); err != nil {
		return err
	}

	// Initialize Git repo
	buf := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	cmd := exec.Command("git", "init")
	cmd.Stdout = buf
	cmd.Stderr = stderr
	cmd.Dir = rootDir

	if err := cmd.Run(); err != nil {
		log.Print(stderr)
		return err
	}

	return nil
}

func getRootDir() string {
	rootDir := config.Query("zet", ".root")

	// TODO: Find a better way to handle null config
	if rootDir == `null` {
		// TODO: Handle this error
		dir, _ := os.UserHomeDir()
		rootDir = fmt.Sprintf("%v/.zet", dir)
	}

	return rootDir
}

func getZetPath(id string) string {
	rootDir := getRootDir()
	return fmt.Sprintf("%v/%v", rootDir, id)
}

func Exists(id string) (bool, error) {
	filename := getZetPath(id)
	fmt.Println(filename)
	if _, err := os.Stat(filename); err != nil {
		// If there was an error, check if it's a file doesn't exist error
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func New(id string) error {
	filename := getZetPath(id)
	exists, err := exists(filename)
	if err != nil {
		return err
	} else if exists {
		return fmt.Errorf("directory %v already exists", filename)
	}

	if err := os.Mkdir(filename, 0750); err != nil {
		return err
	}

	// Create new entry with editor
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi"
	}

	filepath := fmt.Sprintf("%v/%v", filename, DefaultFileName)

	cmd := exec.Command(editor, filepath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = filename

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
