package components

import (
	app "github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	selectFileInput = "select-file-input"
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

	file         []byte
	isBinary     bool
	publicKeyID  string
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
												ID(selectFileInput).
												Type("File").
												Aria("label", "Drag and drop a file or select one").
												ReadOnly(true).
												Placeholder("Drag and drop a file or select one").
												OnChange(func(ctx app.Context, e app.Event) {
													e.PreventDefault()

													reader := app.Window().JSValue().Get("FileReader").New()
													input := app.Window().GetElementByID(selectFileInput)

													reader.Set("onload", app.FuncOf(func(this app.Value, args []app.Value) interface{} {
														go func() {
															rawFileContent := app.Window().Get("Uint8Array").New(args[0].Get("target").Get("result"))

															fileContent := make([]byte, rawFileContent.Get("length").Int())
															app.CopyBytesToGo(fileContent, rawFileContent)

															c.isBinary = true
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

											if c.isBinary {
												c.isBinary = false

												app.Window().GetElementByID(selectFileInput).Set("value", app.Null())
											}
										}).
										Text(func() string {
											if c.isBinary {
												return "File has been selected."
											}

											return string(c.file)
										}()),
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
										Text("Encrypt and Sign"),
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
	app.Window().GetElementByID(selectFileInput).Set("value", app.Null())

	// Clear key
	c.file = []byte{}
	c.isBinary = false
}
