/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   utils.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/23 18:37:08 by lpaulo-m          #+#    #+#             */
/*   Updated: 2022/11/02 23:46:26 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scraper

import (
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Cleans all the unecessary whitespace of a string
func CleanField(field string) string {
	trimmed := strings.TrimSpace(field)
	cleaned := strings.Fields(trimmed)
	joined := strings.Join(cleaned, " ")

	return joined
}

func buildJobsUrl(page int) string {
	baseUrl := jobsUrls[cont.country] + cont.tech
	jobsUrl := baseUrl + "&start=" + strconv.Itoa(page*50)

	return jobsUrl
}

func makeFileName() string {
	nameFragments := []string{makeTimestamp(), cont.country, cont.tech, "jobs.csv"}
	fileName := strings.Join(nameFragments, "_")

	return fileName
}

func makeFilePath(fileName string) string {
	return filepath.Join(scrapesDir, fileName)
}

func makeTimestamp() string {
	now := time.Now()
	timestamp := now.Format(time.Stamp)
	timestamp = strings.ReplaceAll(timestamp, " ", "_")

	return timestamp
}
