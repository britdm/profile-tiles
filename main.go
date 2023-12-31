package main

import (
	"fmt"
	"html/template"
	"os"
	"profile-tiles/pkgs"
	"strings"
)

func main() {
	countries := pkgs.LoadCountries()
	profileTiles := pkgs.ProfileTiles()
	region := ""

	for i, profile := range profileTiles {
		location_length := len(strings.Split(profile.Location, ","))
		country := fmt.Sprintf("|%s|", strings.TrimSpace(strings.Split(profile.Location, ",")[location_length-1]))
		mutableProfile := &profileTiles[i]
		for _, line := range strings.Split(countries, "\n") {
			if strings.Contains(line, country) {
				region = strings.Split(line, "|")[2]
				mutableProfile.AddRegion(region)
			}
		}
	}

	var profileTmpl = "pageHTML.tmpl"
	tmpl, err := template.New(profileTmpl).ParseFiles(profileTmpl)
	if err != nil {
		panic(err)
	}

	// Print raw HTML file to stdout
	// err = tmpl.Execute(os.Stdout, profileTiles)
	// if err != nil {
	// 	panic(err)
	// }

	var profileTilesHTML = "./index.html"
	writer, err := os.Create(profileTilesHTML)
	if err != nil {
		fmt.Println("Could not create file: ", err)
	}
	defer writer.Close()
	err = tmpl.Execute(writer, profileTiles)
	if err != nil {
		fmt.Println("Could write content to file: ", err)
	}
}
