package cli

var flags = struct {
	CLI struct {
		Length            int
		SpecialCharacters bool
		JSON              bool
		Num               int

		XKCD        bool
		PrependDate bool
		Separator   string
	}

	Server struct {
		Port int
	}
}{}
