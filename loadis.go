package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please provide a file name to load\n")
		os.Exit(1)
	}

	filename := os.Args[1]
	fmt.Printf("Loading file %q\n", filename)

	dirname, err := ioutil.TempDir(".", "tmpledis")
	if err != nil {
		fmt.Printf("Failed to create temp dir: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Data will be loaded into %q\n", dirname)

	ledisInfo, err := getLedisInfo(dirname)
	if err != nil {
		fmt.Printf("Failed to connect to ledis: %v\n", err)
		os.Exit(1)
	}

	if _, err := ledisInfo.Conn.LoadDumpFile(filename); err != nil {
		fmt.Printf("Failed to load file: %v\n", err)
		os.Exit(1)
	}

	for {
		cmd, done := nextCommand()
		if done {
			break
		}

		err := cmd.Execute(ledisInfo)
		if err != nil {
			fmt.Printf("Command failed: %v\n", err)
		}
	}

	fmt.Printf("k, bye\n")
}

func nextCommand() (*Command, bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if strings.ToLower(text) == "quit" {
		return nil, true
	}

	return &Command{Text: text}, false
}
