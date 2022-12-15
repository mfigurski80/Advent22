package d11

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/mfigurski80/AOC22/utils"
)

type Operation rune

const (
	ADD = Operation('+')
	SUB = Operation('-')
	MUL = Operation('*')
	DIV = Operation('/')
)

type Monkey = struct {
	Id      uint
	Items   []int
	Operate func(int) int
	Test    func(int) bool
	IfTrue  uint
	IfFalse uint
	Count   uint
}

func DoForMonkeyInFile(fname string, fn func(Monkey)) {
	DoForRegexMatches(fname, `(?m:Monkey (\d+):\W+Starting items:\W?((?:\d+(?:\,\W)?)+))\W+Operation: new = old ([\-\+\*\/]) (\d+|old)\W+Test: divisible by (\d+)\W+If true: throw to monkey (\d+)\W+If false: throw to monkey (\d+)`, func(m []string) {
		id, err := strconv.Atoi(m[1])
		if err != nil {
			panic(err)
		}
		items := []int{}
		for _, item := range strings.Split(m[2], ", ") {
			i, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			items = append(items, i)
		}
		operation := Operation(rune(m[3][0]))
		var operate func(int) int = nil
		operand, err := strconv.Atoi(m[4])
		if err != nil {
			switch operation {
			case ADD:
				operate = func(a int) int {
					return a + a
				}
			case SUB:
				operate = func(a int) int {
					return 0
				}
			case MUL:
				operate = func(a int) int {
					return a * a
				}
			case DIV:
				operate = func(a int) int {
					return 1
				}
			default:
				panic("Unknown operation")
			}
		} else {
			switch operation {
			case ADD:
				operate = func(a int) int {
					return a + operand
				}
			case SUB:
				operate = func(a int) int {
					return a - operand
				}
			case MUL:
				operate = func(a int) int {
					return a * operand
				}
			case DIV:
				operate = func(a int) int {
					return a / operand
				}
			default:
				panic("Unknown operation")
			}
		}

		divisibleBy, err := strconv.Atoi(m[5])
		if err != nil {
			panic(err)
		}
		test := func(a int) bool {
			return a%divisibleBy == 0
		}
		ifTrue, err := strconv.Atoi(m[6])
		if err != nil {
			panic(err)
		}
		ifFalse, err := strconv.Atoi(m[7])
		if err != nil {
			panic(err)
		}
		fn(Monkey{
			Id:      uint(id),
			Items:   items,
			Operate: operate,
			Test:    test,
			IfTrue:  uint(ifTrue),
			IfFalse: uint(ifFalse),
			Count:   0,
		})
	})
}

func Main() {
	monkeys := []Monkey{}
	expectMonkeyId := 0
	DoForMonkeyInFile("d11/temp.txt", func(m Monkey) {
		monkeys = append(monkeys, m)
		if int(m.Id) != expectMonkeyId {
			panic("Monkey IDs are not sequential")
		}
		expectMonkeyId++
	})

	for i := 0; i < 10000; i++ {
		// fmt.Println("ROUND", i+1)
		for id, monkey := range monkeys {
			for _, item := range monkey.Items {
				monkeys[id].Count++
				update := monkey.Operate(item) / 3
				// to := monkey.IfFalse
				if monkey.Test(update) {
					monkeys[monkey.IfTrue].Items = append(monkeys[monkey.IfTrue].Items, update)
					// to = monkey.IfTrue
				} else {
					monkeys[monkey.IfFalse].Items = append(monkeys[monkey.IfFalse].Items, update)
				}
				// fmt.Printf("Monkey %d [%d]: (%d -> %d) to Monkey %d\n", monkey.Id, monkey.Count, item, update, to)
			}
			monkeys[id].Items = []int{}
		}
	}

	// find top 2 most-used monkeys
	mostUsed := []Monkey{monkeys[0], monkeys[1]}
	if mostUsed[0].Count < mostUsed[1].Count {
		mostUsed = []Monkey{monkeys[1], monkeys[0]}
	}
	for _, monkey := range monkeys[2:] {
		if monkey.Count > mostUsed[0].Count {
			mostUsed[0], mostUsed[1] = monkey, mostUsed[0]
		} else if monkey.Count > mostUsed[1].Count {
			mostUsed[1] = monkey
		}
	}

	fmt.Println("Most used monkeys:", mostUsed[0].Id, mostUsed[1].Id)
	fmt.Printf("Monkey Business: %d * %d => %d\n", mostUsed[0].Count, mostUsed[1].Count, mostUsed[0].Count*mostUsed[1].Count)

}
