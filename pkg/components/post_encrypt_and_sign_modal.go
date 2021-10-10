package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type PostEncryptAndSignModal struct {
	app.Compo

	Signed    bool
	Encrypted bool

	OnClose    func()
	OnDownload func()
	OnView     func()
}

func (c *PostEncryptAndSignModal) Render() app.UI {
	title := "File successfully "
	downloadLabel := "Download "
	viewLabel := "View "
	body := "You may now download or view "
	if c.Signed && c.Encrypted {
		title += "encrypted and signed"
		downloadLabel += "cypher and signature"
		viewLabel += "cypher and signature"
		body += "them"
	} else if c.Signed {
		title += "signed"
		downloadLabel += "signature"
		viewLabel += "signature"
		body += "it"
	} else {
		title += "encrypted"
		downloadLabel += "cypher"
		viewLabel += "cypher"
		body += "it"
	}
	title += "!"
	body += "."

	return &Modal{
		ID:    "post-encrypt-and-sign-modal",
		Icon:  "fas fa-check",
		Title: title,
		Class: "pf-m-success",
		Body: []app.UI{
			app.Text(body),
		},
		Footer: []app.UI{
			app.Button().
				Aria("disabled", "false").
				Class("pf-c-button pf-m-primary").
				Type("button").
				Text(downloadLabel).
				OnClick(func(ctx app.Context, e app.Event) {
					c.OnDownload()
				}),
			app.Button().
				Class("pf-c-button pf-m-link").
				Type("button").
				Text(viewLabel).
				OnClick(func(ctx app.Context, e app.Event) {
					c.OnView()
				}),
		},

		OnClose: func() {
			c.OnClose()
		},
	}
}
