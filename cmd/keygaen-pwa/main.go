package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/components"
	"github.com/pojntfx/keygaen/pkg/stories"
)

func main() {
	// Client-side code
	{
		app.Route("/", &components.Home{})
		app.Route("/stories", &stories.Index{})
		app.RunWhenOnBrowser()
	}

	// Server-/build-side code

	// Parse the flags
	serve := flag.Bool("serve", false, "Serve the app instead of building it")
	laddr := flag.String("laddr", "0.0.0.0:22255", "Address to listen on when serving the app")
	dist := flag.String("dist", "out/web", "Directory to build the app to")
	prefix := flag.String("prefix", "/keygaen", "Prefix to build the app for")

	flag.Parse()

	// Define the handler
	h := &app.Handler{
		Title:           "keygaen",
		Name:            "keygaen",
		ShortName:       "keygaen",
		Description:     "Sign, verify, encrypt and decrypt data with PGP.",
		LoadingLabel:    "Sign, verify, encrypt and decrypt data with PGP.",
		Author:          "Felicitas Pojtinger",
		ThemeColor:      "#151515",
		BackgroundColor: "#151515",
		Icon: app.Icon{
			Default: `/web/default.png`,
			Large:   `/web/large.png`,
		},
		Keywords: []string{
			"pgp",
			"pgp",
			"gopenpgp",
			"encryption",
			"signatures",
		},
		RawHeaders: []string{
			`<meta property="og:url" content="https://pojntfx.github.io/keygaen/">`,
			`<meta property="og:title" content="keygaen">`,
			`<meta property="og:description" content="Sign, verify, encrypt and decrypt data with PGP.">`,
			`<meta property="og:image" content="https://pojntfx.github.io/keygaen/web/default.png">`,
		},
		Styles: []string{
			"/web/main.css",
		},
	}

	// Serve if specified
	if *serve {
		log.Println("Listening on", *laddr)
		log.Println("Stories can be found on /stories")

		if err := http.ListenAndServe(*laddr, h); err != nil {
			log.Fatal("could not serve:", err)
		}

		return
	}

	// Build if not specified
	h.Resources = app.GitHubPages(*prefix)

	if err := app.GenerateStaticWebsite(*dist, h); err != nil {
		log.Fatal("could not build static website:", err)
	}
}
