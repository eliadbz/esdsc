package util

import (
	"os"
	"bufio"
)

func readLinesGenerator(path * string, file *os.File, linesChannel chan string){
	/*
	Reads lines into a channel
	*/
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		val := scanner.Text()
    	linesChannel <- val
	}
	close(linesChannel)
	file.Close()
}

func ReadWordlist(path *string) (chan string, error){
	file, err := os.Open(*path)
	linesChannel := make(chan string)
    if err != nil {
        return nil, err
    }

	go readLinesGenerator(path, file, linesChannel)
	return linesChannel, nil
}