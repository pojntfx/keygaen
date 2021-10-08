package components

import (
	app "github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	selectEncryptionFileInput = "select-encryption-file-input"
)

type EncryptionKey struct {
	ID       string
	FullName string
	Email    string
}

type EncryptAndSignModal struct {
	app.Compo

	PrivateKeys []EncryptionKey
	PublicKeys  []EncryptionKey

	OnSubmit func(
		file []byte,
		publicKeyID string,
		privateKeyID string,
	)
	OnCancel func()

	fileIsBinary bool
	file         []byte

	skipEncryption bool
	publicKeyID    string

	skipSigning  bool
	privateKeyID string
}

func (c *EncryptAndSignModal) Render() app.UI {
	return app.Form().
		Class("pf-c-form").
		OnSubmit(func(ctx app.Context, e app.Event) {
			e.PreventDefault()

			c.OnSubmit(
				c.file,
				c.publicKeyID,
				c.privateKeyID,
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
												ID(selectEncryptionFileInput).
												Type("File").
												Aria("label", "Drag and drop a file or select one").
												ReadOnly(true).
												Placeholder("Drag and drop a file or select one").
												OnChange(func(ctx app.Context, e app.Event) {
													e.PreventDefault()

													reader := app.Window().JSValue().Get("FileReader").New()
													input := app.Window().GetElementByID(selectEncryptionFileInput)

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

												app.Window().GetElementByID(selectEncryptionFileInput).Set("value", app.Null())
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
												if !(c.skipSigning && !c.skipEncryption) {
													c.skipEncryption = !c.skipEncryption
												}
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
										Text(
											func() string {
												if c.skipEncryption && !c.skipSigning {
													return "Sign"
												}

												if !c.skipEncryption && c.skipSigning {
													return "Encrypt"
												}

												return "Encrypt and Sign"
											}(),
										),
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

func (c *EncryptAndSignModal) clear() {
	// Clear input value
	app.Window().GetElementByID(selectEncryptionFileInput).Set("value", app.Null())

	// Clear key
	c.file = []byte{}
	c.fileIsBinary = false

	c.skipEncryption = false
	c.publicKeyID = ""

	c.skipSigning = false
	c.privateKeyID = ""
}

func getKeySummary(key EncryptionKey) string {
	return key.ID + " " + key.FullName + " <" + key.Email + ">"
}
