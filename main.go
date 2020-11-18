package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	const versionFile = ".tool-versions"

	home, _ := os.UserHomeDir()
	var pluginsDir = home + "/.asdf/plugins/"

	fd, err := os.Open(versionFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	scanner := lineScanner(fd)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	fmt.Println("#!/bin/bash")
	for _, line := range text {
		parts := strings.Split(line, " ")
		pluginName := parts[0]
		pluginVersion := parts[1]

		config, err := os.Open(pluginsDir + pluginName + "/.git/config")
		if err != nil {
			log.Fatal(err)
		}
		scanner = lineScanner(config)
		var url string
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "url =") {
				parts = strings.Split(line, " = ")
				url = parts[1]
			}
		}
		if url != "" {
			fmt.Printf("asdf plugin-add %s %s\n", pluginName, url)
			fmt.Printf("asdf install %s %s\n", pluginName, pluginVersion)

		}
		_ = config.Close()
	}
}

func lineScanner(file *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}
