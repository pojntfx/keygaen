package stories

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type ExportKeyModalStory struct {
	Story

	modalOpen bool
}

func (c *ExportKeyModalStory) Render() app.UI {
	return app.Div().Body(
		app.Button().
			Class("pf-c-button pf-m-primary").
			Type("button").
			Text("Export key").
			OnClick(func(ctx app.Context, e app.Event) {
				c.modalOpen = !c.modalOpen
			}),
		app.If(
			c.modalOpen,
			c.WithRoot(
				&components.ExportKeyModal{
					OnDownloadPublicKey: func(armor bool) {
						app.Window().Call("alert", fmt.Sprintf("Downloaded public key with armor set to %v", armor))
					},
					OnViewPublicKey: func() {
						app.Window().Call("alert", "Viewed public key")
					},

					OnDownloadPrivateKey: func(armor bool) {
						app.Window().Call("alert", fmt.Sprintf("Downloaded private key with armor set to %v", armor))
					},
					OnViewPrivateKey: func() {
						app.Window().Call("alert", "Viewed private key")
					},

					OnOK: func() {
						c.modalOpen = false

						c.Update()
					},
				},
			),
		),
	)
}
