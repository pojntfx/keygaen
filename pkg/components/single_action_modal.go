package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// SingleActionModal is a modal which has only a singular action
type SingleActionModal struct {
	app.Compo

	ID          string // HTML ID of the modal; must be unique across the page
	Icon        string // Class of the icon to use to the left of the title; may be empty
	Title       string // Title of the modal
	Class       string // Class to be applied to the modal's outmost component
	Body        string // Body text of the modal
	ActionLabel string // Text to display on the modal's primary action

	OnClose  func() // Handler to call when closing/cancelling the modal
	OnAction func() // Handler to call when triggering the modal's primary action
}

func (c *SingleActionModal) Render() app.UI {
	return &Modal{
		ID:           c.ID,
		Icon:         c.Icon,
		Title:        c.Title,
		Class:        c.Class,
		DisableFocus: true,
		Body: []app.UI{
			app.Text(c.Body),
		},
		Footer: []app.UI{
			&Autofocused{
				Component: app.Button().
					Aria("disabled", "false").
					Class("pf-c-button pf-m-primary").
					Type("button").
					Text(c.ActionLabel).
					OnClick(func(ctx app.Context, e app.Event) {
						c.OnAction()
					}),
			},
		},

		OnClose: func() {
			c.OnClose()
		},
	}
}
