/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/22 23:24:59 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/23 00:30:14 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// https://it.indeed.com/offerte-lavoro?q=ruby&limit=50
// https://br.indeed.com/empregos?q=ruby&limit=50
// https://uk.indeed.com/jobs?q=ruby&limit=50

// https://CODE.indeed.com/jobs?q=LANGUAGE&limit=50

var baseUrl string = "https://it.indeed.com/offerte-lavoro?q=ruby&limit=50"

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

func main() {
	var jobs []extractedJob
	totalPages := getTotalPages()

	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}

	fmt.Println(jobs)
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageUrl := baseUrl + "&start=" + strconv.Itoa(page*50)

	fmt.Println("Requesting:", pageUrl)
	res, err := http.Get(pageUrl)
	checkErr(err)
	checkStatus(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	cards := doc.Find(".jobsearch-SerpJobCard")
	cards.Each(func(i int, card *goquery.Selection) {
		job := ectractJob(card)
		jobs = append(jobs, job)
	})

	return jobs
}

func ectractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-jk")
	title := cleanField(
		card.Find(".title>a").Text())
	location := cleanField(
		card.Find(".sjcl").Text())
	salary := cleanField(
		card.Find(".salaryText").Text())
	summary := cleanField(
		card.Find(".summary").Text())

	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func getTotalPages() int {
	pages := 0
	res, err := http.Get(baseUrl)

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
