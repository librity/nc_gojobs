/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   scrapper.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/22 23:24:59 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/23 18:27:34 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scrapper

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

type extractedJob struct {
	id       string
	link     string
	title    string
	location string
	salary   string
	summary  string
}

type scrapeControl struct {
	country string
	tech    string
	pages   int
	jobs    []extractedJob
}

var control = scrapeControl{}
var cont = &control

func initControl(country, tech string, pages int) {
	cont.country = country
	cont.tech = tech
	cont.pages = resolvePages(pages)
}

func Scrape(country, tech string, pages int) {
	initControl(country, tech, pages)

	extractJobs()
	saveToCSV()

	fmt.Println("Successfully scrapped", len(cont.jobs), "jobs")
}
