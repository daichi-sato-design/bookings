package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/daichi-sato-design/bookings/pkg/config"
	"github.com/daichi-sato-design/bookings/pkg/handlers"
	"github.com/daichi-sato-design/bookings/pkg/render"
)

const (
	portNumber = ":4000"
)

var app config.AppConfig
var session *scs.SessionManager

// main はメインアプリケーション機能
func main(){
	// 本番環境ではこれをtrueに変更
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session


	tc, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println("Starting application on port ", portNumber)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}