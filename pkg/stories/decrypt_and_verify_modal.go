package stories

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/components"
)

type DecryptAndVerifyModalStory struct {
	Story

	modalOpen bool
}

func (c *DecryptAndVerifyModalStory) Render() app.UI {
	return app.Div().Body(
		app.Button().
			Class("pf-c-button pf-m-primary").
			Type("button").
			Text("Decrypt/verify").
			OnClick(func(ctx app.Context, e app.Event) {
				c.modalOpen = !c.modalOpen
			}),
		app.If(
			c.modalOpen,
			c.WithRoot(
				&components.DecryptAndVerifyModal{
					Keys: demoKeys,

					OnSubmit: func(file []byte, publicKeyID, privateKeyID string, detachedSignature []byte) {
						app.Window().Call("alert", fmt.Sprintf("Decrypted and verified file %v, using public key ID %v, private key ID %v and detached signature %v", file, publicKeyID, privateKeyID, detachedSignature))

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
