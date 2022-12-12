package d7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	. "github.com/mfigurski80/AOC22/utils"
)

const LIMIT = 100000

type FileType uint

const (
	DIRECTORY FileType = iota
	FILE
)

type File struct {
	Name string
	Size uint
	Type FileType
}

var movePattern = regexp.MustCompile(`^\$ cd (?P<Match>\S+)$`)
var lsPattern = regexp.MustCompile(`^\$ ls$`)
var dirPattern = regexp.MustCompile(`^dir (?P<Match>\S+)$`)
var filePattern = regexp.MustCompile(`^(?P<Size>\d+) (?P<Match>\S+)$`)

func buildTreeFrom(fname string) (TreeNode[File], error) {
	root := TreeNode[File]{Children: nil, Parent: nil, Metadata: File{"/", 0, DIRECTORY}}
	currentDirectory := &root

	_, err := DoByFileLineWithError("d7/in.txt", func(line string) error {
		// CD
		if movePattern.MatchString(line) {
			match := movePattern.FindStringSubmatch(line)
			// fmt.Println("Doing CD", match[1])
			if match[1] == "/" {
				currentDirectory = &root
				return nil
			}
			if match[1] == ".." {
				// sum up sizes
				for _, child := range currentDirectory.Children {
					currentDirectory.Metadata.Size += child.Metadata.Size
				}
				currentDirectory = currentDirectory.Parent
				return nil
			}
			for _, child := range currentDirectory.Children {
				if child.Metadata.Name == match[1] {
					currentDirectory = child
					return nil
				}
			}
			return fmt.Errorf("no such directory")
		}
		// LS
		if lsPattern.MatchString(line) {
			// fmt.Println("Doing LS")
			return nil
		}
		// DIR
		if dirPattern.MatchString(line) {
			match := dirPattern.FindStringSubmatch(line)
			// fmt.Println("Doing DIR", match[1], 0)
			newDir := TreeNode[File]{Children: nil, Parent: currentDirectory, Metadata: File{match[1], 0, DIRECTORY}}
			currentDirectory.Children = append(currentDirectory.Children, &newDir)
			return nil
		}
		// FILE
		if filePattern.MatchString(line) {
			match := filePattern.FindStringSubmatch(line)
			// fmt.Println("Doing FILE", match[2], match[1])
			size, err := strconv.ParseUint(match[1], 10, 64)
			if err != nil {
				return err
			}
			newFile := TreeNode[File]{Children: nil, Parent: currentDirectory, Metadata: File{match[2], uint(size), FILE}}
			currentDirectory.Children = append(currentDirectory.Children, &newFile)
			return nil
		}
		return fmt.Errorf("no match on: %s", line)
	}, 0)
	if err != nil {
		return root, err
	}
	// sum up sizes until root
	for currentDirectory != nil {
		for _, child := range currentDirectory.Children {
			currentDirectory.Metadata.Size += child.Metadata.Size
		}
		currentDirectory = currentDirectory.Parent
	}
	return root, nil
}

func printTree(root TreeNode[File]) {
	root.DfsOnTree(0, func(node TreeNode[File], level int) {
		fmt.Printf("%s%s (%d)\n", strings.Repeat("  ", level), node.Metadata.Name, node.Metadata.Size)
	})
	// root.(root, 0, func(node Node, level int) {
	// 	fmt.Printf("%s%s (%d)\n", strings.Repeat("  ", level), node.Name, node.Size)
	// })
}

func Main() {
	root, err := buildTreeFrom("d7/in.txt")
	if err != nil {
		panic(err)
	}
	// print tree
	// printTree(root)
	// sum up all directories below LIMIT
	count := uint(0)
	root.DfsOnTree(0, func(node TreeNode[File], level int) {
		if node.Metadata.Type == DIRECTORY && node.Metadata.Size < LIMIT {
			count += node.Metadata.Size
			fmt.Printf("Found directory %s (%d)\n", node.Metadata.Name, node.Metadata.Size)
		}
	})
	fmt.Println("Total size:", count)
}
