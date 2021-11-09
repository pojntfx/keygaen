package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygean/pkg/components"
)

var demoKeys = []components.GPGKey{
	{
		ID:       "039292",
		Label:    "039292",
		FullName: "Isegard Example",
		Email:    "isegard@example.com",
		Private:  true,
		Content:  []byte{},
	},
	{
		ID:       "838431",
		Label:    "838431",
		FullName: "Fred Example",
		Email:    "fred@example.com",
		Private:  true,
		Public:   true,
		Content:  []byte{},
	},
	{
		ID:       "123456",
		Label:    "123456",
		FullName: "Alice Example",
		Email:    "alice@example.com",
		Public:   true,
		Content:  []byte{},
	},
	{
		ID:       "319312",
		Label:    "319312",
		FullName: "Bob Example",
		Email:    "bob@example.com",
		Public:   true,
		Content:  []byte{},
	},
}

type KeyListStory struct {
	Story
}

func (c *KeyListStory) Render() app.UI {
	return c.WithRoot(
		&components.KeyList{
			Keys: demoKeys,

			OnExport: func(keyID string) {
				app.Window().Call("alert", "Exported key "+keyID)
			},
			OnDelete: func(keyID string) {
				app.Window().Call("alert", "Deleted key "+keyID)
			},
		},
	)
}
