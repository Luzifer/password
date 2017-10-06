package main

var flags = struct {
	CLI struct {
		Length            int
		SpecialCharacters bool
		JSON              bool
	}

	Server struct {
		Port int
	}
}{}
