package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// DownloadOrViewModal is a modal which provides the actions needed to download or view text
type DownloadOrViewModal struct {
	app.Compo

	SubjectA          bool   // Whether to display the first subject to download or view
	SubjectANoun      string // Noun form the first subject to download or view (i.e. "signature")
	SubjectAAdjective string // Adjective of the first subject to download or view (i.e. "signed")

	SubjectB          bool   // Whether to display the second subject to download or view
	SubjectBNoun      string // Noun form the second subject to download or view (i.e. "signature")
	SubjectBAdjective string // Adjective of the second subject to download or view (i.e. "signed")

	OnClose    func(used bool) // Handler to call when closing/cancelling the modal
	OnDownload func()          // Handler to call to download the subject(s)
	OnView     func()          // Handler to view to download the subject(s)

	ShowView bool // Whether to show the view action

	used bool
}

func (c *DownloadOrViewModal) Render() app.UI {
	title := "File successfully "
	downloadLabel := "Download "
	viewLabel := "View "
	body := "You may now download or view "
	if c.SubjectA && c.SubjectB {
		title += c.SubjectBAdjective + " and " + c.SubjectAAdjective

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
		title += c.SubjectAAdjective
		downloadLabel += c.SubjectANoun
		viewLabel += c.SubjectANoun
		body += "it"
	} else {
		title += c.SubjectBAdjective
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
					c.used = true

					c.OnDownload()
				}),
			app.If(
				c.ShowView,
				app.Button().
					Class("pf-c-button pf-m-link").
					Type("button").
					Text(viewLabel).
					OnClick(func(ctx app.Context, e app.Event) {
						c.used = true

						c.OnView()
					}),
			),
		},

		OnClose: func() {
			c.OnClose(c.used)
		},
	}
}
