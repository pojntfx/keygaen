package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type DownloadOrViewModal struct {
	app.Compo

	SubjectA     bool
	SubjectANoun string
	SubjectAVerb string

	SubjectB     bool
	SubjectBNoun string
	SubjectBVerb string

	OnClose    func()
	OnDownload func()
	OnView     func()
}

func (c *DownloadOrViewModal) Render() app.UI {
	title := "File successfully "
	downloadLabel := "Download "
	viewLabel := "View "
	body := "You may now download or view "
	if c.SubjectA && c.SubjectB {
		title += c.SubjectBVerb + " and " + c.SubjectAVerb

		if c.SubjectANoun == "" {
			downloadLabel += c.SubjectBNoun
			viewLabel += c.SubjectBNoun
			body += "it"
		} else if c.SubjectBNoun == "" {
			downloadLabel += c.SubjectANoun
			viewLabel += c.SubjectANoun
			body += "it"
		} else {
			downloadLabel += c.SubjectBNoun + " and " + c.SubjectANoun
			viewLabel += c.SubjectBNoun + " and " + c.SubjectANoun
			body += "them"
		}
	} else if c.SubjectA {
		title += c.SubjectAVerb
		downloadLabel += c.SubjectANoun
		viewLabel += c.SubjectANoun
		body += "it"
	} else {
		title += c.SubjectBVerb
		downloadLabel += c.SubjectBNoun
		viewLabel += c.SubjectBNoun
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
