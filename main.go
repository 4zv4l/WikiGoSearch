package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var (
		url     = "https://fr.wikipedia.org/wiki/"
		subject = ""
	)
	println("WikiGoSearch :")
	scanner := bufio.NewScanner(os.Stdin)
	print("> ")
	scanner.Scan()
	subject = scanner.Text()
	wiki, err := GetWiki(url+subject, subject)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(wiki)
}

// GetWiki gets the latest blog title headings from the url
// given and returns them as a list.
func GetWiki(url string, subject string) (string, error) {

	// Get the HTML
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// save the page and split after seeing a point.
	wiki := strings.SplitAfter(doc.Find("p").Text(), ".")
	wiki = strings.Split(wiki[0], subject)
	return subject + wiki[1], nil
}
