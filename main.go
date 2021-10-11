package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	err := os.Mkdir("hits-"+now.Format("01-02-2006"), 0755)
	fmt.Println(" ▄▄▄·  ▄▄▄· ▄▄▄·  ▄▄·  ▄ .▄▄▄▄ .    ·▄▄▄▄•      ▪   ")
	fmt.Println("▐█ ▀█ ▐█ ▄█▐█ ▀█ ▐█ ▌▪██▪▐█▀▄.▀·    ▪▀·.█▌▪     ██  ")
	fmt.Println("▄█▀▀█  ██▀·▄█▀▀█ ██ ▄▄██▀▐█▐▀▀▪▄    ▄█▀▀▀• ▄█▀▄ ▐█· ")
	fmt.Println("▐█ ▪▐▌▐█▪·•▐█ ▪▐▌▐███▌██▌▐▀▐█▄▄▌    █▌▪▄█▀▐█▌.▐▌▐█▌ ")
	fmt.Println(" ▀  ▀ .▀    ▀  ▀ ·▀▀▀ ▀▀▀ · ▀▀▀     ·▀▀▀ • ▀█▄▀▪▀▀▀ ")
	fmt.Print("[NDDs List]: ")
	var filename string
	fmt.Scanln(&filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	length := len(lines)
	var wg sync.WaitGroup
	wg.Add(length)
	for v := 0; v < length; v++ {
		go func(v int) {
			defer wg.Done()
			scan(lines[v])
		}(v)
	}

	wg.Wait()
}

func scan(url string) bool {
	response, err := http.Get("http://" + url + "/phpinfo.php")

	if err != nil {
		_, netErrors := http.Get("https://www.google.com")

		if netErrors != nil {
			return false
		}

		return false
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return false
	}

	now := time.Now()

	if response.StatusCode == 200 {
		x := strings.Contains(string(body), "Registered Stream Socket Transports")
		if x {
			fmt.Println("[Valid]: " + url)
			f, err := os.Create("hits-" + now.Format("01-02-2006") + "/" + url + ".txt")

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			_, err2 := f.WriteString(url)

			if err2 != nil {
				log.Fatal(err2)
			}
		} else {
			fmt.Println(url)
		}
	}
	return false
}
