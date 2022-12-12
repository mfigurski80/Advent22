package d7

import (
	"fmt"
	"regexp"
	"strconv"

	. "github.com/mfigurski80/AOC22/utils"
)

type Dir struct {
	Name string
	Size uint
}

func (d Dir) String() string {
	return fmt.Sprintf("%s (%d)", d.Name, d.Size)
}

var structure = regexp.MustCompile(`(?P<Sp>\W*)- (?P<Name>[\w\/\.]+) (?P<Tag>\(.*\))`)
var tag = regexp.MustCompile(`size=(?P<Size>\d+)`)

func parseLine(ln string) (uint, string, uint, error) {
	match := structure.FindStringSubmatch(ln)
	if len(match) <= 3 {
		return 0, "", 0, fmt.Errorf("no match found In: %v", ln)
	}
	// parse uint from string in match[4]
	match_size := tag.FindStringSubmatch(match[3])
	if len(match_size) <= 1 {
		return uint(len(match[1]) / 2), match[2], 0, nil
	}
	size, err := strconv.ParseUint(match_size[1], 10, 32)
	if err != nil {
		return 0, "", 0, err
	}
	return uint(len(match[1]) / 2), match[2], uint(size), nil
}

const LIMIT = 100000

func printFound(name string) {
	fmt.Printf("Found: %s\n", name)
}

func Main() {
	path := Stack[Dir]()
	lastIndent := uint(0)
	totalBelow := uint(0)

	// find size of all directories
	DoByFileLine("d7/in.txt", func(ln string) {
		indent, name, size, err := parseLine(ln)
		if err != nil {
			fmt.Println(err)
			return
		}
		// fmt.Printf("(%d) %-6s %d\n", indent, name, size)
		// init path, check file size...
		if path.Length() == 0 {
			path.Push(Dir{name, size})
			lastIndent = indent
			return
		}
		// fmt.Printf("%s\n", path.Array)
		// if size < LIMIT && size != 0 {
		// 	printFound(name)
		// }

		if lastIndent < indent {
			// if going deeper, add to path
			path.Push(Dir{name, size})
			lastIndent = indent
		} else if lastIndent == indent { // shallower
			// if staying same, resolve previous sibling and add to path
			sibling := path.Pop()
			parent := path.Peek()
			parent.Size += sibling.Size
			path.Push(Dir{name, size})
			lastIndent = indent
		} else {
			// if moving upwards, resolve sibling and nephew
			nephew := path.Pop()
			sibling := path.Pop()
			parent := path.Peek()
			parent.Size += sibling.Size + nephew.Size
			if sibling.Size+nephew.Size < LIMIT {
				printFound(sibling.Name)
				totalBelow += sibling.Size + nephew.Size
			}
			path.Push(Dir{name, size})
			lastIndent = indent
		}
	})

	// clean up remaining path
	for path.Length() > 1 {
		child := path.Pop()
		parent := path.Peek()
		parent.Size += child.Size
		if child.Size < LIMIT {
			printFound(child.Name)
			totalBelow += child.Size
		}
	}

	fmt.Printf("Total below %d: %d\n", LIMIT, totalBelow)
}
