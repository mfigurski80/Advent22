package utils

import (
	"bufio"
	"os"
)

func DoByLine(f *os.File, fn func(string)) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fn(scanner.Text())
	}
}

func DoByFileLine(f string, fn func(string)) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	DoByLine(file, fn)
}
