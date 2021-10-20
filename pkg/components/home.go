package components

import (
	"log"

	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

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

	viewCypherAndSignatureModalOpen bool

	decryptAndVerifyPasswordModalOpen bool
	decryptAndVerifyDownloadModalOpen bool

	viewPlaintextModalOpen bool

	selectedKeyID             string
	deleteKeyConfirmModalOpen bool

	exportKeyModalOpen bool
	viewKeyModalOpen   bool

	viewPrivateKey bool

	err       error
	onRecover func()

	keys []GPGKey
}

func (c *Home) Render() app.UI {
	if c.keys == nil {
		c.keys = []GPGKey{}
	}

	privateKeyLabel := ""
	for _, candidate := range c.keys {
		if candidate.ID == c.privateKeyID {
			privateKeyLabel = candidate.Label

			break
		}
	}

	publicKeyLabel := ""
	for _, candidate := range c.keys {
		if candidate.ID == c.publicKeyID {
			publicKeyLabel = candidate.Label

			break
		}
	}

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

						c.Update()
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
					Title: `Enter password for key "` + c.privateKeyID + `"`,
					OnSubmit: func(password string) {
						c.encryptAndSignPasswordModalOpen = false
						c.encryptAndSignDownloadModalOpen = true
					},
					OnCancel: func() {
						c.confirmModalClose = func() {
							c.encryptAndSignPasswordModalOpen = false
						}
						c.confirmCloseModalOpen = true

						c.Update()
					},
				},
			),
			app.If(
				c.decryptAndVerifyPasswordModalOpen,
				&PasswordModal{
					Title: `Enter password for key "` + c.privateKeyID + `"`,
					OnSubmit: func(password string) {
						c.decryptAndVerifyPasswordModalOpen = false
						c.decryptAndVerifyDownloadModalOpen = true
					},
					OnCancel: func() {
						c.confirmModalClose = func() {
							c.decryptAndVerifyPasswordModalOpen = false
						}
						c.confirmCloseModalOpen = true

						c.Update()
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

					OnClose: func(used bool) {
						if used {
							c.encryptAndSignDownloadModalOpen = false

							c.Update()

							return
						}

						c.confirmModalClose = func() {
							c.encryptAndSignDownloadModalOpen = false
						}
						c.confirmCloseModalOpen = true

						c.Update()
					},
					OnDownload: func() {
						c.download([]byte("Hello, world"), "cypher.txt", "text/plain")

						if c.createDetachedSignature {
							c.download([]byte("asdf"), "signature.asc", "text/plain")
						}
					},
					OnView: func() {
						c.encryptAndSignDownloadModalOpen = false
						c.viewCypherAndSignatureModalOpen = true
					},
				},
			),
			app.If(
				c.decryptAndVerifyDownloadModalOpen,
				&DownloadOrViewModal{
					SubjectA:     c.privateKeyID != "",
					SubjectANoun: "file",
					SubjectAVerb: "decrypted",

					SubjectB:     c.publicKeyID != "",
					SubjectBNoun: "",
					SubjectBVerb: "verified",

					OnClose: func(used bool) {
						if used {
							c.decryptAndVerifyDownloadModalOpen = false

							c.Update()

							return
						}

						c.confirmModalClose = func() {
							c.decryptAndVerifyDownloadModalOpen = false
						}
						c.confirmCloseModalOpen = true

						c.Update()
					},
					OnDownload: func() {
						c.download([]byte("Hello, world"), "plaintext.txt", "text/plain")
					},
					OnView: func() {
						c.decryptAndVerifyDownloadModalOpen = false
						c.viewPlaintextModalOpen = true
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
								Keys: c.keys,

								OnExport: func(keyID string) {
									c.selectedKeyID = keyID
									c.exportKeyModalOpen = !c.exportKeyModalOpen

									for _, key := range c.keys {
										if key.ID == keyID {
											if key.Public {
												c.publicKeyID = key.ID
											} else {
												c.publicKeyID = ""
											}

											if key.Private {
												c.privateKeyID = key.ID
											} else {
												c.privateKeyID = ""
											}

											break
										}
									}
								},
								OnDelete: func(keyID string) {
									c.selectedKeyID = keyID
									c.deleteKeyConfirmModalOpen = !c.deleteKeyConfirmModalOpen
								},
							},
						),
				),
			app.If(
				c.createKeyModalOpen,
				&CreateKeyModal{
					OnSubmit: func(fullName, email, password string) {
						key, err := helper.GenerateKey(fullName, email, []byte(password), "x25519", 0)
						if err != nil {
							c.createKeyModalOpen = false
							c.panic(err, func() {
								c.createKeyModalOpen = true
							})

							return
						}

						c.createKeyModalOpen = false
						c.keySuccessfullyGeneratedModalOpen = true

						fingerprints, err := helper.GetSHA256Fingerprints(key)
						if err != nil {
							c.createKeyModalOpen = false
							c.panic(err, func() {
								c.createKeyModalOpen = true
							})

							return
						}

						c.keys = append(c.keys, GPGKey{
							ID:       fingerprints[0],       // Since we don't generate subkeys, we'll only have one fingerprint
							Label:    fingerprints[0][0:10], // We can safely assume that the fingerprint is at least 10 chars long
							FullName: fullName,
							Email:    email,
							Private:  true,
							Public:   true,
							Content:  key,
						})
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
					OnCancel: func(dirty bool, clear chan struct{}) {
						c.handleCancel(dirty, clear, func() {
							c.importKeyModal = false
						})
					},
				},
			),
			app.If(
				c.encryptAndSignModalOpen,
				&EncryptAndSignModal{
					Keys: c.keys,

					OnSubmit: func(file []byte, publicKeyID, privateKeyID string, createDetachedSignature bool) {
						c.publicKeyID = publicKeyID
						c.privateKeyID = privateKeyID
						c.createDetachedSignature = createDetachedSignature

						c.encryptAndSignModalOpen = false
						c.encryptAndSignPasswordModalOpen = true
					},
					OnCancel: func(dirty bool, clear chan struct{}) {
						c.handleCancel(dirty, clear, func() {
							c.encryptAndSignModalOpen = false
						})
					},
				},
			),
			app.If(
				c.decryptAndVerifyModalOpen,
				&DecryptAndVerifyModal{
					Keys: c.keys,

					OnSubmit: func(file []byte, publicKeyID, privateKeyID, detachedSignature string) {
						c.publicKeyID = publicKeyID
						c.privateKeyID = privateKeyID

						c.decryptAndVerifyModalOpen = false
						c.decryptAndVerifyPasswordModalOpen = true
					},
					OnCancel: func(dirty bool, clear chan struct{}) {
						c.handleCancel(dirty, clear, func() {
							c.decryptAndVerifyModalOpen = false
						})
					},
				},
			),
			app.If(
				c.viewCypherAndSignatureModalOpen,
				func() app.UI {
					tabs := []TextOutputModalTab{
						{
							Language: "text/plain",
							Title:    "cypher.txt",
							Body:     "Hello, world",
						},
					}
					title := "View Cypher"

					if c.createDetachedSignature {
						tabs = append(
							tabs,
							TextOutputModalTab{
								Language: "text/plain",
								Title:    "signature.asc",
								Body:     "uas-02rioj23jd",
							},
						)
						title += " and Signature"
					}

					return &TextOutputModal{
						Title: title,
						Tabs:  tabs,
						OnClose: func() {
							c.viewCypherAndSignatureModalOpen = false
							c.encryptAndSignDownloadModalOpen = true

							c.Update()
						},
					}
				}(),
			),
			app.If(
				c.viewPlaintextModalOpen,
				func() app.UI {
					return &TextOutputModal{
						Title: "View File",
						Tabs: []TextOutputModalTab{
							{
								Language: "text/plain",
								Title:    "plaintext.txt",
								Body:     "Hello, world",
							},
						},
						OnClose: func() {
							c.viewPlaintextModalOpen = false
							c.decryptAndVerifyDownloadModalOpen = true

							c.Update()
						},
					}
				}(),
			),
			app.If(
				c.exportKeyModalOpen,
				&ExportKeyModal{
					PublicKey: c.publicKeyID != "",
					OnDownloadPublicKey: func(armor bool) {
						if armor {
							c.download([]byte("asdfirj230sd"), publicKeyLabel+".pub", "application/octet-stream")
						} else {
							c.download([]byte("asdfirj230sd"), publicKeyLabel+".pub", "text/plain")
						}
					},
					OnViewPublicKey: func() {
						c.exportKeyModalOpen = false
						c.viewPrivateKey = false
						c.viewKeyModalOpen = true
					},

					PrivateKey: c.privateKeyID != "",
					OnDownloadPrivateKey: func(armor bool) {
						if armor {
							c.download([]byte("i34jisdhjs"), privateKeyLabel, "application/octet-stream")
						} else {
							c.download([]byte("i34jisdhjs"), privateKeyLabel, "text/plain")
						}
					},
					OnViewPrivateKey: func() {
						c.exportKeyModalOpen = false
						c.viewPrivateKey = true
						c.viewKeyModalOpen = true
					},

					OnOK: func() {
						c.exportKeyModalOpen = false

						c.Update()
					},
				},
			),
			app.If(
				c.viewKeyModalOpen,
				func() app.UI {
					tabs := []TextOutputModalTab{
						{
							Language: "text/plain",
							Title:    publicKeyLabel + ".pub",
							Body:     "asdfirj230sd",
						},
					}
					title := `View Public Key "` + publicKeyLabel + `"`

					if c.viewPrivateKey {
						tabs = []TextOutputModalTab{
							{
								Language: "text/plain",
								Title:    privateKeyLabel,
								Body:     "i34jisdhjs",
							},
						}
						title = `View Private Key "` + privateKeyLabel + `"`
					}

					return &TextOutputModal{
						Title: title,
						Tabs:  tabs,
						OnClose: func() {
							c.viewKeyModalOpen = false
							c.exportKeyModalOpen = true

							c.Update()
						},
					}
				}(),
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
			app.If(
				c.deleteKeyConfirmModalOpen,
				&ConfirmationModal{
					ID:    "delete-key-confirmation-modal",
					Icon:  "fas fa-exclamation-triangle",
					Title: "Are you sure?",
					Class: "pf-m-danger",
					Body:  "After deletion, you will not be able to restore the key.",

					ActionLabel: `Yes, delete key "` + c.selectedKeyID + `"`,
					ActionClass: "pf-m-danger",

					CancelLabel: "Cancel",

					OnClose: func() {
						c.deleteKeyConfirmModalOpen = false

						c.Update()
					},
					OnAction: func() {
						c.deleteKeyConfirmModalOpen = false

						c.Update()
					},
				},
			),
			app.If(
				c.err != nil,
				&ErrorModal{
					ID:          "error-modal",
					Icon:        "fas fa-times",
					Title:       "An Error Occurred",
					Class:       "pf-m-danger",
					Body:        "The following details may be of help:",
					Error:       c.err,
					ActionLabel: "Close",

					OnClose: func() {
						c.recover()
					},
					OnAction: func() {
						c.recover()
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
		confirm()

		clear <- struct{}{}

		c.Update()

		return
	}

	c.confirmModalClose = func() {
		confirm()

		clear <- struct{}{}
	}
	c.confirmCloseModalOpen = true

	c.Update()
}

func (c *Home) panic(err error, onRecover func()) {
	log.Println(err)

	c.onRecover = onRecover
	c.err = err
}

func (c *Home) recover() {
	c.err = nil
	c.onRecover()
}

func (c *Home) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}
