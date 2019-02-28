package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"path"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	http_helper "github.com/Luzifer/go_helpers/http"
	pwd "github.com/Luzifer/password/lib"
)

func getCmdServe() *cobra.Command {
	cmd := cobra.Command{
		Use:   "serve",
		Short: "start an API server to request passwords",
		Run:   actionCmdServe,
	}

	cmd.Flags().IntVar(&flags.Server.Port, "port", 3000, "port to listen on")

	return &cmd
}

func actionCmdServe(cmd *cobra.Command, args []string) {
	r := mux.NewRouter()
	r.HandleFunc("/", handleFrontend).Methods("GET")
	r.PathPrefix("/assets").HandlerFunc(http_helper.GzipFunc(handleAssets)).Methods("GET")
	r.HandleFunc("/v1/getPassword", handleAPIGetPasswordv1).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(fmt.Sprintf(":%d", flags.Server.Port), nil)
}

func handleAPIGetPasswordv1(res http.ResponseWriter, r *http.Request) {
	length, err := strconv.Atoi(r.URL.Query().Get("length"))
	if err != nil {
		length = 20
	}
	special := r.URL.Query().Get("special") == "true"
	xkcd := r.URL.Query().Get("xkcd") == "true"
	prependDate := r.URL.Query().Get("date") != "false"

	if length > 128 || length < 4 {
		http.Error(res, "Please do not use length with more than 128 or fewer than 4 characters!", http.StatusNotAcceptable)
		return
	}

	var password string
	if xkcd {
		password, err = pwd.DefaultXKCD.GeneratePassword(length, prependDate)
	} else {
		password, err = pwd.NewSecurePassword().GeneratePassword(length, special)
	}

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

	res.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(r.URL.Path)))
	res.Write(buf)
}
