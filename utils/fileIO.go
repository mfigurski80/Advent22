package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func DoByFileLine(f string, fn func(string)) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fn(scanner.Text())
	}
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
	fn(lines)
}

type StopIterationError struct {
	Message string
}

func (e StopIterationError) Error() string {
	return fmt.Sprintf("Iteration stopped: %s", e.Message)
}

func DoByFileLineWithError(f string, fn func(string) error, seek int64) (int64, error) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Seek(seek, 0)
	scanner := bufio.NewScanner(file)
	pos := seek
	for scanner.Scan() {
		txt := scanner.Text()
		pos += int64(len(txt)) + 1
		err := fn(txt)
		if err != nil {
			return pos, err
		}
	}
	return pos, nil
}

func DoForRegexMatches(f string, re string, fn func([]string)) {
	// read file contents
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	matcher := regexp.MustCompile(re)
	matches := matcher.FindAllStringSubmatch(string(body), -1)
	for _, m := range matches {
		fn(m)
	}
}
