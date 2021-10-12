package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type EmptyStateStory struct {
	Story
}

func (c *EmptyStateStory) Render() app.UI {
	return c.WithRoot(
		&components.EmptyState{
			OnCreateKey: func() {
				app.Window().Call("alert", "Created key")
			},
			OnImportKey: func() {
				app.Window().Call("alert", "Imported key")
			},
		},
	)
}
