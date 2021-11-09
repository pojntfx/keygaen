package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygean/pkg/components"
)

type DownloadOrViewModalStory struct {
	Story

	modalOpen bool

	selectedMode string

	subjectA     bool
	subjectANoun string
	subjectAVerb string

	subjectB     bool
	subjectBNoun string
	subjectBVerb string
}

func (c *DownloadOrViewModalStory) Render() app.UI {
	c.EnableShallowReflection()

	if c.selectedMode == "" {
		c.selectedMode = "signed_encrypted"
	}

	switch c.selectedMode {
	case "signed":
		c.subjectA = true
		c.subjectANoun = "signature"
		c.subjectAVerb = "signed"

		c.subjectB = false
		c.subjectBNoun = "cypher"
		c.subjectBVerb = "encrypted"
	case "encrypted":
		c.subjectA = false
		c.subjectANoun = "signature"
		c.subjectAVerb = "signed"

		c.subjectB = true
		c.subjectBNoun = "cypher"
		c.subjectBVerb = "encrypted"
	case "signed_encrypted":
		c.subjectA = true
		c.subjectANoun = "signature"
		c.subjectAVerb = "signed"

		c.subjectB = true
		c.subjectBNoun = "cypher"
		c.subjectBVerb = "encrypted"

	case "verified":
		c.subjectA = true
		c.subjectANoun = "file"
		c.subjectAVerb = "verified"

		c.subjectB = false
		c.subjectBNoun = "file"
		c.subjectBVerb = "decrypted"
	case "decrypted":
		c.subjectA = false
		c.subjectANoun = "file"
		c.subjectAVerb = "verfied"

		c.subjectB = true
		c.subjectBNoun = "file"
		c.subjectBVerb = "decrypted"
	case "verified_decrypted":
		c.subjectA = true
		c.subjectANoun = "file"
		c.subjectAVerb = "verified"

		c.subjectB = true
		c.subjectBNoun = ""
		c.subjectBVerb = "decrypted"
	}

	return app.Div().Body(
		app.Select().
			Class("pf-c-form-control pf-u-mb-lg").
			Required(true).
			OnInput(func(ctx app.Context, e app.Event) {
				c.selectedMode = ctx.JSSrc().Get("value").String()
			}).
			Body(
				app.Option().
					Value("signed").
					Text("Signed").
					Selected(c.selectedMode == "signed"),
				app.Option().
					Value("encrypted").
					Text("Encrypted").
					Selected(c.selectedMode == "encrypted"),
				app.Option().
					Value("signed_encrypted").
					Text("Signed and encrypted").
					Selected(c.selectedMode == "signed_encrypted"),

				app.Option().
					Value("verified").
					Text("Verified").
					Selected(c.selectedMode == "verified"),
				app.Option().
					Value("decrypted").
					Text("Decrypted").
					Selected(c.selectedMode == "decrypted"),
				app.Option().
					Value("verified_decrypted").
					Text("Verified and decrypted").
					Selected(c.selectedMode == "verified_decrypted"),
			),
		app.Button().
			Class("pf-c-button pf-m-primary").
			Type("button").
			Text("Open post encrypt and sign modal").
			OnClick(func(ctx app.Context, e app.Event) {
				c.modalOpen = !c.modalOpen
			}),
		app.If(
			c.modalOpen,
			c.WithRoot(
				&components.DownloadOrViewModal{
					SubjectA:     c.subjectA,
					SubjectANoun: c.subjectANoun,
					SubjectAVerb: c.subjectAVerb,

					SubjectB:     c.subjectB,
					SubjectBNoun: c.subjectBNoun,
					SubjectBVerb: c.subjectBVerb,

					OnClose: func(_ bool) {
						c.modalOpen = false

						c.Update()
					},
					OnDownload: func() {
						app.Window().Call("alert", "Successfully downloaded")
					},
					OnView: func() {
						app.Window().Call("alert", "Successfully viewed")
					},
				},
			),
		),
	)
}
