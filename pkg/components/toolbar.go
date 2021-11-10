package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Toolbar is the secondary navigation menu and contains most actions
type Toolbar struct {
	app.Compo

	OnCreateKey func() // Handler to call to create a key
	OnImportKey func() // Handler to call to import a key

	OnEncryptAndSign   func() // Handler to call to encrypt/sign
	OnDecryptAndVerify func() // Handler to call to decrypt/verify
}

func (c *Toolbar) Render() app.UI {
	return app.Div().
		Class("pf-c-toolbar").
		Body(
			app.Div().
				Class("pf-c-toolbar__content").
				Body(
					app.Div().
						Class("pf-c-toolbar__content-section pf-m-nowrap pf-u-display-none-on-md").
						Body(
							app.Div().
								Class("pf-c-toolbar__group").
								Body(
									app.Button().
										Type("button").
										Class("pf-c-button pf-m-plain pf-x-c-tooltip-parent").
										OnClick(func(ctx app.Context, e app.Event) {
											c.OnCreateKey()
										}).
										Body(
											app.Div().
												Class("pf-c-tooltip pf-m-right pf-x-c-tooltip").
												Aria("role", "tooltip").
												Body(
													app.Div().
														Class("pf-c-tooltip__arrow"),
													app.Div().
														Class("pf-c-tooltip__content").
														Body(
															app.Text("Create key"),
														),
												),
											app.Span().
												Class("pf-c-button__icon").
												Body(
													app.I().
														Class("fas fa-plus").
														Aria("hidden", true),
												),
										),
									app.Button().
										Class("pf-c-button pf-m-plain pf-x-c-tooltip-parent").
										Type("button").
										OnClick(func(ctx app.Context, e app.Event) {
											c.OnImportKey()
										}).
										Body(
											app.Div().
												Class("pf-c-tooltip pf-m-right pf-x-c-tooltip").
												Aria("role", "tooltip").
												Body(
													app.Div().
														Class("pf-c-tooltip__arrow"),
													app.Div().
														Class("pf-c-tooltip__content").
														Body(
															app.Text("Import key"),
														),
												),
											app.Span().
												Class("pf-c-button__icon").
												Body(
													app.I().
														Aria("hidden", true).
														Class("fas fa-file-import"),
												),
										),
								),
							app.Div().
								Class("pf-c-toolbar__item pf-m-pagination").
								Body(
									app.Div().
										Class("pf-c-toolbar__group").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain pf-x-c-tooltip-parent").
												Type("button").
												OnClick(func(ctx app.Context, e app.Event) {
													c.OnEncryptAndSign()
												}).
												Body(
													app.Div().
														Class("pf-c-tooltip pf-m-left pf-x-c-tooltip").
														Aria("role", "tooltip").
														Body(
															app.Div().
																Class("pf-c-tooltip__arrow"),
															app.Div().
																Class("pf-c-tooltip__content").
																Body(
																	app.Text("Encrypt/sign"),
																),
														),
													app.Span().
														Class("pf-c-button__icon").
														Body(
															app.I().
																Class("fas fa-lock").
																Aria("hidden", true),
														),
												),
											app.Button().
												Class("pf-c-button pf-m-plain pf-x-c-tooltip-parent").
												Type("button").
												OnClick(func(ctx app.Context, e app.Event) {
													c.OnDecryptAndVerify()
												}).
												Body(
													app.Div().
														Class("pf-c-tooltip pf-m-left pf-x-c-tooltip").
														Aria("role", "tooltip").
														Body(
															app.Div().
																Class("pf-c-tooltip__arrow"),
															app.Div().
																Class("pf-c-tooltip__content").
																Body(
																	app.Text("Decrypt/verify"),
																),
														),
													app.Span().
														Class("pf-c-button__icon").
														Body(
															app.I().
																Class("fas fa-lock-open").
																Aria("hidden", true),
														),
												),
										),
								),
						),
					app.Div().
						Class("pf-c-toolbar__content-section pf-m-nowrap pf-u-display-none pf-u-display-flex-on-md").
						Body(
							app.Div().
								Class("pf-c-toolbar__group").
								Body(
									app.Div().
										Class("pf-c-toolbar__item ").
										Body(
											app.Button().
												Class("pf-c-button pf-m-control").
												Type("button").
												OnClick(func(ctx app.Context, e app.Event) {
													c.OnCreateKey()
												}).
												Body(
													app.Span().
														Class("pf-c-button__icon pf-m-start").
														Body(
															app.I().
																Class("fas fa-plus").
																Aria("hidden", true),
														),
													app.Text(" Create key"),
												),
										),
									app.Div().
										Class("pf-c-toolbar__item").
										Body(
											app.Button().
												Class("pf-c-button pf-m-control").
												Type("button").
												OnClick(func(ctx app.Context, e app.Event) {
													c.OnImportKey()
												}).
												Body(
													app.Span().
														Class("pf-c-button__icon pf-m-start").
														Body(
															app.I().
																Class("fas fa-file-import").
																Aria("hidden", true),
														),
													app.Text(" Import key"),
												),
										),
								),
							app.Div().
								Class("pf-c-toolbar__item pf-m-pagination").
								Body(
									app.Div().
										Class("pf-c-toolbar__group").
										Body(
											app.Div().
												Class("pf-c-toolbar__item ").
												Body(
													app.Button().
														Class("pf-c-button pf-m-control").
														Type("button").
														OnClick(func(ctx app.Context, e app.Event) {
															c.OnEncryptAndSign()
														}).
														Body(
															app.Span().
																Class("pf-c-button__icon pf-m-start").
																Body(
																	app.I().
																		Class("fas fa-lock").
																		Aria("hidden", true),
																),
															app.Text(" Encrypt/sign"),
														),
												),
											app.Div().
												Class("pf-c-toolbar__item").
												Body(
													app.Button().
														Class("pf-c-button pf-m-control").
														Type("button").
														OnClick(func(ctx app.Context, e app.Event) {
															c.OnDecryptAndVerify()
														}).
														Body(
															app.Span().
																Class("pf-c-button__icon pf-m-start").
																Body(
																	app.I().
																		Class("fas fa-lock-open").
																		Aria("hidden", true),
																),
															app.Text(" Decrypt/verify"),
														),
												),
										),
								),
						),
				),
		)
}
