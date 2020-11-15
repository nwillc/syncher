/*
 * Copyright (c) 2020, nwillc@gmail.com
 *
 * Permission to use, copy, modify, and/or distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

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
