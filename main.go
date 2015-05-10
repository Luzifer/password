package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Luzifer/password/lib"
	"github.com/codegangsta/cli"
	"github.com/gorilla/mux"
)

var pwd *securepassword.SecurePassword

func init() {
	pwd = securepassword.NewSecurePassword()
}

func main() {
	app := cli.NewApp()
	app.Usage = "generates secure random passwords"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:   "serve",
			Usage:  "start an API server to request passwords",
			Action: startAPIServer,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "port",
					Value: 3000,
					Usage: "port to listen on",
				},
			},
		},
		{
			Name:   "get",
			Usage:  "generate and return a secure random password",
			Action: printPassword,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "length, l",
					Value: 20,
					Usage: "length of the generated password",
				},
				cli.BoolFlag{
					Name:  "special, s",
					Usage: "use special characters in your password",
				},
			},
		},
	}

	app.Run(os.Args)
}

func startAPIServer(c *cli.Context) {
	r := mux.NewRouter()
	r.HandleFunc("/", handleFrontend).Methods("GET")
	r.PathPrefix("/assets").HandlerFunc(handleAssets).Methods("GET")
	r.HandleFunc("/v1/getPassword", handleAPIGetPasswordv1).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), nil)
}

func printPassword(c *cli.Context) {
	password, err := pwd.GeneratePassword(c.Int("length"), c.Bool("special"))
	if err != nil {
		switch {
		case err == securepassword.ErrLengthTooLow:
			fmt.Println("The password has to be more than 4 characters long to meet the security considerations")
		default:
			fmt.Println("An unknown error occured")
		}
		os.Exit(1)
	}

	fmt.Println(password)
}

func handleAPIGetPasswordv1(res http.ResponseWriter, r *http.Request) {
	length, err := strconv.Atoi(r.URL.Query().Get("length"))
	if err != nil {
		length = 20
	}
	special := r.URL.Query().Get("special") == "true"

	if length > 128 || length < 4 {
		http.Error(res, "Please do not use length with more than 128 or fewer than 4 characters!", http.StatusNotAcceptable)
		return
	}

	password, err := pwd.GeneratePassword(length, special)

	res.Header().Add("Content-Type", "text/plain")
	res.Header().Add("Cache-Control", "no-cache")
	res.Write([]byte(password))
}

func handleFrontend(res http.ResponseWriter, r *http.Request) {
	res.Header().Add("Content-Type", "text/html")
	buf, err := Asset("frontend/index.html")
	if err != nil {
		http.Error(res, "Unable to load interface", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	res.Write(buf)
}

func handleAssets(res http.ResponseWriter, r *http.Request) {
	buf, err := Asset(fmt.Sprintf("frontend%s", r.URL.Path))
	if err != nil {
		http.Error(res, "Unable to load interface", http.StatusInternalServerError)
		return
	}
	res.Write(buf)
}
