package main

var flags = struct {
	CLI struct {
		Length            int
		SpecialCharacters bool
	}

	Server struct {
		Port int
	}
}{}
