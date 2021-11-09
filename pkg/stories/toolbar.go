package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/components"
)

type ToolbarStory struct {
	Story
}

func (c *ToolbarStory) Render() app.UI {
	return c.WithRoot(
		&components.Toolbar{
			OnCreateKey: func() {
				app.Window().Call("alert", "Created key")
			},
			OnImportKey: func() {
				app.Window().Call("alert", "Imported key")
			},

			OnEncryptAndSign: func() {
				app.Window().Call("alert", "Encrypted and signed")
			},
			OnDecryptAndVerify: func() {
				app.Window().Call("alert", "Decrypted and verified")
			},
		},
	)
}
