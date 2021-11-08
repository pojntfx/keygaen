package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	selectDecryptionFileInput        = "select-decryption-file-input"
	selectDetachedSignatureFileInput = "select-detached-signature-file-input"
)

type DecryptAndVerifyModal struct {
	app.Compo

	Keys []GPGKey

	OnSubmit func(
		file []byte,
		publicKeyID string,
		privateKeyID string,
		detachedSignature []byte,
	)
	OnCancel func(dirty bool, clear chan struct{})

	fileContents []byte

	skipDecryption bool
	publicKeyID    string

	skipVerification bool
	privateKeyID     string

	useDetachedSignature bool
	detachedSignature    []byte

	dirty bool
}

func (c *DecryptAndVerifyModal) Render() app.UI {
	privateKeys := []GPGKey{}
	publicKeys := []GPGKey{}
	for _, key := range c.Keys {
		if key.Private {
			privateKeys = append(privateKeys, key)
		}

		if key.Public {
			publicKeys = append(publicKeys, key)
		}
	}

	return &Modal{
		ID:    "decrypt-and-verify-modal",
		Title: "Decrypt/Verify",
		Body: []app.UI{
			app.Form().
				Class("pf-c-form").
				ID("decrypt-and-verify-form").
				OnSubmit(func(ctx app.Context, e app.Event) {
					e.PreventDefault()

					c.OnSubmit(
						c.fileContents,
						c.publicKeyID,
						c.privateKeyID,
						c.detachedSignature,
					)

					c.clear()
				}).
				Body(
					app.Div().
						Class("pf-c-form__group").
						Body(
							&FileUpload{
								ID:                    selectDecryptionFileInput,
								FileSelectionLabel:    "Drag and drop a file or select one",
								ClearLabel:            "Clear",
								TextEntryLabel:        "Or enter text here",
								TextEntryBlockedLabel: "File has been selected.",
								FileContents:          c.fileContents,

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
														if !(c.skipVerification && !c.skipDecryption) {
															c.skipDecryption = !c.skipDecryption
														}

														if c.skipDecryption {
															c.privateKeyID = ""
														}

														c.dirty = true
													}),
												Properties: map[string]interface{}{
													"checked": !c.skipDecryption,
												},
											},
											app.Label().
												Class("pf-c-check__label").
												For("encryption-checkbox").
												Body(
													app.I().
														Class("fas fa-unlock pf-u-mr-sm"),
													app.Text("Decrypt file"),
												),
											app.If(
												c.skipDecryption,
												app.Span().
													Class("pf-c-check__description").
													Text("If enabled, decrypt the file using the select key."),
											).Else(
												app.Span().
													Class("pf-c-check__description").
													Text("Decrypt the file using the following key:"),
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
														if !(!c.skipVerification && c.skipDecryption) {
															c.skipVerification = !c.skipVerification
														}

														if c.skipVerification {
															c.publicKeyID = ""
														}

														c.dirty = true
													}),
												Properties: map[string]interface{}{
													"checked": !c.skipVerification,
												},
											},
											app.Label().
												Class("pf-c-check__label").
												For("signature-checkbox").
												Body(
													app.I().
														Class("fas fa-user-check pf-u-mr-sm"),
													app.Text("Verify file"),
												),
											app.If(
												c.skipVerification,
												app.Span().
													Class("pf-c-check__description").
													Text("If enabled, verify that the file originates from the person with the selected key."),
											).Else(
												app.Span().
													Class("pf-c-check__description").
													Text("Verify that the file originates from the person with the following key:"),
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
																							c.useDetachedSignature = !c.useDetachedSignature

																							c.dirty = true
																						}),
																					Properties: map[string]interface{}{
																						"checked": c.useDetachedSignature,
																					},
																				},
																				app.Label().
																					Class("pf-c-check__label").
																					For("detached-signature-checkbox").
																					Body(
																						app.I().
																							Class("fas fa-unlink pf-u-mr-sm"),
																						app.Text("Use detached signature"),
																					),
																				app.If(
																					!c.useDetachedSignature,
																					app.Span().
																						Class("pf-c-check__description").
																						Text("If enabled, validate the file using a detached signature (.asc file)."),
																				).Else(
																					app.Span().
																						Class("pf-c-check__description").
																						Text("Validate the file using the following detached signature (.asc file):"),
																					app.Div().
																						Class("pf-c-check__body pf-u-w-100").
																						Body(
																							&FileUpload{
																								ID:                    selectDetachedSignatureFileInput,
																								FileSelectionLabel:    "Drag and drop a detached signature or select one",
																								ClearLabel:            "Clear",
																								TextEntryLabel:        "Or enter the detached signature's content here",
																								TextEntryBlockedLabel: "File has been selected.",
																								FileContents:          []byte(c.detachedSignature),

																								OnChange: func(fileContents []byte) {
																									c.detachedSignature = fileContents

																									c.dirty = true
																								},
																								OnClear: func() {
																									c.detachedSignature = []byte{}
																								},
																							},
																						),
																				),
																			),
																	),
															),
													),
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
				Form("decrypt-and-verify-form").
				Text(
					func() string {
						if c.skipDecryption && !c.skipVerification {
							return "Verify"
						}

						if !c.skipDecryption && c.skipVerification {
							return "Decrypt"
						}

						return "Decrypt and verify"
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

func (c *DecryptAndVerifyModal) clear() {
	// Clear file input values
	app.Window().GetElementByID(selectDecryptionFileInput).Set("value", app.Null())
	if c.useDetachedSignature {
		app.Window().GetElementByID(selectDetachedSignatureFileInput).Set("value", app.Null())
	}

	// Clear key
	c.fileContents = []byte{}

	c.skipDecryption = false
	c.publicKeyID = ""

	c.skipVerification = false
	c.privateKeyID = ""

	c.useDetachedSignature = false
	c.detachedSignature = []byte{}

	c.dirty = false
}
