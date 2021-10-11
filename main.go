package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, v := range lines {
		scan(v)
	}
}

func scan(url string) bool {
	response, err := http.Get(url)

	if err != nil {
		_, netErrors := http.Get("https://www.google.com")

		if netErrors != nil {
			fmt.Fprintf(os.Stderr, "no internet\n")
			os.Exit(1)
		}

		return false
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == 200 {
		a := strings.Contains(string(body), "Registered Stream Socket Transports")
		fmt.Println(a)
	}
	return false
}
