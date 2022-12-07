package xx

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	Name   string
	Files  []*File
	Dirs   []*Dir
	Parent *Dir
	Size   int
}

func NewDir(dirName string, parent *Dir) *Dir {
	return &Dir{
		Name:   dirName,
		Files:  []*File{},
		Dirs:   []*Dir{},
		Size:   0,
		Parent: parent,
	}
}

type File struct {
	Name string
	Size int
}

func NewFile(fileName string, size int) *File {
	return &File{
		Name: fileName,
		Size: size,
	}
}

func readFile(fileName string) []string {
	file, _ := os.Open(fileName)
	defer file.Close()

	result := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func processDir(console []string, dirName string, parent *Dir) *Dir {
	dir := NewDir(dirName, parent)

	for _, row := range console {
		parts := strings.Split(row, " ")
		if parts[0] == "dir" {
			dir.Dirs = append(dir.Dirs, NewDir(parts[1], dir))
		} else {
			fileSize, _ := strconv.Atoi(parts[0])
			dir.Files = append(dir.Files, NewFile(parts[1], fileSize))
			dir.Size += fileSize
		}

	}

	return dir
}

func processConsole(console []string, startDir *Dir) {
	currentDir := startDir
	for i := 0; i < len(console); i++ {
		cmd := console[i]
		if strings.HasPrefix(cmd, "$ cd") {
			parts := strings.Split(cmd, " ")
			dirName := parts[2]
			if dirName == ".." {
				currentDir = currentDir.Parent
			} else {
				for _, d := range currentDir.Dirs {
					if parts[2] == d.Name {
						currentDir = d
					}
				}
			}
		} else {
			i++
			buffer := []string{}
			for !strings.HasPrefix(console[i], "$ cd") {
				buffer = append(buffer, console[i])
			}
			currentDir = processDir(buffer, currentDir.Name, currentDir.Parent)
		}
	}
}

const (
	fileName = "input.txt"
)

func main() {
	console := readFile(fileName)
	startDir := NewDir("/", nil)
	processConsole(console[1:])
}
