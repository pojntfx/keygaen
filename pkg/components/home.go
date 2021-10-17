package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var demoKeys = []GPGKey{
	{
		ID:       "039292",
		FullName: "Isegard Example",
		Email:    "isegard@example.com",
		Private:  true,
	},
	{
		ID:       "838431",
		FullName: "Fred Example",
		Email:    "fred@example.com",
		Private:  true,
		Public:   true,
	},
	{
		ID:       "123456",
		FullName: "Alice Example",
		Email:    "alice@example.com",
		Public:   true,
	},
	{
		ID:       "319312",
		FullName: "Bob Example",
		Email:    "bob@example.com",
		Public:   true,
	},
}

type Home struct {
	app.Compo

	createKeyModalOpen                bool
	importKeyModal                    bool
	encryptAndSignModalOpen           bool
	decryptAndVerifyModalOpen         bool
	keySuccessfullyGeneratedModalOpen bool

	keyImportPasswordModalOpen       bool
	keySuccessfullyImportedModalOpen bool

	publicKeyID             string
	privateKeyID            string
	createDetachedSignature bool

	encryptAndSignPasswordModalOpen bool
	encryptAndSignDownloadModalOpen bool

	confirmCloseModalOpen bool
	confirmModalClose     func()
}

func (c *Home) Render() app.UI {
	return app.Div().
		Class("pf-c-page").
		Body(
			app.A().
				Class("pf-c-skip-to-content pf-c-button pf-m-primary").
				Href("#gridge-main").
				Body(
					app.Text("Skip to content"),
				),
			&Navbar{},
			app.If(
				c.keySuccessfullyGeneratedModalOpen,
				&SuccessModal{
					ID:          "key-successfully-generated-modal",
					Icon:        "fas fa-check",
					Title:       "Key Successfully Generated!",
					Class:       "pf-m-success",
					Body:        "It has been added to the key list.",
					ActionLabel: "Continue to key list",

					OnClose: func() {
						c.keySuccessfullyGeneratedModalOpen = false

						c.Update()
					},
					OnAction: func() {
						c.keySuccessfullyGeneratedModalOpen = false
					},
				},
			),
			app.If(
				c.keyImportPasswordModalOpen,
				&PasswordModal{
					Title: "Enter Key Password",
					OnSubmit: func(password string) {
						c.keyImportPasswordModalOpen = false
						c.keySuccessfullyImportedModalOpen = true
					},
					OnCancel: func() {
						c.confirmModalClose = func() {
							c.keyImportPasswordModalOpen = false
						}
						c.confirmCloseModalOpen = true
					},
				},
			),
			app.If(
				c.keySuccessfullyImportedModalOpen,
				&SuccessModal{
					ID:          "key-successfully-imported-modal",
					Icon:        "fas fa-check",
					Title:       "Key Successfully Imported!",
					Class:       "pf-m-success",
					Body:        "It has been added to the key list.",
					ActionLabel: "Continue to key list",

					OnClose: func() {
						c.keySuccessfullyImportedModalOpen = false

						c.Update()
					},
					OnAction: func() {
						c.keySuccessfullyImportedModalOpen = false
					},
				},
			),
			app.If(
				c.encryptAndSignPasswordModalOpen,
				&PasswordModal{
					Title: "Enter Key Password",
					OnSubmit: func(password string) {
						c.encryptAndSignPasswordModalOpen = false
						c.encryptAndSignDownloadModalOpen = true
					},
					OnCancel: func() {
						c.confirmModalClose = func() {
							c.encryptAndSignPasswordModalOpen = false
						}
						c.confirmCloseModalOpen = true
					},
				},
			),
			app.If(
				c.encryptAndSignDownloadModalOpen,
				&DownloadOrViewModal{
					SubjectA:     c.privateKeyID != "",
					SubjectANoun: "signature",
					SubjectAVerb: "signed",

					SubjectB:     c.publicKeyID != "",
					SubjectBNoun: "cypher",
					SubjectBVerb: "encrypted",

					OnClose: func() {
						c.confirmModalClose = func() {
							c.encryptAndSignDownloadModalOpen = false
						}
						c.confirmCloseModalOpen = true
					},
					OnDownload: func() {
						c.download([]byte("Hello, world"), "cypher.txt", "text/plain")

						if c.createDetachedSignature {
							c.download([]byte("asdf"), "signature.asc", "text/plain")
						}
					},
					OnView: func() {
						app.Window().Call("alert", "Successfully viewed")
					},
				},
			),
			app.Main().
				Class("pf-c-page__main").
				ID("gridge-main").
				TabIndex(-1).
				Body(
					app.Section().
						Class("pf-c-page__main-section pf-m-light pf-m-no-padding pf-u-px-sm-on-xl").
						Body(
							&Toolbar{
								OnCreateKey: func() {
									c.createKeyModalOpen = !c.createKeyModalOpen
								},
								OnImportKey: func() {
									c.importKeyModal = !c.importKeyModal
								},

								OnEncryptAndSign: func() {
									c.encryptAndSignModalOpen = !c.encryptAndSignModalOpen
								},
								OnDecryptAndVerify: func() {
									c.decryptAndVerifyModalOpen = !c.decryptAndVerifyModalOpen
								},
							},
						),
					app.Section().
						Class("pf-c-page__main-section pf-m-fill").
						Body(
							&KeyList{
								Keys: demoKeys,
							},
						),
				),
			app.If(
				c.createKeyModalOpen,
				&CreateKeyModal{
					OnSubmit: func(fullName, email, _ string) {
						c.createKeyModalOpen = false
						c.keySuccessfullyGeneratedModalOpen = true
					},
					OnCancel: func(dirty bool, clear chan struct{}) {
						c.handleCancel(dirty, clear, func() {
							c.createKeyModalOpen = false
						})
					},
				},
			),
			app.If(
				c.importKeyModal,
				&ImportKeyModal{
					OnSubmit: func(key string) {
						c.importKeyModal = false
						c.keyImportPasswordModalOpen = true
					},
					OnCancel: func() {
						c.confirmModalClose = func() {
							c.importKeyModal = false
						}
						c.confirmCloseModalOpen = true
					},
				},
			),
			app.If(
				c.encryptAndSignModalOpen,
				&EncryptAndSignModal{
					Keys: demoKeys,

					OnSubmit: func(file []byte, publicKeyID, privateKeyID string, createDetachedSignature bool) {
						c.publicKeyID = publicKeyID
						c.privateKeyID = privateKeyID
						c.createDetachedSignature = true

						c.encryptAndSignModalOpen = false
						c.encryptAndSignPasswordModalOpen = true
					},
					OnCancel: func() {
						c.confirmModalClose = func() {
							c.encryptAndSignModalOpen = false
						}
						c.confirmCloseModalOpen = true
					},
				},
			),
			app.If(
				c.decryptAndVerifyModalOpen,
				&DecryptAndVerifyModal{
					Keys: demoKeys,

					OnSubmit: func(file []byte, publicKeyID, privateKeyID, detachedSignature string) {
						app.Window().Call("alert", fmt.Sprintf("Decrypted and verified file %v, using public key ID %v, private key ID %v and detached signature %v", file, publicKeyID, privateKeyID, detachedSignature))

						c.decryptAndVerifyModalOpen = false
					},
					OnCancel: func() {
						c.confirmModalClose = func() {
							c.decryptAndVerifyModalOpen = false
						}
						c.confirmCloseModalOpen = true
					},
				},
			),
			app.If(
				c.confirmCloseModalOpen,
				&ConfirmationModal{
					ID:    "confirmation-modal",
					Icon:  "fas fa-exclamation-triangle",
					Title: "Are you sure?",
					Class: "pf-m-danger",
					Body:  "Unsaved changes might be lost.",

					ActionLabel: "Yes, delete unsaved changes",
					ActionClass: "pf-m-danger",

					CancelLabel: "Cancel",

					OnClose: func() {
						c.confirmCloseModalOpen = false

						c.Update()
					},
					OnAction: func() {
						c.confirmCloseModalOpen = false

						c.confirmModalClose()

						c.Update()
					},
				},
			),
		)
}

func (c *Home) download(content []byte, name string, mimetype string) {
	buf := app.Window().JSValue().Get("Uint8Array").New(len(content))
	app.CopyBytesToJS(buf, content)

	blob := app.Window().JSValue().Get("Blob").New(app.Window().JSValue().Get("Array").New(buf), map[string]interface{}{
		"type": mimetype,
	})

	link := app.Window().Get("document").Call("createElement", "a")
	link.Set("href", app.Window().Get("URL").Call("createObjectURL", blob))
	link.Set("download", name)
	link.Call("click")
}

func (c *Home) handleCancel(dirty bool, clear chan struct{}, confirm func()) {
	if !dirty {
		c.createKeyModalOpen = false

		clear <- struct{}{}

		c.Update()

		return
	}

	c.confirmModalClose = func() {
		confirm()

		clear <- struct{}{}
	}
	c.confirmCloseModalOpen = true
}

func (c *Home) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}
