/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   scrapper.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/22 23:24:59 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/23 18:15:47 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const scrapesDir = "scrapes"

var jobsUrls = map[string]string{
	"it": "https://it.indeed.com/offerte-lavoro?limit=50&q=",
	"br": "https://br.indeed.com/empregos?limit=50&q=",
	"uk": "https://uk.indeed.com/jobs?limit=50&q=",
}

var viewJobUrls = map[string]string{
	"it": "https://it.indeed.com/viewjob?jk=",
	"br": "https://br.indeed.com/viewjob?jk=",
	"uk": "https://uk.indeed.com/viewjob?jk=",
}

var control = scrapeControl{}
var cont = &control

type scrapeControl struct {
	country string
	tech    string
	pages   int
	jobs    []extractedJob
}

type extractedJob struct {
	id       string
	link     string
	title    string
	location string
	salary   string
	summary  string
}

func Scrape(country, tech string, pages int) {
	initControl(country, tech, pages)

	extractJobs()
	saveJobsAsCSV()

	fmt.Println("Successfully scrapped", len(cont.jobs), "jobs")
}

func initControl(country, tech string, pages int) {
	cont.country = country
	cont.tech = tech
	cont.pages = resolvePages(pages)
}

func resolvePages(pages int) int {
	if pages > 0 {
		return pages
	}
	return getTotalPages()
}

func saveJobsAsCSV() {
	cCSV := make(chan []string)

	for _, job := range cont.jobs {
		go jobToRow(cCSV, job)
	}

	w := initCSVWriter()
	writeHeader(w)

	for i := 0; i < len(cont.jobs); i++ {
		jobRow := <-cCSV
		wErr := w.Write(jobRow)
		checkErr(wErr)
	}
}

func jobToRow(cCSV chan<- []string, job extractedJob) {
	jobRow := []string{
		job.id,
		job.link,
		job.title,
		job.location,
		job.salary,
		job.summary,
	}

	cCSV <- jobRow
}

func writeHeader(w *csv.Writer) {
	headers := []string{"id", "link", "title", "location", "salary", "summary"}
	wErr := w.Write(headers)
	checkErr(wErr)
}

func initCSVWriter() *csv.Writer {
	file := initFile()

	w := csv.NewWriter(file)
	defer w.Flush()

	return w
}

func initFile() *os.File {
	os.MkdirAll(scrapesDir, os.ModePerm)
	filePath := makeFilePath()
	file, err := os.Create(filePath)
	checkErr(err)

	return file
}

func makeFilePath() string {
	nameFragments := []string{makeTimestamp(), cont.country, cont.tech, "jobs.csv"}
	fileName := strings.Join(nameFragments, "_")

	return filepath.Join(scrapesDir, fileName)
}

func makeTimestamp() string {
	now := time.Now()
	timestamp := now.Format(time.Stamp)
	timestamp = strings.ReplaceAll(timestamp, " ", "_")

	return timestamp
}

func extractJobs() {
	var jobs []extractedJob
	cJobs := make(chan []extractedJob)

	for i := 0; i < cont.pages; i++ {
		go extractBatch(cJobs, i)
	}

	for i := 0; i < cont.pages; i++ {
		jobsBatch := <-cJobs
		jobs = append(jobs, jobsBatch...)
	}

	cont.jobs = jobs
}

func extractBatch(cJobs chan<- []extractedJob, page int) {
	var jobsBatch []extractedJob
	cJob := make(chan extractedJob)
	body := getBody(page)

	cards := body.Find(".jobsearch-SerpJobCard")
	cards.Each(func(i int, card *goquery.Selection) {
		go extractJob(cJob, card)
	})

	for i := 0; i < cards.Length(); i++ {
		job := <-cJob
		jobsBatch = append(jobsBatch, job)
	}

	cJobs <- jobsBatch
}

func extractJob(cJob chan<- extractedJob, card *goquery.Selection) {
	id, _ := card.Attr("data-jk")
	link := viewJobUrls[cont.country] + id
	title := cleanField(
		card.Find(".title>a").Text())
	location := cleanField(
		card.Find(".sjcl").Text())
	salary := cleanField(
		card.Find(".salaryText").Text())
	summary := cleanField(
		card.Find(".summary").Text())

	cJob <- extractedJob{
		id:       id,
		link:     link,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func buildJobsUrl(page int) string {
	baseUrl := jobsUrls[cont.country] + cont.tech
	jobsUrl := baseUrl + "&start=" + strconv.Itoa(page*50)

	return jobsUrl
}

func getBody(page int) *goquery.Document {
	jobsUrl := buildJobsUrl(page)

	fmt.Println("Requesting:", jobsUrl)
	res, err := http.Get(jobsUrl)
	checkErr(err)
	checkStatus(res)

	defer res.Body.Close()
	body, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	return body
}

func getTotalPages() int {
	pages := 0
	res, err := http.Get(buildJobsUrl(0))

	checkErr(err)
	checkStatus(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length() + 1
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatus(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with status", res.StatusCode)
	}
}

func cleanField(field string) string {
	trimmed := strings.TrimSpace(field)
	cleaned := strings.Fields(trimmed)
	joined := strings.Join(cleaned, " ")

	return joined
}
