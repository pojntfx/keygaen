package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type SuccessModal struct {
	app.Compo

	ID          string
	Icon        string
	Title       string
	Class       string
	Body        string
	ActionLabel string

	OnClose  func()
	OnAction func()
}

func (c *SuccessModal) Render() app.UI {
	return &Modal{
		ID:    c.ID,
		Icon:  c.Icon,
		Title: c.Title,
		Class: c.Class,
		Body:  app.Text(c.Body),
		Footer: app.Button().
			Aria("disabled", "false").
			Class("pf-c-button pf-m-primary").
			Type("button").
			Text(c.ActionLabel).
			OnClick(func(ctx app.Context, e app.Event) {
				c.OnAction()
			}),

		OnClose: func() {
			c.OnClose()
		},
	}
}
