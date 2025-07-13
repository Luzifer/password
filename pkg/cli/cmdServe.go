package cli

import (
	"embed"
	"fmt"
	"log"
	"mime"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	http_helper "github.com/Luzifer/go_helpers/v2/http"
	pwd "github.com/Luzifer/password/lib/v2"
)

const defaultHTTPListenPort = 3000

//go:embed frontend/**
var frontend embed.FS

func getCmdServe() *cobra.Command {
	cmd := cobra.Command{
		Use:   "serve",
		Short: "start an API server to request passwords",
		RunE:  actionCmdServe,
	}

	cmd.Flags().IntVar(&flags.Server.Port, "port", defaultHTTPListenPort, "port to listen on")

	return &cmd
}

func actionCmdServe(_ *cobra.Command, _ []string) error {
	r := mux.NewRouter()
	r.HandleFunc("/", handleFrontend).Methods("GET")
	r.PathPrefix("/assets").HandlerFunc(http_helper.GzipFunc(handleAssets)).Methods("GET")
	r.HandleFunc("/v1/getPassword", handleAPIGetPasswordv1).Methods("GET")
	r.HandleFunc("/v1/healthz", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", flags.Server.Port),
		Handler:           r,
		ReadHeaderTimeout: time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("listening for HTTP on port %d: %w", flags.Server.Port, err)
	}

	return nil
}

func handleAPIGetPasswordv1(res http.ResponseWriter, r *http.Request) {
	length, err := strconv.Atoi(r.URL.Query().Get("length"))
	if err != nil {
		length = 20
	}

	special := r.URL.Query().Get("special") == "true"
	xkcd := r.URL.Query().Get("xkcd") == "true"
	prependDate := r.URL.Query().Get("date") != "false"
	xkcdSeparator := r.URL.Query().Get("separator")

	if length > 128 || length < 4 {
		http.Error(res, "Please do not use length with more than 128 or fewer than 4 characters!", http.StatusNotAcceptable)
		return
	}

	pwd.DefaultXKCD.Separator = xkcdSeparator

	var password string
	if xkcd {
		password, err = pwd.DefaultXKCD.GeneratePassword(length, prependDate)
	} else {
		password, err = pwd.NewSecurePassword().GeneratePassword(length, special)
	}

	if err != nil {
		http.Error(res, errors.Wrap(err, "generating password").Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Add("Content-Type", "text/plain")
	res.Header().Add("Cache-Control", "no-cache")
	if _, err = res.Write([]byte(password)); err != nil {
		logrus.WithError(err).Error("writing password")
	}
}

func handleFrontend(res http.ResponseWriter, _ *http.Request) {
	res.Header().Add("Content-Type", "text/html")
	buf, err := frontend.ReadFile("frontend/index.html")
	if err != nil {
		http.Error(res, "Unable to load interface", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	if _, err = res.Write(buf); err != nil {
		logrus.WithError(err).Error("writing frontend buffer")
	}
}

func handleAssets(res http.ResponseWriter, r *http.Request) {
	buf, err := frontend.ReadFile(fmt.Sprintf("frontend%s", r.URL.Path))
	if err != nil {
		http.Error(res, "Unable to load interface", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", mime.TypeByExtension(path.Ext(r.URL.Path)))
	if _, err = res.Write(buf); err != nil {
		logrus.WithError(err).Error("writing assets buffer")
	}
}
