package d5

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/mfigurski80/AOC22/utils"
)

// ERROR STRUCT
type EndOfInputSectionError struct {
	Section string
}

func (e EndOfInputSectionError) Error() string {
	return fmt.Sprintf("End of %s input section", e.Section)
}

// CARGO STACKS STRUCT
type CargoStacks []StackStruct[rune]

func (cs CargoStacks) New(size uint) CargoStacks {
	stacks := make(CargoStacks, 0, size)
	i := uint(0)
	for i < size {
		stacks = append(stacks, Stack[rune]())
		i++
	}
	return stacks
}

func (cs CargoStacks) String() string {
	res := ""
	for i, stack := range cs {
		res += fmt.Sprintf("%d: [", i)
		for _, v := range *stack.Array {
			res += fmt.Sprintf("%c", v)
		}
		res += "]\n"
	}
	return res
}

func parseOperation(s string) (int, int, int, error) {
	tokens := strings.Split(s, " ")
	if len(tokens) != 6 || tokens[0] != "move" || tokens[2] != "from" || tokens[4] != "to" {
		return 0, 0, 0, fmt.Errorf("invalid operation: %s", s)
	}
	times, err := strconv.Atoi(tokens[1])
	if err != nil {
		return 0, 0, 0, err
	}
	from, err := strconv.Atoi(tokens[3])
	if err != nil {
		return 0, 0, 0, err
	}
	to, err := strconv.Atoi(tokens[5])
	if err != nil {
		return 0, 0, 0, err
	}
	return times, from, to, nil
}

// MAIN FUNCTION
func Main() {
	// stack structure
	var stacks CargoStacks
	n_elements := uint(0)

	// read stack structure
	endOfA, err := DoByFileLineWithError("d5/in.txt", func(s string) error {
		fmt.Println(s)
		if len(s) == 0 {
			return &EndOfInputSectionError{"Stack"}
		}
		// init if needed
		if n_elements == 0 || stacks == nil {
			n_elements = uint((len(s) + 1) / 4)
			stacks = stacks.New(n_elements)
		}
		// add elements to stacks
		for i := uint(0); i < n_elements; i++ {
			ch := rune(s[i*4+1])
			if ch == ' ' || s[i*4] != '[' || s[i*4+2] != ']' { // if not a box
				continue
			}
			stacks[i].Push(ch)
		}
		return nil
	}, 0)
	if err != nil {
		if _, ok := err.(*EndOfInputSectionError); !ok {
			panic(err)
		}
	}

	// flip stacks
	for _, stack := range stacks {
		ReverseInPlace(*stack.Array)
	}

	// read operations input (from same file)
	_, err = DoByFileLineWithError("d5/in.txt", func(s string) error {
		times, from, to, err := parseOperation(s)
		if err != nil {
			return err
		}
		// pick up #times boxes from from stack
		from_s := stacks[from-1]
		to_s := stacks[to-1]
		target_from_length := from_s.Length() - times
		if target_from_length < 0 {
			return fmt.Errorf("not enough boxes in stack %d to take %d", from, times)
		}
		// fmt.Printf("moving %c from %d to %d\n", (*from_s.Array)[target_from_length:], from, to)
		*to_s.Array = append(*to_s.Array, (*from_s.Array)[target_from_length:]...)
		*from_s.Array = (*from_s.Array)[:target_from_length]

		return nil
	}, endOfA)
	if err != nil {
		fmt.Println(stacks)
		panic(err)
	}

	fmt.Println(stacks)

	res := ""
	for _, stack := range stacks {
		res += fmt.Sprintf("%c", stack.Pop())
	}
	fmt.Println(res)
}
