/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   extract_jobs.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/23 18:26:53 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/23 18:27:01 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scrapper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

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

func cleanField(field string) string {
	trimmed := strings.TrimSpace(field)
	cleaned := strings.Fields(trimmed)
	joined := strings.Join(cleaned, " ")

	return joined
}
