package stories

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type ImportKeyModalStory struct {
	Story

	modalOpen bool
}

func (c *ImportKeyModalStory) Render() app.UI {
	c.EnableShallowReflection()

	return c.WithRoot(
		app.Div().Body(
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("button").
				Text("Import Key").
				OnClick(func(ctx app.Context, e app.Event) {
					c.modalOpen = !c.modalOpen
				}),
			app.If(
				c.modalOpen,
				&components.ImportKeyModal{
					OnSubmit: func(key string) {
						app.Window().Call("alert", fmt.Sprintf("Imported key with contents %v", key))

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
