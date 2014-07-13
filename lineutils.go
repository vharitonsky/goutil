package lineutils

import(
    "os"
    "bufio"
    "strings"
    "log"
)

func readLines(file_path string) chan string {
	ch := make(chan string)
	go func() {
		file, err := os.Open(file_path)
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewReader(file)
		defer file.Close()
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				break
			}
			ch <- strings.TrimSuffix(string(line), "\n")
		}
		close(ch)
	}()
	return ch
}

func splitLine(line string) (string, string) {
	parts := strings.Split(line, "|")
	return parts[0], parts[1]
}
