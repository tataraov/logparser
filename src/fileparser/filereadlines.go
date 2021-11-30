package fileparser

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadLinebyLine() {
	file, err := os.Open("/Users/tataraov/apps/code-frog-cli/input/artifactory.log")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
}
