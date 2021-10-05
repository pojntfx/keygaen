package stories

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type ImportKeyModalStory struct {
	Story
}

func (c *ImportKeyModalStory) Render() app.UI {
	return c.WithRoot(
		&components.ImportKeyModal{
			OnSubmit: func(key string) {
				app.Window().Call("alert", fmt.Sprintf("Imported key with contents %v", key))
			},
			OnCancel: func() {},
		},
	)
}
