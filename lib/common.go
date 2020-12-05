package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ConvertStringsToInts(lines []string) ([]int, error) {
	var results []int
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to int: %v", line, err)
		}

		results = append(results, n)
	}

	return results, nil
}

func ReadFile(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)

	var results []string

	for ; scanner.Scan(); {
		results = append(results, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("scanner error: %v", err)
	}

	return results, nil
}