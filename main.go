package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
	"github.com/pojntfx/gridge/pkg/stories"
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
	prefix := flag.String("prefix", "/gridge", "Prefix to build the app for")

	flag.Parse()

	// Define the handler
	h := &app.Handler{
		Title:           "gridge",
		Name:            "gridge",
		ShortName:       "gridge",
		Description:     "Sign, verify, encrypt and decrypt data with GPG.",
		LoadingLabel:    "Sign, verify, encrypt and decrypt data with GPG.",
		Author:          "Felicitas Pojtinger",
		ThemeColor:      "#151515",
		BackgroundColor: "#151515",
		Icon: app.Icon{
			Default: `/web/default.png`,
			Large:   `/web/large.png`,
		},
		Keywords: []string{
			"gpg",
			"pgp",
			"gopenpgp",
			"encryption",
			"signatures",
		},
		RawHeaders: []string{
			`<meta property="og:url" content="https://pojntfx.github.io/gridge/">`,
			`<meta property="og:title" content="gridge">`,
			`<meta property="og:description" content="Sign, verify, encrypt and decrypt data with GPG.">`,
			`<meta property="og:image" content="https://pojntfx.github.io/gridge/web/default.png">`,
		},
		Styles: []string{
			`https://unpkg.com/@patternfly/patternfly@4.135.2/patternfly.css`,
			`https://unpkg.com/@patternfly/patternfly@4.135.2/patternfly-addons.css`,
			`/web/index.css`,
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
