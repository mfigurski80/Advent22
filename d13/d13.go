package d13

import (
	"encoding/json"
	"fmt"

	. "github.com/mfigurski80/AOC22/utils"
)

type Packet = []interface{}

type Decision uint8

const (
	GOOD Decision = iota
	UNDECIDED
	BAD
)

func ComparePackets(a Packet, b Packet) Decision {
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
	i := 0
	sum_i := 0
	DoByNFileLines(3, "d13/in.txt", func(lines []string) {
		// fmt.Println(lines)
		a_body := Packet{}
		b_body := Packet{}
		err := json.Unmarshal([]byte(lines[0]), &a_body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal([]byte(lines[1]), &b_body)
		if err != nil {
			panic(err)
		}
		match := ComparePackets(a_body, b_body)
		fmt.Println(" == Example", i+1, "==", match == GOOD)
		if match == GOOD {
			sum_i += i + 1
		}
		i++
	})
	fmt.Println("Sum of good examples:", sum_i)
}
