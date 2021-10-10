package stories

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
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
					PrivateKeys: []components.EncryptionKey{
						{
							ID:       "123456",
							FullName: "Alice Example",
							Email:    "alice@example.com",
						},
						{
							ID:       "319312",
							FullName: "Bob Example",
							Email:    "bob@example.com",
						},
					},
					PublicKeys: []components.EncryptionKey{
						{
							ID:       "039292",
							FullName: "Isegard Example",
							Email:    "isegard@example.com",
						},
						{
							ID:       "838431",
							FullName: "Fred Example",
							Email:    "fred@example.com",
						},
					},

					OnSubmit: func(file []byte, publicKeyID, privateKeyID, detachedSignature string) {
						app.Window().Call("alert", fmt.Sprintf("Decrypted and verified file %v, using public key ID %v, private key ID %v and detached signature %v", file, publicKeyID, privateKeyID, detachedSignature))

						c.modalOpen = false
					},
					OnCancel: func() {
						c.modalOpen = false

						c.Update()
					},
				},
			),
		),
	)
}
