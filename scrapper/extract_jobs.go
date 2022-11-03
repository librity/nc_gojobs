/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   extract_jobs.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/23 18:26:53 by lpaulo-m          #+#    #+#             */
/*   Updated: 2022/11/02 23:46:26 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scraper

import (
	"fmt"
	"net/http"

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
