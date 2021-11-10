package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// ConfirmationModal is a modal with callbacks intended to enable confirm destructive operations such as deleting something
type ConfirmationModal struct {
	app.Compo

	ID          string // HTML ID of the modal; must be unique across the page
	Icon        string // Class of the icon to use to the left of the title; may be empty
	Title       string // Title of the modal
	Class       string // Class to be applied to the modal's outmost component
	Body        string // Body text of the modal
	ActionClass string // Class to be applied to the modal's primary action
	ActionLabel string // Text to display on the modal's primary action
	CancelLabel string // Text to display on the modal's cancel action
	CancelLink  string // Link to display as the cancel action; if empty, `OnClose` is being called

	OnClose  func() // Handler to call when closing/cancelling the modal
	OnAction func() // Handler to call when triggering the modal's primary action
}

func (c *ConfirmationModal) Render() app.UI {
	actionClass := "pf-c-button"
	if c.ActionClass == "" {
		actionClass += " " + "pf-m-primary"
	} else {
		actionClass += " " + c.ActionClass
	}

	return &Modal{
		ID:    c.ID,
		Icon:  c.Icon,
		Title: c.Title,
		Class: c.Class,
		Body: []app.UI{
			app.Text(c.Body),
		},
		Footer: []app.UI{
			app.Button().
				Aria("disabled", "false").
				Class(actionClass).
				Type("button").
				Text(c.ActionLabel).
				OnClick(func(ctx app.Context, e app.Event) {
					c.OnAction()
				}),
			app.If(
				c.CancelLink == "",
				app.Button().
					Class("pf-c-button pf-m-link").
					Type("button").
					Text(c.CancelLabel).
					OnClick(func(ctx app.Context, e app.Event) {
						c.OnClose()
					}),
			).Else(
				app.A().
					Class("pf-c-button pf-m-link").
					Target("_blank").
					Href(c.CancelLink).
					Text(c.CancelLabel),
			),
		},

		OnClose: func() {
			c.OnClose()
		},
	}
}
