package util

import (
	"os"
	"bufio"
)

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}


func ReadWordlist(wordlistPath *string) []string{
	data, err := readLines(*wordlistPath)
	if err != nil{
		panic(err)
	}
	return data
}