/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   save_to_csv.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/23 18:23:51 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/23 18:27:31 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scrapper

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func saveToCSV() {
	cCSV := make(chan []string)

	for _, job := range cont.jobs {
		go jobToRow(cCSV, job)
	}

	w := initCSVWriter()
	writeHeader(w)

	for i := 0; i < len(cont.jobs); i++ {
		jobRow := <-cCSV
		wErr := w.Write(jobRow)
		checkErr(wErr)
	}
}

func jobToRow(cCSV chan<- []string, job extractedJob) {
	jobRow := []string{
		job.id,
		job.link,
		job.title,
		job.location,
		job.salary,
		job.summary,
	}

	cCSV <- jobRow
}

func writeHeader(w *csv.Writer) {
	headers := []string{"id", "link", "title", "location", "salary", "summary"}
	wErr := w.Write(headers)
	checkErr(wErr)
}

func initCSVWriter() *csv.Writer {
	file := initFile()

	w := csv.NewWriter(file)
	defer w.Flush()

	return w
}

func initFile() *os.File {
	os.MkdirAll(scrapesDir, os.ModePerm)
	filePath := makeFilePath()
	file, err := os.Create(filePath)
	checkErr(err)

	return file
}

func makeFilePath() string {
	nameFragments := []string{makeTimestamp(), cont.country, cont.tech, "jobs.csv"}
	fileName := strings.Join(nameFragments, "_")

	return filepath.Join(scrapesDir, fileName)
}

func makeTimestamp() string {
	now := time.Now()
	timestamp := now.Format(time.Stamp)
	timestamp = strings.ReplaceAll(timestamp, " ", "_")

	return timestamp
}
