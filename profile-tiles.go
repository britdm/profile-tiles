package main

import (
	"fmt"
	"os"
	"text/template"
)

type Tile struct {
	Description string
	Image       string
	Links       string
	Location    string
	Name        string
	Resume      string
}

func main() {
	// Modify the data in profileTiles before generating the profile tile(s)
	profileTiles := []Tile{
		{
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse ultricies tempus elit.",
			Image:       "./resources/image.png",
			Links:       "https://linktr.ee/",
			Location:    "Austin, TX, US",
			Name:        "FName\nL/SName",
			Resume:      "./resources/resume.pdf",
		},
		// {
		// 	Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse ultricies tempus elit.",
		// 	Image:       "./resources/image.png",
		// 	Links:       "https://linktr.ee/",
		// 	Location:    "Austin, TX, US",
		// 	Name:        "FName\nL/SName",
		// 	Resume:      "./resources/resume.pdf",
		// },
		// {
		// 	Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse ultricies tempus elit.",
		// 	Image:       "./resources/image.png",
		// 	Links:       "https://linktr.ee/",
		// 	Location:    "Austin, TX, US",
		// 	Name:        "FName\nL/SName",
		// 	Resume:      "./resources/resume.pdf",
		// },
	}

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
