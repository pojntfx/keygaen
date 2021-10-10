package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type PostEncryptAndSignModalStory struct {
	Story

	modalOpen bool

	selectedMode string

	signed    bool
	encrypted bool
}

func (c *PostEncryptAndSignModalStory) Render() app.UI {
	c.EnableShallowReflection()

	if c.selectedMode == "" {
		c.selectedMode = "signed_encrypted"
	}

	switch c.selectedMode {
	case "signed":
		c.signed = true
		c.encrypted = false
	case "encrypted":
		c.signed = false
		c.encrypted = true
	case "signed_encrypted":
		c.signed = true
		c.encrypted = true
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
				&components.PostEncryptAndSignModal{
					Signed:    c.signed,
					Encrypted: c.encrypted,

					OnClose: func() {
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
