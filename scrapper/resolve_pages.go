/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   resolve_pages.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/23 18:25:44 by lpaulo-m          #+#    #+#             */
/*   Updated: 2022/11/02 23:46:26 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scraper

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func resolvePages(pages int) int {
	if pages > 0 {
		return pages
	}
	return getTotalPages()
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
