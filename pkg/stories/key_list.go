package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type KeyListStory struct {
	Story
}

func (c *KeyListStory) Render() app.UI {
	return c.WithRoot(
		&components.KeyList{
			Keys: []components.GPGKey{
				{
					ID:       "039292",
					FullName: "Isegard Example",
					Email:    "isegard@example.com",
					Private:  true,
				},
				{
					ID:       "838431",
					FullName: "Fred Example",
					Email:    "fred@example.com",
					Private:  true,
					Public:   true,
				},
				{
					ID:       "123456",
					FullName: "Alice Example",
					Email:    "alice@example.com",
					Public:   true,
				},
				{
					ID:       "319312",
					FullName: "Bob Example",
					Email:    "bob@example.com",
					Public:   true,
				},
			},
		},
	)
}
