package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	selectEncryptionFileInput = "select-encryption-file-input"
)

// EncryptAndSignModal is a modal which provides the information needed to encrypt/sign something
type EncryptAndSignModal struct {
	app.Compo

	Keys []PGPKey // PGP keys to be available for encryption/signing

	OnSubmit func(
		file []byte,
		publicKeyID string,
		privateKeyID string,
		createDetachedSignature bool,
		armor bool,
	) // Handle to call to encrypt/sign
	OnCancel func(dirty bool, clear chan struct{}) // Handler to call when closing/cancelling the modal

	fileIsBinary bool
	fileContents []byte

	skipEncryption bool
	publicKeyID    string

	skipSigning  bool
	privateKeyID string

	createDetachedSignature bool
	skipArmor               bool

	dirty bool
}

func (c *EncryptAndSignModal) Render() app.UI {
	privateKeys := []PGPKey{}
	publicKeys := []PGPKey{}
	for _, key := range c.Keys {
		if key.Private {
			privateKeys = append(privateKeys, key)
		}

		if key.Public {
			publicKeys = append(publicKeys, key)
		}
	}

	return &Modal{
		ID:    "encrypt-and-sign-modal",
		Title: "Encrypt/Sign",
		Body: []app.UI{
			app.Form().
				Class("pf-c-form").
				ID("encrypt-and-sign-form").
				OnSubmit(func(ctx app.Context, e app.Event) {
					e.PreventDefault()

					c.OnSubmit(
						c.fileContents,
						c.publicKeyID,
						c.privateKeyID,
						c.createDetachedSignature,
						!c.skipArmor,
					)

					c.clear()
				}).
				Body(
					app.Div().
						Class("pf-c-form__group").
						Body(
							&FileUpload{
								ID:                         selectEncryptionFileInput,
								FileSelectionLabel:         "Drag and drop a file or select one",
								ClearLabel:                 "Clear",
								TextEntryInputPlaceholder:  "Or enter text here",
								TextEntryInputBlockedLabel: "File has been selected.",
								FileContents:               c.fileContents,

								OnChange: func(fileContents []byte) {
									c.fileContents = fileContents

									c.dirty = true
								},
								OnClear: func() {
									c.fileContents = []byte{}
								},
							},
						),
					app.Div().
						Class("pf-c-form__group").
						Aria("role", "group").
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
													ID("encryption-checkbox").
													OnInput(func(ctx app.Context, e app.Event) {
														if !(c.skipSigning && !c.skipEncryption) {
															c.skipEncryption = !c.skipEncryption
														}

														if c.skipEncryption {
															c.publicKeyID = ""
														}

														c.dirty = true
													}),
												Properties: map[string]interface{}{
													"checked": !c.skipEncryption,
												},
											},
											app.Label().
												Class("pf-c-check__label").
												For("encryption-checkbox").
												Body(
													app.I().
														Class("fas fa-lock pf-u-mr-sm"),
													app.Text("Encrypt file"),
												),
											app.If(
												c.skipEncryption,
												app.Span().
													Class("pf-c-check__description").
													Text("If enabled, only the person with the correct key will be able to read the message."),
											).Else(
												app.Span().
													Class("pf-c-check__description").
													Text("Allow only the person with the following key to read the message:"),
												app.Div().
													Class("pf-c-check__body pf-u-w-100").
													Body(
														app.Select().
															Class("pf-c-form-control").
															ID("public-key-selector").
															Required(true).
															OnInput(func(ctx app.Context, e app.Event) {
																c.publicKeyID = ctx.JSSrc().Get("value").String()

																c.dirty = true
															}).
															Body(
																app.Option().
																	Value("").
																	Text("Select one").
																	Selected(c.publicKeyID == ""),
																app.Range(publicKeys).Slice(func(i int) app.UI {
																	return app.Option().
																		Value(publicKeys[i].ID).
																		Text(getKeySummary(publicKeys[i])).
																		Selected(c.publicKeyID == publicKeys[i].ID)
																}),
															),
													),
											),
										),
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Aria("role", "group").
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
													ID("signature-checkbox").
													OnInput(func(ctx app.Context, e app.Event) {
														if !(!c.skipSigning && c.skipEncryption) {
															c.skipSigning = !c.skipSigning
														}

														if c.skipSigning {
															c.privateKeyID = ""
														}

														c.dirty = true
													}),
												Properties: map[string]interface{}{
													"checked": !c.skipSigning,
												},
											},
											app.Label().
												Class("pf-c-check__label").
												For("signature-checkbox").
												Body(
													app.I().
														Class("fas fa-signature pf-u-mr-sm"),
													app.Text("Sign file"),
												),
											app.If(
												c.skipSigning,
												app.Span().
													Class("pf-c-check__description").
													Text("If enabled, anyone will be able to verify that the file originates from the person with the selected key."),
											).Else(
												app.Span().
													Class("pf-c-check__description").
													Text("This will anyone to verify that the file originates from the person with the following key:"),
												app.Div().
													Class("pf-c-check__body pf-u-w-100").
													Body(
														app.Select().
															Class("pf-c-form-control").
															ID("private-key-selector").
															Required(true).
															OnInput(func(ctx app.Context, e app.Event) {
																c.privateKeyID = ctx.JSSrc().Get("value").String()

																c.dirty = true
															}).
															Body(
																app.Option().
																	Value("").
																	Text("Select one").
																	Selected(c.privateKeyID == ""),
																app.Range(privateKeys).Slice(func(i int) app.UI {
																	return app.Option().
																		Value(privateKeys[i].ID).
																		Text(getKeySummary(privateKeys[i])).
																		Selected(c.privateKeyID == privateKeys[i].ID)
																}),
															),
														app.Div().
															Class("pf-c-form__group pf-u-mt-lg").
															Aria("role", "group").
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
																						ID("detached-signature-checkbox").
																						OnInput(func(ctx app.Context, e app.Event) {
																							c.createDetachedSignature = !c.createDetachedSignature

																							c.dirty = true
																						}),
																					Properties: map[string]interface{}{
																						"checked": c.createDetachedSignature,
																					},
																				},
																				app.Label().
																					Class("pf-c-check__label").
																					For("detached-signature-checkbox").
																					Body(
																						app.I().
																							Class("fas fa-unlink pf-u-mr-sm"),
																						app.Text("Create detached signature"),
																					),
																				app.Span().
																					Class("pf-c-check__description").
																					Text("If enabled, create a separate .asc file containing the signature."),
																			),
																	),
															),
													),
											),
										),
								),
						),
					app.Div().
						Class("pf-c-form__group").
						Aria("role", "group").
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
													ID("armor-checkbox").
													OnInput(func(ctx app.Context, e app.Event) {
														c.skipArmor = !c.skipArmor

														c.dirty = true
													}),
												Properties: map[string]interface{}{
													"checked": !c.skipArmor,
												},
											},
											app.Label().
												Class("pf-c-check__label").
												For("armor-checkbox").
												Body(
													app.I().
														Class("fas fa-shield-alt pf-u-mr-sm"),
													app.Text("Armor"),
												),
											app.Span().
												Class("pf-c-check__description").
												Text("To increase portability, ASCII armor the cyphertext."),
										),
								),
						),
				),
		},
		Footer: []app.UI{
			app.Button().
				Class("pf-c-button pf-m-primary").
				Type("submit").
				Form("encrypt-and-sign-form").
				Text(
					func() string {
						if c.skipEncryption && !c.skipSigning {
							return "Sign"
						}

						if !c.skipEncryption && c.skipSigning {
							return "Encrypt"
						}

						return "Encrypt and sign"
					}(),
				),
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

func (c *EncryptAndSignModal) clear() {
	// Clear input value
	app.Window().GetElementByID(selectEncryptionFileInput).Set("value", app.Null())

	// Clear key
	c.fileContents = []byte{}
	c.fileIsBinary = false

	c.skipEncryption = false
	c.publicKeyID = ""

	c.skipSigning = false
	c.privateKeyID = ""

	c.createDetachedSignature = false

	c.dirty = false
}

func getKeySummary(key PGPKey) string {
	return key.Label + " " + key.FullName + " <" + key.Email + ">"
}
