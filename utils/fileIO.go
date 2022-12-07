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

func DoByNFileLines(N uint, f string, fn func([]string)) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines = make([]string, N)
	scanner := bufio.NewScanner(file)
	i := uint(0)
	for scanner.Scan() {
		lines[i%N] = scanner.Text()
		if i%N == N-1 {
			fn(lines)
		}
		i++
	}
}
