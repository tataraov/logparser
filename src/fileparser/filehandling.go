package fileparser

import (
	"fmt"
	"io/ioutil"
)

func TestParser() {
	data, err := ioutil.ReadFile("/Users/tataraov/apps/code-frog-cli/input/artifactory.log")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}
