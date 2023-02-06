package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())  // where we get the info line by line
	}
}

// this will be to read the config file to find where our notes are stored and any other configs we need to save
func ReadConfig(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		keyVal := strings.Split(scanner.Text(), "=")
		fmt.Println(keyVal[0])
		fmt.Println(keyVal[1])
	}
}

func main() {

	// need to open a file and save it in the program in some way
	ReadFile("testNote.txt")
	

	// should probably use a struct and edit it using functions
}