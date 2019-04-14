package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Slots struct {
	Agenda string   `json:"agenda"`
	Date   string   `json:"date"`
	Times  []string `json:"times"`
}

func main() {
	resp, err := http.Get("https://app.bookitit.com/onlinebookings/datetime/?callback=jQuery21107171938286945687_1555098924310&type=default&publickey=21a8d76163e6f2dc0e5ca528c922d37c3&lang=es&services%5B%5D=bkt180488&agendas%5B%5D=bkt86361&version=1&src=https%3A%2F%2Fapp.bookitit.com%2Fes%2Fhosteds%2Fwidgetdefault%2F21a8d76163e6f2dc0e5ca528c922d37c3&srvsrc=https%3A%2F%2Fapp.bookitit.com&start=2019-04-01&end=2019-04-30&selectedPeople=1&_=1555098924314")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer resp.Body.Close()


	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %s", err.Error())
	}

	str := string(bytes)
	r, err := regexp.Compile("^callback=jQuery\\d+\\_\\d+\\(|\\|\\)|;")
	if err != nil {
		log.Fatalf("Unable to compile regex: %s", err.Error())
	}

	resultStr := r.ReplaceAllString(str, "")
	resultStr = strings.ReplaceAll(resultStr,")","")
	var list map[string][]*Slots
	err = json.Unmarshal([]byte(resultStr), &list)
	if err != nil {
		log.Fatalf("Unable to unmarshall JSON: %s: %s", err.Error(), resultStr)
	}

	for _, slot := range list["Slots"]{
		if len(slot.Times) > 0 {
			log.Printf("Slot found! Check the consulate site")
			os.Exit(0)
		}
	}
	log.Fatalf("No slots found :(")
}
