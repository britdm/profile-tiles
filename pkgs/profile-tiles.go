package pkgs

type Tile struct {
	Description string
	Image       string
	Links       string
	Location    string
	Name        string
	Resume      string
}

func ProfileTiles() []Tile {
	// Modify the data in profileTiles before generating the profile tile(s)
	profileTiles := []Tile{
		{
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse ultricies tempus elit.",
			Image:       "./resources/image.png",
			Links:       "https://linktr.ee/",
			Location:    "Austin, TX, US",
			Name:        "F_Name L/S_Name",
			Resume:      "./resources/resume.pdf",
		},
		{
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse ultricies tempus elit.",
			Image:       "./resources/image.png",
			Links:       "https://linktr.ee/",
			Location:    "London, England, GB",
			Name:        "F_Name L/S_Name",
			Resume:      "./resources/resume.pdf",
		},
		{
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse ultricies tempus elit.",
			Image:       "./resources/image.png",
			Links:       "https://linktr.ee/",
			Location:    "Canberra ACT 2601, AU",
			Name:        "F_Name L/S_Name",
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

	return profileTiles
}
