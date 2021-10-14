package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Toolbar struct {
	app.Compo

	OnCreateKey func()
	OnImportKey func()

	OnEncryptAndSign   func()
	OnDecryptAndVerify func()
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
										Class("pf-c-button pf-m-plain").
										OnClick(func(ctx app.Context, e app.Event) {
											c.OnCreateKey()
										}).
										Body(
											app.Span().
												Class("pf-c-button__icon").
												Body(
													app.I().
														Class("fas fa-plus").
														Aria("hidden", true),
												),
										),
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										OnClick(func(ctx app.Context, e app.Event) {
											c.OnImportKey()
										}).
										Body(
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
												Class("pf-c-button pf-m-plain").
												Type("button").
												OnClick(func(ctx app.Context, e app.Event) {
													c.OnEncryptAndSign()
												}).
												Body(
													app.Span().
														Class("pf-c-button__icon").
														Body(
															app.I().
																Class("fas fa-lock").
																Aria("hidden", true),
														),
												),
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												OnClick(func(ctx app.Context, e app.Event) {
													c.OnDecryptAndVerify()
												}).
												Body(
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
													app.Text(" Create Key"),
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
													app.Text(" Import Key"),
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
															app.Text(" Encrypt/Sign"),
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
															app.Text(" Decrypt/Verify"),
														),
												),
										),
								),
						),
				),
		)
}
