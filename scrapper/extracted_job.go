/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   extracted_job.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/23 18:36:49 by lpaulo-m          #+#    #+#             */
/*   Updated: 2022/11/02 23:46:26 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scraper

import "github.com/PuerkitoBio/goquery"

type extractedJob struct {
	id       string
	link     string
	title    string
	location string
	salary   string
	summary  string
}

func extractJob(cJob chan<- extractedJob, card *goquery.Selection) {
	id := extractId(card)

	cJob <- extractedJob{
		id:       id,
		link:     makeLink(id),
		title:    extractTitle(card),
		location: extractLocation(card),
		salary:   extractSalary(card),
		summary:  extractSummary(card),
	}
}

func extractId(card *goquery.Selection) string {
	id, _ := card.Attr("data-jk")

	return id
}

func makeLink(id string) string {
	return viewJobUrls[cont.country] + id
}

func extractTitle(card *goquery.Selection) string {
	rawField := card.Find(".title>a").Text()

	return CleanField(rawField)
}

func extractLocation(card *goquery.Selection) string {
	rawField := card.Find(".sjcl").Text()

	return CleanField(rawField)
}

func extractSalary(card *goquery.Selection) string {
	rawField := card.Find(".salaryText").Text()

	return CleanField(rawField)
}

func extractSummary(card *goquery.Selection) string {
	rawField := card.Find(".summary").Text()

	return CleanField(rawField)
}
