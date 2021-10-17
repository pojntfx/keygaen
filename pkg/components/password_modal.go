package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type PasswordModal struct {
	app.Compo

	Title    string
	OnSubmit func(password string)
	OnCancel func()

	password string
}

func (c *PasswordModal) Render() app.UI {
	return &Modal{
		ID:           "password-modal",
		Title:        c.Title,
		DisableFocus: true,
		Body: []app.UI{
			app.Form().
				Class("pf-c-form").
				ID("password-modal-form").
				OnSubmit(func(ctx app.Context, e app.Event) {
					e.PreventDefault()

					// Submit the form
					c.OnSubmit(
						c.password,
					)

					c.clear()
				}).
				Body(
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									&Autofocused{
										Component: app.Input().
											Class("pf-c-form-control").
											Required(true).
											Type("password").
											Placeholder("Password").
											Aria("label", "Password").
											OnInput(func(ctx app.Context, e app.Event) {
												c.password = ctx.JSSrc().Get("value").String()
											}).
											Value(c.password),
									},
								),
						),
				),
		},
		Footer: []app.UI{
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("submit").
				Form("password-modal-form").
				Text("Continue"),
			app.Button().
				Class("pf-c-button pf-m-link").
				Type("button").
				Text("Cancel").
				OnClick(func(ctx app.Context, e app.Event) {
					c.clear()
					c.OnCancel()
				}),
		},
		OnClose: func() {
			c.clear()
			c.OnCancel()
		},
	}
}

func (c *PasswordModal) clear() {
	c.password = ""
}
