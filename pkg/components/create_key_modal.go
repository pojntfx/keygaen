package components

import (
	app "github.com/maxence-charriere/go-app/v9/pkg/app"
)

type CreateKeyModal struct {
	app.Compo

	OnSubmit func(
		fullName string,
		email string,
		password string,
	)

	fullName             string
	email                string
	password             string
	passwordConfirmation string
}

func (c *CreateKeyModal) Render() app.UI {
	return app.Form().
		Class("pf-c-form").
		OnSubmit(func(ctx app.Context, e app.Event) {
			e.PreventDefault()

			// TODO: Call `OnSubmit`
		}).
		Body(
			app.Div().
				Class("pf-c-form__group").
				Body(
					app.Div().
						Class("pf-c-form__group-label").
						Body(
							app.Label().
								Class("pf-c-form__label").
								For("full-name-input").
								Body(
									app.Span().
										Class("pf-c-form__label-text").
										Text("Full name"),
									app.Span().
										Class("pf-c-form__label-required").
										Aria("hidden", true).
										Text("*"),
								),
						),
					app.Div().
						Class("pf-c-form__group-control").
						Body(
							app.Input().
								Class("pf-c-form-control").
								Required(true).
								Type("text").
								ID("full-name-input").
								Aria("describedby", "form-demo-basic-name-helper").
								OnInput(func(ctx app.Context, e app.Event) {
									c.fullName = ctx.JSSrc().Get("value").String()
								}).
								Value(c.fullName),
						),
				),
			app.Div().
				Class("pf-c-form__group").
				Body(
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-label").
								Body(
									app.Label().
										Class("pf-c-form__label").
										For("email-input").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Text("Email"),
											app.Span().
												Class("pf-c-form__label-required").
												Aria("hidden", true).
												Text("*"),
										),
								),
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Input().
										Class("pf-c-form-control").
										Type("email").
										ID("email-input").
										Required(true).
										OnInput(func(ctx app.Context, e app.Event) {
											c.email = ctx.JSSrc().Get("value").String()
										}).
										Value(c.email),
								),
						),
				),
			app.Div().
				Class("pf-c-form__group").
				Body(
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-label").
								Body(
									app.Label().
										Class("pf-c-form__label").
										For("password-input").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Text("Password"),
											app.Span().
												Class("pf-c-form__label-required").
												Aria("hidden", true).
												Text("*"),
										),
								),
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Input().
										Class("pf-c-form-control").
										Type("password").
										ID("password-input").
										Required(true).
										OnInput(func(ctx app.Context, e app.Event) {
											c.password = ctx.JSSrc().Get("value").String()
										}).
										Value(c.password),
								),
						),
				),
			app.Div().
				Class("pf-c-form__group").
				Body(
					app.Div().
						Class("pf-c-form__group").
						Body(
							app.Div().
								Class("pf-c-form__group-label").
								Body(
									app.Label().
										Class("pf-c-form__label").
										For("confirm-password-input").
										Body(
											app.Span().
												Class("pf-c-form__label-text").
												Text("Confirm Password"),
											app.Span().
												Class("pf-c-form__label-required").
												Aria("hidden", true).
												Text("*"),
										),
								),
							app.Div().
								Class("pf-c-form__group-control").
								Body(
									app.Input().
										Class("pf-c-form-control").
										Type("password").
										ID("confirm-password-input").
										Required(true).
										OnInput(func(ctx app.Context, e app.Event) {
											c.passwordConfirmation = ctx.JSSrc().Get("value").String()
										}).
										Value(c.passwordConfirmation),
								),
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
										Text("Create Key"),
									app.Button().
										Class("pf-c-button pf-m-link").
										Type("button").
										Text("Cancel"),
								),
						),
				),
		)
}
