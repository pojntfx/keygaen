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
	c.EnableShallowReflection()

	return c.WithRoot(
		app.Div().Body(
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("button").
				Text("Encrypt/Sign").
				OnClick(func(ctx app.Context, e app.Event) {
					c.modalOpen = !c.modalOpen
				}),
			app.If(
				c.modalOpen,
				&components.EncryptAndSignModal{
					PrivateKeys: []components.EncryptionKey{
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
					PublicKeys: []components.EncryptionKey{
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

					OnSubmit: func(file []byte, publicKeyID, privateKeyID string, createDetachedSignature bool) {
						app.Window().Call("alert", fmt.Sprintf("Encrypted and signed file %v, using public key ID %v and private key ID %v and createDetachedSignature set to %v", file, publicKeyID, privateKeyID, createDetachedSignature))

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
