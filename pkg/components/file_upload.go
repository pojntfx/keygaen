package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// FileUpload provides a way to upload/import a file by entering it's contents or selecting it
type FileUpload struct {
	app.Compo

	ID                         string // HTML ID of the modal; must be unique across the page
	FileSelectionLabel         string // Text to display on the file selection input
	ClearLabel                 string // Text to display on the clear action
	TextEntryInputPlaceholder  string // Placeholder to display in the text input field
	TextEntryInputBlockedLabel string // Text to display in the input if it is blocked (i.e. because a file has been selected)
	FileContents               []byte // Contents of the file selected/text entered

	OnChange func(fileContents []byte) // Handler to call to set `FileContents`
	OnClear  func()                    // Handler to call to clear `FileContents`

	fileIsBinary bool
}

func (c *FileUpload) Render() app.UI {
	return app.Div().
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
								ID(c.ID).
								Type("File").
								Aria("label", c.FileSelectionLabel).
								ReadOnly(true).
								Placeholder(c.FileSelectionLabel).
								OnChange(func(ctx app.Context, e app.Event) {
									e.PreventDefault()

									reader := app.Window().JSValue().Get("FileReader").New()
									input := app.Window().GetElementByID(c.ID)

									reader.Set("onload", app.FuncOf(func(this app.Value, args []app.Value) interface{} {
										go func() {
											rawFileContent := app.Window().Get("Uint8Array").New(args[0].Get("target").Get("result"))

											fileContent := make([]byte, rawFileContent.Get("length").Int())
											app.CopyBytesToGo(fileContent, rawFileContent)

											c.fileIsBinary = true
											c.OnChange(fileContent)
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
								Disabled(len(c.FileContents) == 0).
								Text(c.ClearLabel).
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
						Aria("label", c.TextEntryInputPlaceholder).
						Placeholder(c.TextEntryInputPlaceholder).
						Required(true).
						OnInput(func(ctx app.Context, e app.Event) {
							c.OnChange([]byte(ctx.JSSrc().Get("value").String()))

							if c.fileIsBinary {
								c.fileIsBinary = false

								app.Window().GetElementByID(c.ID).Set("value", app.Null())
							}
						}).
						Text(func() string {
							if c.fileIsBinary {
								return c.TextEntryInputBlockedLabel
							}

							return string(c.FileContents)
						}()),
				),
		)
}

func (c *FileUpload) clear() {
	// Clear input value
	app.Window().GetElementByID(c.ID).Set("value", app.Null())

	// Clear key
	c.fileIsBinary = false

	c.OnClear()
}
