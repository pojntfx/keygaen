package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ConfirmationModal struct {
	app.Compo

	ID          string
	Icon        string
	Title       string
	Class       string
	Body        string
	ActionClass string
	ActionLabel string
	CancelLabel string
	CancelLink  string

	OnClose  func()
	OnAction func()
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
