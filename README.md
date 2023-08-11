# profile-tiles

This is a Go project created to generate an HTML file with profile "tiles" that can be a standalone page or integrated into a larger HTML file.

The project can be modified to use any template file format to generate HTML using the the same Tile sturct values in `profile-tiles.go`. [See this article](https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go) for an example of how to build and use templates (.tmpl) in Go.

## Tools Used

 - Random Latin word generator [Loreum Ipsum](https://www.lipsum.com/).
 - Locations add regions based [countries.md](https://gist.githubusercontent.com/richjenks/15b75f1960bc3321e295/raw/e9b473faed0c7512d6720d71d485b662cd743d25/countries.md) raw Markdown from [Rick Jenks](https://gist.github.com/richjenks/15b75f1960bc3321e295).
 - Icons from the [iconify material design icon library](https://icon-sets.iconify.design/mdi/).

### Profile Suggestions

Description for each profile can look like [this](https://sg.indeed.com/career-advice/interviewing/describe-yourself-in-one-sentence). Long descriptions will enable a scroll bar on the tile.

## How to use

Populate the `profileTiles` defined in the `profile-tiles.go` file with profile data, the returned variable will be a list of `Tiles`. Uncomment the additional tiles to see how spacing will look with multiple rows of tiles.

For a single profile, replace the dummy files `./resources/image.png` and `./resources/resume.pdf`, then update the data in `profile-tiles.go` before generating the HTML.

Run the script `go run main.go pprofile-tiles.go`. This should generate an HTML file, which can be used as a single page or integrated into a larger HTML file. To preview in VSCode using the HTML Preview plugin use `CTRL+SHIFT+V`.

### Build

Generate the executable `go build .`