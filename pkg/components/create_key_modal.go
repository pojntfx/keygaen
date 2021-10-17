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
	OnCancel func(dirty bool, clear chan struct{})

	fullName             string
	email                string
	password             string
	passwordConfirmation string

	passwordInvalid bool

	dirty bool
}

func (c *CreateKeyModal) Render() app.UI {
	return &Modal{
		ID:           "create-key-modal",
		Title:        "Create Key",
		DisableFocus: true,
		Body: []app.UI{
			app.Form().
				Class("pf-c-form").
				ID("create-key-form").
				OnSubmit(func(ctx app.Context, e app.Event) {
					e.PreventDefault()

					// Check if the password confirmation matches
					if c.password != c.passwordConfirmation {
						c.passwordInvalid = true

						return
					}

					c.passwordInvalid = false

					// Submit the form
					c.OnSubmit(
						c.fullName,
						c.email,
						c.password,
					)

					c.clear()
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
									&Autofocused{
										Component: app.Input().
											Class("pf-c-form-control").
											Required(true).
											Type("text").
											Placeholder("Jean Doe").
											ID("full-name-input").
											Aria("describedby", "form-demo-basic-name-helper").
											OnInput(func(ctx app.Context, e app.Event) {
												c.fullName = ctx.JSSrc().Get("value").String()

												c.dirty = true
											}).
											Value(c.fullName),
									},
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
												Placeholder("jean@example.com").
												ID("email-input").
												Required(true).
												OnInput(func(ctx app.Context, e app.Event) {
													c.email = ctx.JSSrc().Get("value").String()

													c.dirty = true
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

													c.dirty = true
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
														Text("Confirm password"),
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
												Aria("invalid", c.passwordInvalid).
												Aria("describedby", func() string {
													if c.passwordInvalid {
														return "password-invalid-helper"
													}

													return ""
												}).
												Required(true).
												OnInput(func(ctx app.Context, e app.Event) {
													c.passwordConfirmation = ctx.JSSrc().Get("value").String()

													c.dirty = true
												}).
												Value(c.passwordConfirmation),
											app.If(
												c.passwordInvalid,
												app.P().
													Class("pf-c-form__helper-text pf-m-error").
													ID("password-invalid-helper").
													Aria("live", "polite").
													Text("The passwords don't match."),
											),
										),
								),
						),
				),
		},
		Footer: []app.UI{
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("submit").
				Text("Create key").
				Form("create-key-form"),
			app.Button().
				Class("pf-c-button pf-m-link").
				Type("button").
				Text("Cancel").
				OnClick(func(ctx app.Context, e app.Event) {
					handleCancel(c.clear, c.dirty, c.OnCancel)
				}),
		},
		OnClose: func() {
			handleCancel(c.clear, c.dirty, c.OnCancel)
		},
	}

}

func handleCancel(clear func(), dirty bool, cancel func(bool, chan struct{})) {
	done := make(chan struct{})

	go func() {
		<-done

		clear()
	}()
	cancel(dirty, done)
}

func (c *CreateKeyModal) clear() {
	c.fullName = ""
	c.email = ""
	c.password = ""
	c.passwordConfirmation = ""
}
