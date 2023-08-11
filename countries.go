package main

import (
	"fmt"
	"io"
	"net/http"
)

func LoadCountries() string {
	resp, err := http.Get("https://gist.githubusercontent.com/richjenks/15b75f1960bc3321e295/raw/e9b473faed0c7512d6720d71d485b662cd743d25/countries.md")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return string(body)
}

func (t *Tile) addRegion(region string) {
	old_location := t.Location
	t.Location = fmt.Sprintf("%s - %s", old_location, region)
}
