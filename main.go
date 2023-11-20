package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// TODO: Take gzip file / files as input

	checkArgs()

	// ^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3}  (INFO|ERROR|WARN) \d+ \-\-\- \[.+\] .+\: (GET|POST|PUT|DELETE) .+ User: .+ Organisation: .+ Time: .+
	logLineRegexp := "^\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}\\.\\d{3}  (INFO|ERROR|WARN) \\d+ \\-\\-\\- \\[.+\\] .+\\: GET \\/txNewChecklistItemEntity.+ User: abc1\\.xyz Organisation: .+ Time: .+"

	re := regexp.MustCompile(logLineRegexp)

	logFilePath := os.Args[1]

	readFile, err := os.Open(logFilePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		logLine := fileScanner.Text()

		if re.MatchString(logLine) {
			fmt.Println(logLine)
		}
	}

	readFile.Close()
}

func checkArgs() {
	if len(os.Args) != 2 {
		log.Fatalf("expected exactly 1 argument but got %d. Usage: tiny-log-parser <log-file>", len(os.Args))
	}
}
