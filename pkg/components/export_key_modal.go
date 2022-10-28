package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	exportPublicKeyForm  = "export-public-key-form"
	exportPrivateKeyForm = "export-private-key-form"
)

// ExportKeyModal is a modal which provides the actions needed to export a key
type ExportKeyModal struct {
	app.Compo

	PublicKey           bool                           // Whether to display the options for a public key
	OnDownloadPublicKey func(armor, base64encode bool) // Handler to call to download the public key
	OnViewPublicKey     func(armor, base64encode bool) // Handler to call to view the public key

	PrivateKey           bool                           // Whether to display the options for a private key
	OnDownloadPrivateKey func(armor, base64encode bool) // Handler to call to download the private key
	OnViewPrivateKey     func(armor, base64encode bool) // Handler to call to view the private key

	OnOK func() // Handler to call when dismissing the modal

	skipPublicKeyArmor    bool
	publicKeyBase64Encode bool

	skipPrivateKeyArmor    bool
	privateKeyBase64Encode bool
}

func (c *ExportKeyModal) Render() app.UI {
	return &Modal{
		ID:    "export-key-modal",
		Title: "Export Key",
		Body: []app.UI{
			app.If(
				c.PublicKey,
				app.Div().
					Class("pf-c-card pf-m-compact pf-m-flat").
					Body(
						app.Div().
							Class("pf-c-card__title").
							Body(
								app.I().
									Class("fas fa-globe pf-u-mr-sm"),
								app.Text("Public Key"),
							),
						app.Div().
							Class("pf-c-card__body").
							Body(
								app.Text("Anyone can use this key to encrypt messages to you and verify your identity; you may share it with the public."),
								app.Form().
									Class("pf-c-form pf-u-mt-lg").
									ID(exportPublicKeyForm).
									OnSubmit(func(ctx app.Context, e app.Event) {
										e.PreventDefault()
									}).
									Body(
										app.Div().
											Aria("role", "group").
											Class("pf-c-form__group").
											Body(
												app.Div().
													Class("pf-c-form__group-control").
													Body(
														app.Div().
															Class("pf-c-check").
															Body(
																&Controlled{
																	Component: app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("public-armor-checkbox").
																		OnInput(func(ctx app.Context, e app.Event) {
																			c.skipPublicKeyArmor = !c.skipPublicKeyArmor
																		}),
																	Properties: map[string]interface{}{
																		"checked": !c.skipPublicKeyArmor,
																	},
																},
																app.Label().
																	Class("pf-c-check__label").
																	For("public-armor-checkbox").
																	Body(
																		app.I().
																			Class("fas fa-shield-alt pf-u-mr-sm"),
																		app.Text("Armor"),
																	),
																app.Span().
																	Class("pf-c-check__description").
																	Text("To increase portability, ASCII armor the key."),
															),
													),
											),
										app.Div().
											Aria("role", "group").
											Class("pf-c-form__group").
											Body(
												app.Div().
													Class("pf-c-form__group-control").
													Body(
														app.Div().
															Class("pf-c-check").
															Body(
																&Controlled{
																	Component: app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("public-base64-checkbox").
																		OnInput(func(ctx app.Context, e app.Event) {
																			c.publicKeyBase64Encode = !c.publicKeyBase64Encode
																		}),
																	Properties: map[string]interface{}{
																		"checked": c.publicKeyBase64Encode,
																	},
																},
																app.Label().
																	Class("pf-c-check__label").
																	For("public-base64-checkbox").
																	Body(
																		app.I().
																			Class("fas fa-shield-alt pf-u-mr-sm"),
																		app.Text("Base64 encode"),
																	),
																app.Span().
																	Class("pf-c-check__description").
																	Text("Use a reduced alphabet for better portability."),
															),
													),
											),
									),
							),
						app.Div().
							Class("pf-c-card__footer").
							Body(
								app.Button().
									Class(func() string {
										classes := "pf-c-button pf-m-control pf-u-mr-sm pf-u-display-block pf-u-display-inline-block-on-md pf-u-w-100 pf-u-w-initial-on-md"
										if !c.skipPublicKeyArmor {
											classes += " pf-u-mb-md pf-u-mb-0-on-md"
										}

										return classes
									}()).
									Type("submit").
									Form(exportPublicKeyForm).
									OnClick(func(ctx app.Context, e app.Event) {
										c.OnDownloadPublicKey(!c.skipPublicKeyArmor, c.publicKeyBase64Encode)
									}).
									Body(
										app.Span().
											Class("pf-c-button__icon pf-m-start").
											Body(
												app.I().
													Class("fas fa-download").
													Aria("hidden", true),
											),
										app.Text("Download public key"),
									),
								app.If(
									!c.skipPublicKeyArmor || c.publicKeyBase64Encode,
									app.Button().
										Class("pf-c-button pf-m-control pf-u-mr-sm pf-u-display-block pf-u-display-inline-block-on-md pf-u-w-100 pf-u-w-initial-on-md").
										Type("submit").
										Form(exportPublicKeyForm).
										OnClick(func(ctx app.Context, e app.Event) {
											c.OnViewPublicKey(!c.skipPublicKeyArmor, c.publicKeyBase64Encode)
										}).
										Body(
											app.Span().
												Class("pf-c-button__icon pf-m-start").
												Body(
													app.I().
														Class("fas fa-eye").
														Aria("hidden", true),
												),
											app.Text("View public key"),
										),
								),
							),
					),
			),
			app.If(
				c.PrivateKey,
				app.Div().
					Class("pf-c-card pf-m-compact pf-m-flat pf-u-mt-md").
					Body(
						app.Div().
							Class("pf-c-card__title").
							Body(
								app.I().
									Class("fas fa-user-lock pf-u-mr-sm"),
								app.Text("Private Key"),
							),
						app.Div().
							Class("pf-c-card__body").
							Body(
								app.Text("You can use this key to decrypt messages and sign with your identity; don't share it with anyone."),
								app.Form().
									Class("pf-c-form pf-u-mt-lg").
									ID(exportPrivateKeyForm).
									OnSubmit(func(ctx app.Context, e app.Event) {
										e.PreventDefault()
									}).
									Body(
										app.Div().
											Aria("role", "group").
											Class("pf-c-form__group").
											Body(
												app.Div().
													Class("pf-c-form__group-control").
													Body(
														app.Div().
															Class("pf-c-check").
															Body(
																&Controlled{
																	Component: app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("private-armor-checkbox").
																		OnInput(func(ctx app.Context, e app.Event) {
																			c.skipPrivateKeyArmor = !c.skipPrivateKeyArmor
																		}),
																	Properties: map[string]interface{}{
																		"checked": !c.skipPrivateKeyArmor,
																	},
																},
																app.Label().
																	Class("pf-c-check__label").
																	For("private-armor-checkbox").
																	Body(
																		app.I().
																			Class("fas fa-shield-alt pf-u-mr-sm"),
																		app.Text("Armor"),
																	),
																app.Span().
																	Class("pf-c-check__description").
																	Text("To increase portability, ASCII armor the key."),
															),
													),
											),
										app.Div().
											Aria("role", "group").
											Class("pf-c-form__group").
											Body(
												app.Div().
													Class("pf-c-form__group-control").
													Body(
														app.Div().
															Class("pf-c-check").
															Body(
																&Controlled{
																	Component: app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("private-base64-checkbox").
																		OnInput(func(ctx app.Context, e app.Event) {
																			c.privateKeyBase64Encode = !c.privateKeyBase64Encode
																		}),
																	Properties: map[string]interface{}{
																		"checked": c.privateKeyBase64Encode,
																	},
																},
																app.Label().
																	Class("pf-c-check__label").
																	For("private-base64-checkbox").
																	Body(
																		app.I().
																			Class("fas fa-shield-alt pf-u-mr-sm"),
																		app.Text("Base64 encode"),
																	),
																app.Span().
																	Class("pf-c-check__description").
																	Text("Use a reduced alphabet for better portability."),
															),
													),
											),
									),
							),
						app.Div().
							Class("pf-c-card__footer").
							Body(
								app.Button().
									Class(func() string {
										classes := "pf-c-button pf-m-control pf-u-mr-sm pf-u-display-block pf-u-display-inline-block-on-md pf-u-w-100 pf-u-w-initial-on-md"
										if !c.skipPrivateKeyArmor {
											classes += " pf-u-mb-md pf-u-mb-0-on-md"
										}

										return classes
									}()).
									Type("submit").
									Form(exportPrivateKeyForm).
									OnClick(func(ctx app.Context, e app.Event) {
										c.OnDownloadPrivateKey(!c.skipPrivateKeyArmor, c.privateKeyBase64Encode)
									}).
									Body(
										app.Span().
											Class("pf-c-button__icon pf-m-start").
											Body(
												app.I().
													Class("fas fa-download").
													Aria("hidden", true),
											),
										app.Text("Download private key"),
									),
								app.If(
									!c.skipPrivateKeyArmor || c.privateKeyBase64Encode,
									app.Button().
										Class("pf-c-button pf-m-control pf-u-mr-sm pf-u-display-block pf-u-display-inline-block-on-md pf-u-w-100 pf-u-w-initial-on-md").
										Type("submit").
										Form(exportPrivateKeyForm).
										OnClick(func(ctx app.Context, e app.Event) {
											c.OnViewPrivateKey(!c.skipPrivateKeyArmor, c.privateKeyBase64Encode)
										}).
										Body(
											app.Span().
												Class("pf-c-button__icon pf-m-start").
												Body(
													app.I().
														Class("fas fa-eye").
														Aria("hidden", true),
												),
											app.Text("View private key"),
										),
								),
							),
					),
			),
		},
		Footer: []app.UI{
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("button").
				Text("OK").
				OnClick(func(ctx app.Context, e app.Event) {
					c.clear()
					c.OnOK()
				}),
		},
		OnClose: func() {
			c.clear()
			c.OnOK()
		},
	}
}

func (c *ExportKeyModal) clear() {
	c.skipPublicKeyArmor = false
	c.publicKeyBase64Encode = false

	c.skipPrivateKeyArmor = false
	c.privateKeyBase64Encode = false
}
