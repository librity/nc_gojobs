/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/23 04:20:57 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/12/12 00:01:52 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/librity/nc_gojobs/scrapper"
)

var port = os.Getenv("PORT")

const cleanupScrapes = false

func main() {
	if port == "" {
		port = "2000"
	}

	e := echo.New()

	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)

	e.Logger.Fatal(e.Start(":" + port))
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
