package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type Book struct {
	Code     string
	Chapters int
}

type Collection struct {
	Code  string
	Books []Book
}

type ScriptureJSON struct {
	Book    string `json:"book"`
	Chapter int    `json:"chapter"`
	Verse   string `json:"verse"`
	Text    string `json:"text"`
}

func ScrapeScripture() error {
	scriptures := []Collection{
		{
			Code: "bofm",
			Books: []Book{
				{"1-ne", 22},
				{"2-ne", 33},
				{"jacob", 7},
				{"enos", 1},
				{"jarom", 1},
				{"omni", 1},
				{"w-of-m", 1},
				{"mosiah", 29},
				{"alma", 63},
				{"hel", 16},
				{"3-ne", 30},
				{"4-ne", 1},
				{"morm", 9},
				{"ether", 15},
				{"moro", 10},
			},
		},
		{
			Code: "dc-testament",
			Books: []Book{
				{"dc", 138},
			},
		},
		{
			Code: "pgp",
			Books: []Book{
				{"moses", 8},
				{"abr", 5},
				{"js-m", 1},
				{"js-h", 1},
			},
		},
		{
			Code: "ot",
			Books: []Book{
				{"gen", 50},
				{"ex", 40},
				{"lev", 27},
				{"num", 36},
				{"deut", 34},
				{"josh", 24},
				{"judg", 21},
				{"ruth", 4},
				{"1-sam", 31},
				{"2-sam", 24},
				{"1-kgs", 22},
				{"2-kgs", 25},
				{"1-chr", 29},
				{"2-chr", 36},
				{"ezra", 10},
				{"neh", 13},
				{"esth", 10},
				{"job", 42},
				{"ps", 150},
				{"prov", 31},
				{"eccl", 12},
				{"song", 8},
				{"isa", 66},
				{"jer", 52},
				{"lam", 5},
				{"ezek", 48},
				{"dan", 12},
				{"hosea", 14},
				{"joel", 3},
				{"amos", 9},
				{"obad", 1},
				{"jonah", 4},
				{"micah", 7},
				{"nahum", 3},
				{"hab", 3},
				{"zeph", 3},
				{"hag", 2},
				{"zech", 14},
				{"mal", 4},
			},
		},
		{
			Code: "nt",
			Books: []Book{
				{"matt", 28},
				{"mark", 16},
				{"luke", 24},
				{"john", 21},
				{"acts", 28},
				{"rom", 16},
				{"1-cor", 16},
				{"2-cor", 13},
				{"gal", 6},
				{"eph", 6},
				{"philip", 4},
				{"col", 4},
				{"1-thes", 5},
				{"2-thes", 3},
				{"1-tim", 6},
				{"2-tim", 4},
				{"titus", 3},
				{"philem", 1},
				{"heb", 13},
				{"james", 5},
				{"1-pet", 5},
				{"2-pet", 3},
				{"1-jn", 5},
				{"2-jn", 1},
				{"3-jn", 1},
				{"jude", 1},
				{"rev", 22},
			},
		},
	}

	c := colly.NewCollector(
		colly.AllowedDomains("www.churchofjesuschrist.org"),
	)

	if err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*churchofjesuschrist.org*",
		Parallelism: 1,
		Delay:       300 * time.Millisecond,
	}); err != nil {
		return fmt.Errorf("setting rate limit: %w", err)
	}

	scriptureJSONs := make([]ScriptureJSON, 0, 45_000)

	c.OnHTML("div.body-block", func(e *colly.HTMLElement) {
		parts := strings.Split(e.Request.URL.Path, "/")

		if len(parts) < 6 {
			fmt.Printf("Unexpected URL path: %s\n", e.Request.URL.Path)
			return
		}

		book := parts[4]
		chapterStr := parts[5]

		chapter, err := strconv.Atoi(chapterStr)
		if err != nil {
			fmt.Printf("Invalid chapter %q in URL %s: %v\n", chapterStr, e.Request.URL.String(), err)
			return
		}

		e.ForEach("p.verse", func(_ int, verse *colly.HTMLElement) {
			verseNumber := strings.TrimSpace(verse.ChildText("span.verse-number"))

			text := strings.TrimSpace(verse.Text)
			text = strings.TrimPrefix(text, verseNumber)
			text = strings.TrimSpace(text)

			if verseNumber == "" || text == "" {
				return
			}

			scriptureJSONs = append(scriptureJSONs, ScriptureJSON{
				Book:    book,
				Chapter: chapter,
				Verse:   verseNumber,
				Text:    text,
			})
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error visiting %s: status=%d err=%v\n", r.Request.URL.String(), r.StatusCode, err)
	})

	for _, collection := range scriptures {
		for _, book := range collection.Books {
			for chapter := 1; chapter <= book.Chapters; chapter++ {
				url := fmt.Sprintf(
					"https://www.churchofjesuschrist.org/study/scriptures/%s/%s/%d?lang=eng",
					collection.Code,
					book.Code,
					chapter,
				)

				if err := c.Visit(url); err != nil {
					fmt.Printf("Error scheduling URL %s: %v\n", url, err)
				}
			}
		}
	}

	file, err := os.Create("scriptures.json")
	if err != nil {
		return fmt.Errorf("creating scriptures.json: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(scriptureJSONs); err != nil {
		return fmt.Errorf("writing JSON: %w", err)
	}

	return nil
}
