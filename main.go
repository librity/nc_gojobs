/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/23 04:20:57 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/23 18:15:57 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "github.com/librity/nc_gojobs/scrapper"

var country = "it"
var tech = "ruby"
var pages = 10

func main() {
	scrapper.Scrape(country, tech, pages)
}
