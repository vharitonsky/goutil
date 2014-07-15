package goutil

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func ReadLines(file_path string) chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		file, err := os.Open(file_path)
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewReader(file)
		defer file.Close()
		for {
			line, err := reader.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			ch <- strings.TrimSuffix(string(line), "\n")
		}
	}()
	return ch
}

func SplitLine(line string) (string, string) {
	parts := strings.Split(line, "|")
	return parts[0], parts[1]
}
