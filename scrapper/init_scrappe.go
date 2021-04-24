/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   init_scrappe.go                                    :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/23 21:18:50 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/23 21:26:38 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scrapper

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type scrapeParams struct {
	country string
	tech    string
	pages   int
}

// Initializes the scrapper
func InitScrape(c echo.Context) scrapeParams {
	params := scrapeParams{}
	params.country = getCountry(c)
	params.tech = getTech(c)
	params.pages = getPages(c)
	logScrape(params)

	return params
}

func getCountry(c echo.Context) (country string) {
	country = c.FormValue("country")
	country = CleanField(country)
	return
}

func getTech(c echo.Context) (tech string) {
	tech = c.FormValue("tech")
	tech = CleanField(tech)
	return
}

func getPages(c echo.Context) int {
	raw := c.FormValue("pages")
	pages, err := strconv.Atoi(raw)

	if err == nil {
		return pages
	}
	return 0
}

func logScrape(params scrapeParams) {
	fmt.Println("=== Scrape request ===")
	fmt.Println("COUNTRY:", params.country)
	fmt.Println("TECHS:", params.tech)
	fmt.Println("PAGES:", params.pages)
}
