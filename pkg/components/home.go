package components

import (
	"errors"
	"log"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/gridge/pkg/crypt"
)

type Home struct {
	app.Compo

	createKeyModalOpen                bool
	importKeyModalOpen                bool
	encryptAndSignModalOpen           bool
	decryptAndVerifyModalOpen         bool
	keySuccessfullyGeneratedModalOpen bool

	keyPasswordModalOpen             bool
	keyPasswordModalKeyID            string
	keySuccessfullyImportedModalOpen bool

	publicKeyID             string
	privateKeyID            string
	createDetachedSignature bool

	encryptAndSignDownloadModalOpen bool

	confirmCloseModalOpen bool
	confirmModalClose     func()

	viewCypherAndSignatureModalOpen bool

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

	keyPasswordChan chan string

	wrongPassword bool

	keyDuplicateModalOpen bool
}

func (c *Home) Render() app.UI {
	if c.keys == nil {
		c.keys = []GPGKey{}
	}

	if c.keyPasswordChan == nil {
		c.keyPasswordChan = make(chan string)
	}

	privateKey := GPGKey{}
	privateKeyExport := []byte{}
	privateKeyExportArmored := ""
	for _, candidate := range c.keys {
		if candidate.ID == c.privateKeyID {
			privateKey = candidate

			parsedKey, err := crypto.NewKeyFromArmored(string(privateKey.Content)) // TODO: Use crypt package's implementation to support raw keys
			if err != nil {
				c.panic(err, func() {})

				break
			}

			privateKeyExport, err = parsedKey.Serialize()
			if err != nil {
				c.panic(err, func() {})

				break
			}

			privateKeyExportArmored, err = parsedKey.Armor()
			if err != nil {
				c.panic(err, func() {})

				break
			}

			break
		}
	}

	publicKey := GPGKey{}
	publicKeyExport := []byte{}
	publicKeyExportArmored := ""
	for _, candidate := range c.keys {
		if candidate.ID == c.publicKeyID {
			publicKey = candidate

			parsedKey, err := crypto.NewKeyFromArmored(string(publicKey.Content)) // TODO: Use crypt package's implementation to support raw keys
			if err != nil {
				c.panic(err, func() {})

				break
			}

			publicKeyExport, err = parsedKey.GetPublicKey()
			if err != nil {
				c.panic(err, func() {})

				break
			}

			publicKeyExportArmored, err = parsedKey.GetArmoredPublicKey()
			if err != nil {
				c.panic(err, func() {})

				break
			}

			break
		}
	}

	selectedKey := GPGKey{}
	for _, candidate := range c.keys {
		if candidate.ID == c.selectedKeyID {
			selectedKey = candidate

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
				&SingleActionModal{
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
				c.keyPasswordModalOpen,
				&PasswordModal{
					Title:         `Enter Password for Key "` + c.keyPasswordModalKeyID + `"`,
					WrongPassword: c.wrongPassword,
					ClearWrongPassword: func() {
						c.wrongPassword = false
					},
					OnSubmit: func(password string) {
						c.keyPasswordChan <- password
					},
					OnCancel: func() {
						c.confirmModalClose = func() {
							c.keyPasswordChan <- ""
						}
						c.confirmCloseModalOpen = true

						c.Update()
					},
				},
			),
			app.If(
				c.keySuccessfullyImportedModalOpen,
				&SingleActionModal{
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
				c.keyDuplicateModalOpen,
				&SingleActionModal{
					ID:          "key-successfully-imported-modal",
					Icon:        "fas fa-info-circle",
					Title:       "Key Already Exists",
					Class:       "pf-m-info",
					Body:        "This key is already in the key list.",
					ActionLabel: "Continue to key list",

					OnClose: func() {
						c.keyDuplicateModalOpen = false

						c.Update()
					},
					OnAction: func() {
						c.keyDuplicateModalOpen = false
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
						c.download([]byte("Hello, world"), "cypher.txt", "text/plain") // TODO: Use `.gpg` if unarmored, `.asc` if armored

						if c.createDetachedSignature {
							c.download([]byte("asdf"), "signature.asc", "text/plain") // TODO: Use `.gpg` if unarmored, `.asc` if armored
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
						c.download([]byte("Hello, world"), "plaintext.txt", "text/plain") // TODO: Use `.gpg` if unarmored, `.asc` if armored
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
					app.If(
						len(c.keys) != 0,
						app.Section().
							Class("pf-c-page__main-section pf-m-light pf-m-no-padding pf-u-px-sm-on-xl").
							Body(
								&Toolbar{
									OnCreateKey: func() {
										c.createKeyModalOpen = !c.createKeyModalOpen
									},
									OnImportKey: func() {
										c.importKeyModalOpen = !c.importKeyModalOpen
									},

									OnEncryptAndSign: func() {
										c.encryptAndSignModalOpen = !c.encryptAndSignModalOpen
									},
									OnDecryptAndVerify: func() {
										c.decryptAndVerifyModalOpen = !c.decryptAndVerifyModalOpen
									},
								},
							),
					),
					app.Section().
						Class("pf-c-page__main-section pf-m-fill").
						Body(
							app.If(
								len(c.keys) == 0,
								&EmptyState{
									OnCreateKey: func() {
										c.createKeyModalOpen = !c.createKeyModalOpen
									},
									OnImportKey: func() {
										c.importKeyModalOpen = !c.importKeyModalOpen
									},
								},
							).Else(
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

						parsedKey, fingerprint, err := crypt.ReadKey([]byte(key), password)
						if err != nil {
							c.panic(err, func() {})

							return
						}

						id := parsedKey.PrimaryIdentity()
						if id == nil {
							c.panic(errors.New("no identity found in key"), func() {})

							return
						}

						c.keys = append(c.keys, GPGKey{
							ID:       fingerprint,       // Since we don't generate subkeys, we'll only have one fingerprint
							Label:    fingerprint[0:10], // We can safely assume that the fingerprint is at least 10 chars long
							FullName: id.Name,
							Email:    id.UserId.Email,
							Private:  parsedKey.PrivateKey != nil,
							Public:   true,
							Content:  []byte(key),
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
				c.importKeyModalOpen,
				&ImportKeyModal{
					OnSubmit: func(key []byte) {
						c.importKeyModalOpen = false

						go func() {
							var parsedKey *openpgp.Entity

							// We might have to unlock a private key first
							locked, fingerprint, err := crypt.IsKeyLocked(key)
							if err != nil {
								c.panic(err, func() {})

								return
							}

							if locked {
								password := c.getPasswordForKey(
									fingerprint,
									func(password string) error {
										pk, fp, err := crypt.ReadKey(key, password)
										if err == nil {
											parsedKey = pk
											fingerprint = fp
										}

										return err
									},
								)

								// Stop import if no password has been entered
								if password == "" {
									c.keyPasswordModalOpen = false

									return
								}
							}

							if parsedKey == nil {
								parsedKey, fingerprint, err = crypt.ReadKey(key, "")
								if err != nil {
									c.panic(err, func() {})

									return
								}
							}

							id := parsedKey.PrimaryIdentity()
							if id == nil {
								c.panic(errors.New("no identity found in key"), func() {})

								return
							}

							newKeys := []GPGKey{}
							for _, candidate := range c.keys {
								if candidate.ID == fingerprint {
									// Replace the key if the existing key is a public key and the imported key is a private key
									if !candidate.Private && parsedKey.PrivateKey != nil {
										continue
									} else {
										// Don't add the duplicate key
										c.keyDuplicateModalOpen = true

										c.Update()

										return
									}
								}

								newKeys = append(newKeys, candidate)
							}

							newKeys = append(newKeys, GPGKey{
								ID:       fingerprint,       // Since we don't generate subkeys, we'll only have one fingerprint
								Label:    fingerprint[0:10], // We can safely assume that the fingerprint is at least 10 chars long
								FullName: id.Name,
								Email:    id.UserId.Email,
								Private:  parsedKey.PrivateKey != nil,
								Public:   true,
								Content:  key,
							})
							c.keys = newKeys

							c.keySuccessfullyImportedModalOpen = true

							c.Update()
						}()
					},
					OnCancel: func(dirty bool, clear chan struct{}) {
						c.handleCancel(dirty, clear, func() {
							c.importKeyModalOpen = false
						})
					},
				},
			),
			app.If(
				c.encryptAndSignModalOpen,
				&EncryptAndSignModal{
					Keys: c.keys,

					OnSubmit: func(file []byte, publicKeyID, privateKeyID string, createDetachedSignature, enableArmor bool) {
						c.publicKeyID = publicKeyID
						c.privateKeyID = privateKeyID
						c.createDetachedSignature = createDetachedSignature

						c.encryptAndSignModalOpen = false

						go func() {
							// Encrypt
							var encryptConfig *crypt.EncryptConfig
							if c.publicKeyID != "" {
								rawPublicKey, err := c.getPublicKeyByID(c.publicKeyID)
								if err != nil {
									c.panic(err, func() {})

									return
								}

								publicKey, _, err := crypt.ReadKey([]byte(rawPublicKey), "")
								if err != nil {
									c.panic(err, func() {})

									return
								}

								encryptConfig = &crypt.EncryptConfig{
									PublicKey:       publicKey,
									ArmorCyphertext: enableArmor,
								}
							}

							// Sign
							var signatureConfig *crypt.SignatureConfig
							if c.privateKeyID != "" {
								rawPrivateKey, err := c.getPrivateKeyByID(c.privateKeyID)
								if err != nil {
									c.panic(err, func() {})

									return
								}

								var privateKey *openpgp.Entity

								// We might have to unlock a private key first
								locked, fingerprint, err := crypt.IsKeyLocked([]byte(rawPrivateKey))
								if err != nil {
									c.panic(err, func() {})

									return
								}

								if locked {
									password := c.getPasswordForKey(
										fingerprint,
										func(password string) error {
											pk, fp, err := crypt.ReadKey([]byte(rawPrivateKey), password)
											if err == nil {
												privateKey = pk
												fingerprint = fp
											}

											return err
										},
									)

									// Stop import if no password has been entered
									if password == "" {
										c.keyPasswordModalOpen = false

										return
									}
								}

								if privateKey == nil {
									privateKey, fingerprint, err = crypt.ReadKey([]byte(rawPrivateKey), "")
									if err != nil {
										c.panic(err, func() {})

										return
									}
								}

								signatureConfig = &crypt.SignatureConfig{
									PrivateKey:      privateKey,
									ArmorSignature:  enableArmor,
									DetachSignature: createDetachedSignature,
								}
							}

							cyphertext, signature, err := crypt.EncryptSign(
								encryptConfig,
								signatureConfig,
								file,
							)
							if err != nil {
								c.panic(err, func() {})

								return
							}

							if enableArmor {
								log.Printf("Cyphertext: %s\n", cyphertext)
								log.Printf("Signature: %s\n", signature)
							} else {
								log.Print(cyphertext)
								log.Print(signature)
							}

							c.encryptAndSignDownloadModalOpen = true

							c.Update()
						}()
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

					OnSubmit: func(file []byte, publicKeyID, privateKeyID string, detachedSignature []byte) {
						c.publicKeyID = publicKeyID
						c.privateKeyID = privateKeyID

						c.decryptAndVerifyModalOpen = false

						go func() {
							// Decrypt
							if c.privateKeyID != "" && c.publicKeyID == "" {
								rawPrivateKey, err := c.getPrivateKeyByID(c.privateKeyID)
								if err != nil {
									c.panic(err, func() {})

									return
								}

								var privateKey *openpgp.Entity

								// We might have to unlock a private key first
								locked, fingerprint, err := crypt.IsKeyLocked([]byte(rawPrivateKey))
								if err != nil {
									c.panic(err, func() {})

									return
								}

								if locked {
									password := c.getPasswordForKey(
										fingerprint,
										func(password string) error {
											pk, fp, err := crypt.ReadKey([]byte(rawPrivateKey), password)
											if err == nil {
												privateKey = pk
												fingerprint = fp
											}

											return err
										},
									)

									// Stop import if no password has been entered
									if password == "" {
										c.keyPasswordModalOpen = false

										return
									}
								}

								if privateKey == nil {
									privateKey, fingerprint, err = crypt.ReadKey([]byte(rawPrivateKey), "")
									if err != nil {
										c.panic(err, func() {})

										return
									}
								}

								plaintext, _, err := crypt.DecryptVerify(
									&crypt.DecryptConfig{
										PrivateKey: privateKey,
									},
									nil,
									file,
								)
								if err != nil {
									c.panic(err, func() {})

									return
								}

								log.Printf("Plaintext: %s\n", plaintext)

								c.decryptAndVerifyDownloadModalOpen = true

								c.Update()

								return
							}

							// Verify
							if c.publicKeyID != "" && c.privateKeyID == "" {
								rawPublicKey, err := c.getPublicKeyByID(c.publicKeyID)
								if err != nil {
									c.panic(err, func() {})

									return
								}

								publicKey, _, err := crypt.ReadKey([]byte(rawPublicKey), "")
								if err != nil {
									c.panic(err, func() {})

									return
								}

								plaintext, verified, err := crypt.DecryptVerify(
									nil,
									&crypt.VerifyConfig{
										PublicKey:         publicKey,
										DetachedSignature: detachedSignature,
									},
									file,
								)
								if err != nil {
									c.panic(err, func() {})

									return
								}

								log.Printf("Plaintext: %s\n", plaintext)
								log.Printf("Verified: %v\n", verified)

								c.decryptAndVerifyDownloadModalOpen = true

								c.Update()

								return
							}
						}()
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
							Title:    "cypher.txt", // TODO: Use `.gpg` if unarmored, `.asc` if armored
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
								Title:    "plaintext.txt", // TODO: Use `.gpg` if unarmored, `.asc` if armored
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
							c.download([]byte(publicKeyExportArmored), publicKey.Label+".asc", "text/plain")
						} else {
							c.download(publicKeyExport, publicKey.Label+".gpg", "application/octet-stream")
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
							c.download([]byte(privateKeyExportArmored), privateKey.Label+".asc", "text/plain")
						} else {
							c.download(privateKeyExport, privateKey.Label+".gpg", "application/octet-stream")
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
							Title:    publicKey.Label + ".pub",
							Body:     publicKeyExportArmored,
						},
					}
					title := `View Public Key "` + publicKey.Label + `"`

					if c.viewPrivateKey {
						tabs = []TextOutputModalTab{
							{
								Language: "text/plain",
								Title:    privateKey.Label,
								Body:     privateKeyExportArmored,
							},
						}
						title = `View Private Key "` + privateKey.Label + `"`
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

					ActionLabel: `Yes, delete key "` + selectedKey.Label + `"`,
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

		// Remove the native confirmation prompt
		app.Window().Set("onbeforeunload", app.Undefined())

		clear <- struct{}{}
	}

	// Add the native confirmation prompt
	app.Window().Set(
		"onbeforeunload",
		app.FuncOf(func(this app.Value, args []app.Value) interface{} {
			args[0].Set("returnValue", "")

			return nil
		}),
	)
	c.confirmCloseModalOpen = true

	c.Update()
}

func (c *Home) panic(err error, onRecover func()) {
	log.Println(err)

	c.onRecover = onRecover
	c.err = err

	c.Update()
}

func (c *Home) recover() {
	c.err = nil
	c.onRecover()
}

func (c *Home) handleWrongPassword(err error) {
	log.Println(err)

	c.wrongPassword = true

	c.Update()
}

func (c *Home) clearWrongPassword() {
	c.wrongPassword = false
}

func (c *Home) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}

func (c *Home) getPublicKeyByID(ID string) (string, error) {
	publicKeyExportArmored := ""
	for _, candidate := range c.keys {
		if candidate.ID == c.publicKeyID {
			publicKey := candidate

			parsedKey, err := crypto.NewKeyFromArmored(string(publicKey.Content)) // TODO: Use crypt package's implementation to support raw keys
			if err != nil {
				return "", err
			}

			publicKeyExportArmored, err = parsedKey.GetArmoredPublicKey()
			if err != nil {
				return "", err
			}

			break
		}
	}

	return publicKeyExportArmored, nil
}

func (c *Home) getPrivateKeyByID(ID string) (string, error) {
	privateKey := GPGKey{}
	privateKeyExportArmored := ""
	for _, candidate := range c.keys {
		if candidate.ID == c.privateKeyID {
			privateKey = candidate

			parsedKey, err := crypto.NewKeyFromArmored(string(privateKey.Content)) // TODO: Use crypt package's implementation to support raw keys
			if err != nil {
				c.panic(err, func() {})

				break
			}

			privateKeyExportArmored, err = parsedKey.Armor()
			if err != nil {
				c.panic(err, func() {})

				break
			}

			break
		}
	}

	return privateKeyExportArmored, nil
}

func (c *Home) getPasswordForKey(ID string, checkPassword func(password string) error) string {
	c.keyPasswordModalOpen = true
	c.keyPasswordModalKeyID = ID

	c.Update()

	for {
		password := <-c.keyPasswordChan

		// Stop if no password has been entered
		if password == "" {
			c.keyPasswordModalOpen = false

			return ""
		}

		if err := checkPassword(password); err != nil {
			c.handleWrongPassword(err)

			continue
		}

		c.keyPasswordModalOpen = false
		c.clearWrongPassword()

		c.Update()

		return password
	}
}
