package components

import (
	app "github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	selectDecryptionFileInput = "select-decryption-file-input"
)

type DecryptAndVerifyModal struct {
	app.Compo

	PrivateKeys []EncryptionKey
	PublicKeys  []EncryptionKey

	OnSubmit func(
		file []byte,
		publicKeyID string,
		privateKeyID string,
		detachedSignature string,
	)
	OnCancel func()

	fileIsBinary bool
	file         []byte

	skipDecryption bool
	publicKeyID    string

	skipVerification bool
	privateKeyID     string

	useDetachedSignature bool
	detachedSignature    string
}

func (c *DecryptAndVerifyModal) Render() app.UI {
	return app.Form().
		Class("pf-c-form").
		OnSubmit(func(ctx app.Context, e app.Event) {
			e.PreventDefault()

			c.OnSubmit(
				c.file,
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
					app.Div().
						Class("pf-c-file-upload").
						Body(
							app.Div().
								Class("pf-c-file-upload__file-select").
								Body(
									app.Div().
										Class("pf-c-input-group").
										Body(
											app.Input().
												Class("pf-c-form-control").
												ID(selectDecryptionFileInput).
												Type("File").
												Aria("label", "Drag and drop a file or select one").
												ReadOnly(true).
												Placeholder("Drag and drop a file or select one").
												OnChange(func(ctx app.Context, e app.Event) {
													e.PreventDefault()

													reader := app.Window().JSValue().Get("FileReader").New()
													input := app.Window().GetElementByID(selectDecryptionFileInput)

													reader.Set("onload", app.FuncOf(func(this app.Value, args []app.Value) interface{} {
														go func() {
															rawFileContent := app.Window().Get("Uint8Array").New(args[0].Get("target").Get("result"))

															fileContent := make([]byte, rawFileContent.Get("length").Int())
															app.CopyBytesToGo(fileContent, rawFileContent)

															c.fileIsBinary = true
															c.file = fileContent
														}()

														return nil
													}))

													if file := input.Get("files").Get("0"); !file.IsUndefined() {
														reader.Call("readAsArrayBuffer", file)
													} else {
														c.clear()
													}
												}),
											app.Button().
												Class("pf-c-button pf-m-control").
												Type("button").
												Disabled(len(c.file) == 0).
												Text("Clear").
												OnClick(func(ctx app.Context, e app.Event) {
													c.clear()
												}),
										),
								),
							app.Div().
								Class("pf-c-file-upload__file-details").
								Body(
									app.Textarea().
										Class("pf-c-form-control pf-m-resize-vertical").
										ID("enter-key-input").
										Aria("label", "Enter text here").
										Placeholder("Or enter text here").
										Required(true).
										OnInput(func(ctx app.Context, e app.Event) {
											c.file = []byte(ctx.JSSrc().Get("value").String())

											if c.fileIsBinary {
												c.fileIsBinary = false

												app.Window().GetElementByID(selectDecryptionFileInput).Set("value", app.Null())
											}
										}).
										Text(func() string {
											if c.fileIsBinary {
												return "File has been selected."
											}

											return string(c.file)
										}()),
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
											ID("encryption-checkbox").
											OnInput(func(ctx app.Context, e app.Event) {
												if !(c.skipVerification && !c.skipDecryption) {
													c.skipDecryption = !c.skipDecryption
												}
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
													}).
													Body(
														app.Option().
															Value("").
															Text("Select one").
															Selected(c.privateKeyID == ""),
														app.Range(c.PrivateKeys).Slice(func(i int) app.UI {
															key := c.PrivateKeys[i]

															return app.Option().
																Value(key.ID).
																Text(getKeySummary(key)).
																Selected(c.privateKeyID == key.ID)
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
													}).
													Body(
														app.Option().
															Value("").
															Text("Select one").
															Selected(c.publicKeyID == ""),
														app.Range(c.PublicKeys).Slice(func(i int) app.UI {
															key := c.PublicKeys[i]

															return app.Option().
																Value(key.ID).
																Text(getKeySummary(key)).
																Selected(c.publicKeyID == key.ID)
														}),
													),
											),
									),
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
										Text("Decrypt and Verify"),
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

func (c *DecryptAndVerifyModal) clear() {
	// Clear input value
	app.Window().GetElementByID(selectDecryptionFileInput).Set("value", app.Null())

	// Clear key
	c.file = []byte{}
	c.fileIsBinary = false

	c.skipDecryption = false
	c.publicKeyID = ""

	c.skipVerification = false
	c.privateKeyID = ""

	c.useDetachedSignature = false
	c.detachedSignature = ""
}
