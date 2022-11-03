/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   save_to_csv.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/23 18:23:51 by lpaulo-m          #+#    #+#             */
/*   Updated: 2022/11/02 23:46:26 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scraper

import (
	"encoding/csv"
	"os"
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
	file, err := os.Create(cont.filePath)
	checkErr(err)

	return file
}
