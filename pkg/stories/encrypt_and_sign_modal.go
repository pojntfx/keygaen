package stories

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type EncryptAndSignModalStory struct {
	Story

	modalOpen bool
}

func (c *EncryptAndSignModalStory) Render() app.UI {
	return app.Div().Body(
		app.Button().
			Class("pf-c-button pf-m-primary").
			Type("button").
			Text("Encrypt/sign").
			OnClick(func(ctx app.Context, e app.Event) {
				c.modalOpen = !c.modalOpen
			}),
		app.If(
			c.modalOpen,
			c.WithRoot(
				&components.EncryptAndSignModal{
					Keys: demoKeys,

					OnSubmit: func(file []byte, publicKeyID, privateKeyID string, createDetachedSignature, armor bool) {
						app.Window().Call("alert", fmt.Sprintf("Encrypted and signed file %v, using public key ID %v and private key ID %v, createDetachedSignature set to %v and armor set to %v", file, publicKeyID, privateKeyID, createDetachedSignature, armor))

						c.modalOpen = false
					},
					OnCancel: func(dirty bool, clear chan struct{}) {
						c.modalOpen = false

						c.Update()

						clear <- struct{}{}
					},
				},
			),
		),
	)
}
