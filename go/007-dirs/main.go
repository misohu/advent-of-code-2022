package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func processConsole(console []string, currentDir *Dir) *Dir {
	start := new(Dir)
	currentDir = NewDir("/", nil)
	start = currentDir
	for _, line := range console[1:] {
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == ".." {
					currentDir = currentDir.Parent
				} else {
					for _, d := range currentDir.Dirs {
						if d.Name == parts[2] {
							currentDir = d
						}
					}
				}
			}
			continue
		} else {
			if parts[0] == "dir" {
				currentDir.Dirs = append(currentDir.Dirs, NewDir(parts[1], currentDir))
			} else {
				fileSize, _ := strconv.Atoi(parts[0])
				currentDir.Files = append(currentDir.Files, NewFile(parts[1], fileSize))
			}
		}
	}
	return start
}

func findFiles(dir *Dir) []*Dir {
	size := 0
	res := []*Dir{}
	for _, f := range dir.Files {
		size += f.Size
	}

	if len(dir.Dirs) == 0 {
		dir.Size = size
		return []*Dir{dir}
	}

	for _, d := range dir.Dirs {
		res = append(res, findFiles(d)...)
		size += d.Size
	}
	dir.Size = size
	res = append(res, dir)

	return res
}

const (
	fileName = "input.txt"
	maxSize  = 100000
)

func main() {
	console := readFile(fileName)
	start := processConsole(console, nil)
	fmt.Println(start)
	allDirs := findFiles(start)
	usedSize := 0
	for _, d := range allDirs {
		if d.Name == "/" {
			usedSize = d.Size
		}
	}
	fmt.Println(usedSize)
	total := 70000000
	need := 30000000
	unused := total - usedSize
	missing := need - unused
	sort.Slice(allDirs, func(i, j int) bool {
		return allDirs[i].Size < allDirs[j].Size
	})

	for _, d := range allDirs {
		fmt.Printf("%q -> %d\n", d.Name, d.Size)
	}

	for _, x := range allDirs {
		if x.Size > missing {
			fmt.Println(x.Size)
			break
		}
	}
	// fmt.Println(res)
}
