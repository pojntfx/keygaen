package stories

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type EncryptAndSignModalStory struct {
	Story
}

func (c *EncryptAndSignModalStory) Render() app.UI {
	return c.WithRoot(
		&components.EncryptAndSignModal{
			PrivateKeys: []components.EncryptionKey{},
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

			OnSubmit: func(file []byte, publicKeyID, privateKeyID string) {
				app.Window().Call("alert", fmt.Sprintf("Encrypted and signed file %v, using public key ID %v and private key ID %v", file, publicKeyID, privateKeyID))
			},
			OnCancel: func() {},
		},
	)
}
