package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
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
			akia, err := regexp.MatchString(`AKIA[A-Z0-9]{16}`, string(body))
			other, err2 := regexp.MatchString(`smtp\.sendgrid\.net|smtp\.mailgun\.org|smtp-relay\.sendinblue\.com|smtp.tipimail.com|smtp.sparkpostmail.com|vonage|nexmo|twilo|smtp.deliverabilitymanager.net|smtp.mailendo.com|mail.smtpeter.com|mail.smtp2go.com|smtp.socketlabs.com|secure.emailsrvr.com|mail.infomaniak.com|smtp.pepipost.com|smtp.elasticemail.com|smtp25.elasticemail.com|pro.turbo-smtp.com|smtp-pulse.com|in-v3.mailjet.com`, string(body))
			if akia {
				fmt.Println("[AKIA]: " + url)
				f, err := os.Create("hits-" + now.Format("01-02-2006") + "/" + "AKIA-" + url + ".txt")
				if err != nil {
					log.Fatal(err)
				}

				defer f.Close()

				_, err2 := f.WriteString(url)

				if err2 != nil {
					log.Fatal(err2)
				}
			}

			if err != nil {
				log.Fatal(err)
			}

			if other {
				fmt.Println("[OTHER]: " + url)
				f, err := os.Create("hits-" + now.Format("01-02-2006") + "/" + "OTHER-" + url + ".txt")
				if err != nil {
					log.Fatal(err)
				}

				defer f.Close()

				_, err2 := f.WriteString(url)

				if err2 != nil {
					log.Fatal(err2)
				}
			}

			if err2 != nil {
				log.Fatal(err2)
			}

			fmt.Println("[NOTHING]: " + url)
			f, err := os.Create("hits-" + now.Format("01-02-2006") + "/NOTHING-" + url + ".txt")

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			_, err3 := f.WriteString(url)

			if err3 != nil {
				log.Fatal(err3)
			}

		} else {
			fmt.Println(url)
		}
	}

	return false
}
