package stories

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/components"
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
					PublicKey: true,
					OnDownloadPublicKey: func(armor, base64encode bool) {
						app.Window().Call("alert", fmt.Sprintf("Downloaded public key with armor set to %v with base64encode %v", armor, base64encode))
					},
					OnViewPublicKey: func(armor, base64encode bool) {
						app.Window().Call("alert", fmt.Sprintf("Viewed public key with armor set to %v with base64encode %v", armor, base64encode))
					},

					PrivateKey: true,
					OnDownloadPrivateKey: func(armor, base64encode bool) {
						app.Window().Call("alert", fmt.Sprintf("Downloaded private key with armor set to %v with base64encode %v", armor, base64encode))
					},
					OnViewPrivateKey: func(armor, base64encode bool) {
						app.Window().Call("alert", fmt.Sprintf("Viewed private key with armor set to %v with base64encode %v", armor, base64encode))
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
