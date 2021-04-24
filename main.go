/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/23 04:20:57 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/23 21:52:09 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/librity/nc_gojobs/scrapper"
)

const url = "localhost:2000"
const cleanupScrapes = false

func main() {
	e := echo.New()

	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)

	e.Logger.Fatal(e.Start(url))
}

func handleHome(c echo.Context) error {
	return c.File("pages/home.html")
}

func handleScrape(c echo.Context) error {
	scrape := scrapper.InitScrape(c)
	scrapeResults, fileName := scrapper.Scrape(scrape)
	if cleanupScrapes {
		defer os.Remove(scrapeResults)
	}

	return c.Attachment(scrapeResults, fileName)
}
