/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   scrape.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/22 23:24:59 by lpaulo-m          #+#    #+#             */
/*   Updated: 2022/11/02 23:46:26 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scraper

import (
	"fmt"
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

type scrapeControl struct {
	country  string
	tech     string
	pages    int
	fileName string
	filePath string
	jobs     []extractedJob
}

var control = scrapeControl{}
var cont = &control

// Scrape Indeed.com listings by term, country and number of pages
func Scrape(params scrapeParams) (string, string) {
	initControl(params)
	extractJobs()
	saveToCSV()

	fmt.Println("Successfully scrapped", len(cont.jobs), "jobs")
	return cont.filePath, cont.filePath
}

func initControl(params scrapeParams) {
	cont.country = params.country
	cont.tech = params.tech
	cont.pages = resolvePages(params.pages)
	cont.fileName = makeFileName()
	cont.filePath = makeFilePath(cont.fileName)
}
