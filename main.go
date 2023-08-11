package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	profileTiles := ProfileTiles()

	var profileTmpl = "pageHTML.tmpl"
	tmpl, err := template.New(profileTmpl).ParseFiles(profileTmpl)
	if err != nil {
		panic(err)
	}

	// Print raw HTML file to stdout
	// err = tmpl.Execute(os.Stdout, profileTile)
	// if err != nil {
	// 	panic(err)
	// }

	var profileTilesHTML = "./profileTile.html"
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
