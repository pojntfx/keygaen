package components

import (
	app "github.com/maxence-charriere/go-app/v9/pkg/app"
)

type PasswordModal struct {
	app.Compo

	OnSubmit func(password string)
	OnCancel func()

	password string
}

func (c *PasswordModal) Render() app.UI {
	return app.Form().
		Class("pf-c-form").
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
							app.Input().
								Class("pf-c-form-control").
								Required(true).
								Type("password").
								Placeholder("Enter password").
								Aria("label", "Enter password").
								OnInput(func(ctx app.Context, e app.Event) {
									c.password = ctx.JSSrc().Get("value").String()
								}).
								Value(c.password),
						),
				),
			app.Div().
				Class("pf-c-form__group pf-m-action").
				Body(
					app.Div().
						Class("pf-c-form__group-control").
						Body(
							app.Div().
								Class("pf-c-form__actions").
								Body(
									app.Button().
										Class("pf-c-button pf-m-primary").
										Type("submit").
										Text("Continue"),
									app.Button().
										Class("pf-c-button pf-m-link").
										Type("button").
										Text("Cancel").
										OnClick(func(ctx app.Context, e app.Event) {
											c.clear()

											c.OnCancel()
										}),
								),
						),
				),
		)
}

func (c *PasswordModal) clear() {
	c.password = ""
}
