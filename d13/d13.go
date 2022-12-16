package d13

import (
	"encoding/json"
	"fmt"
	"sort"

	. "github.com/mfigurski80/AOC22/utils"
)

type Packet = []interface{}

type Decision int8

const (
	BAD       Decision = -1
	UNDECIDED Decision = 0
	GOOD      Decision = 1
)

func ComparePackets(a, b Packet) Decision {
	// fmt.Println("Comparing", a, b)
	for i, _ := range a {
		if i >= len(b) { // if right runs out of items first, good order
			return BAD
		}
		// fmt.Println("  Comparing", a[i], b[i])
		dec := UNDECIDED
		switch a[i].(type) {
		case float64:
			switch b[i].(type) {
			case float64: // if both values are integers...
				if a[i].(float64) > b[i].(float64) { // if left is higher than right, bad order
					return BAD
				} else if a[i].(float64) < b[i].(float64) { // if left is lower than right, good order
					return GOOD
				}
			case []interface{}: // if mismatch, convert to list and compare
				dec = ComparePackets(Packet{a[i].(float64)}, b[i].(Packet))
			default:
				panic(fmt.Sprintf("Unknown B type: %T, %v", b[i], b[i]))
			}
		case []interface{}:
			switch b[i].(type) {
			case float64:
				dec = ComparePackets(a[i].(Packet), Packet{b[i].(float64)})
			case []interface{}: // if both values are lists...
				dec = ComparePackets(a[i].(Packet), b[i].(Packet)) // compare internal values
			default:
				panic(fmt.Sprintf("Unknown B type: %T, %v", b[i], b[i]))
			}
		default:
			panic(fmt.Sprintf("Unknown A type: %T, %v", b[i], b[i]))
		}
		if dec == UNDECIDED {
			continue
		}
		return dec
	}
	if len(a) < len(b) { // if left runs out of items first, good order
		return GOOD
	}
	return UNDECIDED
}

func Main() {
	packets := make([]Packet, 2)
	packets[0] = Packet{Packet{float64(2)}}
	packets[1] = Packet{Packet{float64(6)}}
	DoByFileLine("d13/in.txt", func(line string) {
		if len(line) == 0 {
			return
		}
		body := Packet{}
		err := json.Unmarshal([]byte(line), &body)
		if err != nil {
			panic(err)
		}
		packets = append(packets, body)
	})

	// sort
	sort.SliceStable(packets, func(i, j int) bool {
		return ComparePackets(packets[i], packets[j]) == GOOD
	})

	// find [[2]] and [[6]] index
	st, en := 0, 0
	for i, p := range packets {
		s := fmt.Sprintf("%v", p)
		fmt.Printf("%-2d %s\n", i+1, s)
		if s == "[[2]]" {
			st = i + 1
		}
		if s == "[[6]]" {
			en = i + 1
		}
	}
	fmt.Println("Start:", st, ", End:", en)
	fmt.Println("Answer:", en*st)
}
