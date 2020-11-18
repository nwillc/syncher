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
	"bytes"
	"flag"
	"fmt"
	"github/nwillc/syncher/gen/version"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type pair struct {
	a, b string
}

func main() {
	flag.Parse()

	if *flags.version {
		fmt.Println("version", version.Version)
		os.Exit(0)
	}

	fd, err := os.Open(*flags.input)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	pluginUrls, err := pluginsUrls()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("#!/bin/bash")
	for _, pair := range pairs(fd) {
		pluginName := pair.a
		pluginVersion := pair.b

		url := pluginUrls[pluginName]
		if url != "" {
			fmt.Println("asdf plugin-add", pluginName, url)
			fmt.Println("asdf install", pluginName, pluginVersion)
		}
	}
}

func pluginsUrls() (map[string]string, error) {
	asdf, err := exec.LookPath("asdf")
	if err != nil {
		return nil, err
	}

	pluginList := exec.Command(asdf, "plugin", "list", "--urls")

	output, err := pluginList.Output()
	if err != nil {
		return nil, err
	}

	plugins := make(map[string]string)
	for _, pair := range pairs(bytes.NewReader(output)) {
		plugins[pair.a] = pair.b
	}
	return plugins, nil
}

func pairs(reader io.Reader) []pair {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var pairs []pair
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		pairs = append(pairs, pair{a: parts[0], b: parts[1]})
	}
	return pairs
}
