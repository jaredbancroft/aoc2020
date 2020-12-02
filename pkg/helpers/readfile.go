package helpers

import (
	"io/ioutil"
	"strconv"
	"strings"
)

//ReadIntFile will read a file of integers separated by a newline
//and return a slice of integers
func ReadIntFile(fname string) ([]int, error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	nums := make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

//ReadStringFile will read a file of strings separated by a newline
//and return a slice of strings
func ReadStringFile(fname string) ([]string, error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	return lines, nil
}
