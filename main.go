/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/22 23:24:59 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/22 23:51:53 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// https://it.indeed.com/offerte-lavoro?q=ruby&limit=50
// https://br.indeed.com/empregos?q=ruby&limit=50
// https://uk.indeed.com/jobs?q=ruby&limit=50

// https://CODE.indeed.com/jobs?q=LANGUAGE&limit=50

var baseUrl string = "https://it.indeed.com/offerte-lavoro?q=ruby&limit=50"

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)
}

func getPages() int {
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
