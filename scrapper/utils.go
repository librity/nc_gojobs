/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   utils.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/23 18:37:08 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/23 18:44:19 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package scrapper

import (
	"strconv"
	"strings"
)

func cleanField(field string) string {
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