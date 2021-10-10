package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/components"
)

type ConfirmationModalStory struct {
	Story

	modalOpen bool
}

func (c *ConfirmationModalStory) Render() app.UI {
	return app.Div().Body(
		app.Button().
			Class("pf-c-button pf-m-danger").
			Type("button").
			Text("Delete key").
			OnClick(func(ctx app.Context, e app.Event) {
				c.modalOpen = !c.modalOpen
			}),
		app.If(
			c.modalOpen,
			c.WithRoot(
				&components.ConfirmationModal{
					ID:    "confirmation-modal-story",
					Icon:  "fas fa-exclamation-triangle",
					Title: "Are you sure?",
					Class: "pf-m-danger",
					Body:  "After deletion, you will not be able to restore the key.",

					ActionLabel: "Yes, delete the key",
					ActionClass: "pf-m-danger",

					CancelLabel: "Cancel",

					OnClose: func() {
						c.modalOpen = false

						c.Update()
					},
					OnAction: func() {
						c.modalOpen = false

						c.Update()
					},
				},
			),
		),
	)
}
