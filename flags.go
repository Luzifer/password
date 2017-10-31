package main

var flags = struct {
	CLI struct {
		Length            int
		SpecialCharacters bool
		JSON              bool

		XKCD        bool
		PrependDate bool
	}

	Server struct {
		Port int
	}
}{}
