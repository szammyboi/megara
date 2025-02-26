package main

import (
	"context"
	"fmt"
	"time"
	"strconv"
	"github.com/valyala/fasthttp"
	"os"
	"sort"
	"strings"
	"github.com/gocolly/colly/v2"
	"net/url"
)

// App struct
type App struct {
	ctx context.Context
	client *fasthttp.Client
	color_schemes []ColorScheme
	active ColorScheme
}

type SearchResult struct {
	Language string `json:"language"`
	Word string `json:"word"`
	Strength int `json:"strength"`
}

type WordDetails struct {
	FrWord string `json:"fr_word"`
	ToWord string `json:"to_word"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}


// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	readTimeout, _ := time.ParseDuration("500ms")
	writeTimeout, _ := time.ParseDuration("500ms")
	maxIdleConnDuration, _ := time.ParseDuration("1h")
	a.client = &fasthttp.Client{
		ReadTimeout:                   readTimeout,
		WriteTimeout:                  writeTimeout,
		MaxIdleConnDuration:           maxIdleConnDuration,
		NoDefaultUserAgentHeader:      true,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096, DNSCacheDuration: time.Hour,
		}).Dial,
	}
	a.LoadColors()
}

// domReady is called after front-end resources have been loaded
func (a *App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) SearchWord(lang1 string, lang2 string, word string) []SearchResult {
	results := []SearchResult{}
	req := fasthttp.AcquireRequest()
		req.SetRequestURI("https://www.wordreference.com/autocomplete?dict=" +lang1 + lang2 +"&query=" + url.QueryEscape(word))
	req.Header.SetMethod(fasthttp.MethodGet)
	resp := fasthttp.AcquireResponse()
	err := a.client.Do(req, resp)
	fasthttp.ReleaseRequest(req)
	if err == nil && resp.StatusCode() == 200 {
		current := [4][]byte{}
		current_index := 0
		for _, char := range resp.Body() {
			if char == '\t' {
				current_index++
			} else if char == '\n' {
				current_index = 0
				var search SearchResult
				search.Word = string(current[0])
				search.Language = string(current[1])
				num, _ := strconv.Atoi(string(current[2]))
				search.Strength =num
				results = append(results, search)
				current = [4][]byte{}
			} else {
				current[current_index] = append(current[current_index], char)
			}
		}
		sort.Slice(results, func(i, j int) bool {
			return results[i].Strength > results[j].Strength // && results[i].Language > results[j].Language
		})

		if (len(results) > 10) {
			return results[0:10]
		}
		return results
	} else {
		fmt.Fprintf(os.Stderr, "ERR Connection error: %v\n", err)
	}
	fasthttp.ReleaseResponse(resp)
	return results
}

func (a *App) FetchDetails(lang1 string, lang2 string, word string) []WordDetails {
	c := colly.NewCollector()
	var details []WordDetails

	c.OnHTML(".WRD tbody tr[id]", func(e *colly.HTMLElement) {
		fr := e.ChildText(".FrWrd strong")
		
		char_count := 0
		char_count += len(e.ChildText(".ToWrd i"))
		char_count += len(e.ChildText(".ToWrd em"))
		to_total := e.ChildText(".ToWrd")
		to := to_total[0:len(to_total)-char_count]
		to = strings.ReplaceAll(to, "⇒", "")
		fr = strings.ReplaceAll(fr, "⇒", "")
		to = strings.TrimSpace(to)
		fr = strings.TrimSpace(fr)

		fr_count := 0
		fr_count += len(e.ChildText(".FrWrd strong span"))
		fr = fr[0:len(fr)-fr_count]
		new_word := WordDetails{fr, to}
		details = append(details, new_word)
		fmt.Println(fr, to)
	})

	c.Visit("https://www.wordreference.com/" +lang1 + lang2 +"/" + word)
	return details
}

func (a *App) GetColors() []ColorScheme {
	return a.color_schemes
}

func (a *App) GetActiveColor() ColorScheme {
	return a.active
}
