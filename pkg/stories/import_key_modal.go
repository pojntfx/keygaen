package stories

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygean/pkg/components"
)

type ImportKeyModalStory struct {
	Story

	modalOpen bool
}

func (c *ImportKeyModalStory) Render() app.UI {
	return app.Div().Body(
		app.Button().
			Class("pf-c-button pf-m-primary").
			Type("button").
			Text("Import key").
			OnClick(func(ctx app.Context, e app.Event) {
				c.modalOpen = !c.modalOpen
			}),
		app.If(
			c.modalOpen,
			c.WithRoot(
				&components.ImportKeyModal{
					OnSubmit: func(key []byte) {
						app.Window().Call("alert", fmt.Sprintf("Imported key with contents %v", key))

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
