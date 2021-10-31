package crypto

// In the future, the following package design may be sensible as a replacement for the `Helper` package:
//
// package crypto
//
// func ReadKey(
// key []byte,
// password string, // May also be armored
// ) (Key, error) // Try to unarmor first, if it doesn't work ignore and continue as though it were plain text
//
// type EncryptConfig struct {
// PublicKey Key
// ArmorCyphertext bool
// }
//
// type SignatureConfig struct {
// PrivateKey Key
// ArmorSignature bool
// DetachSignature bool
// }
//
// func EncryptSign(
// encryptConfig EncryptConfig, // May also be nil
// signatureConfig SignatureConfig, // May also be nil
//
// plaintext []byte,
// ) (cyphertext []byte, signature []byte, err error) // May also be only cyphertext or strings
//
// type DecryptConfig struct {
// PrivateKey Key
// }
//
// type VerifyConfig struct {
// PublicKey Key
// DetachedSignature []byte // May also be armored
// }
//
// func DecryptVerify(
// decryptConfig DecryptConfig,
// verfiyConfig VerifyConfig,
//
// cyphertext []byte, // May also be armored
// ) (plaintext []byte, verified bool, err error) // verified is always false if verifyConfig == nil

import (
	"bytes"
	"errors"
	"io/ioutil"

	"github.com/ProtonMail/go-crypto/openpgp"
)

func parseEntities(entities openpgp.EntityList, password string) (*openpgp.Entity, error) {
	if len(entities) <= 0 {
		return nil, errors.New("no entities found in keyring")
	}

	entity := entities[0]
	if entity.PrivateKey != nil && entity.PrivateKey.Encrypted {
		if err := entity.PrivateKey.Decrypt([]byte(password)); err != nil {
			return nil, err
		}
	}

	return entity, nil
}

func ReadKey(key []byte, password string) (*openpgp.Entity, error) {
	entities, err := openpgp.ReadKeyRing(bytes.NewReader(key))
	if err != nil {
		return nil, err
	}

	return parseEntities(entities, password)
}

func ReadKeyArmored(key []byte, password string) (*openpgp.Entity, error) {
	entities, err := openpgp.ReadArmoredKeyRing(bytes.NewReader(key))
	if err != nil {
		return nil, err
	}

	return parseEntities(entities, password)

}

func EncryptSign(publicKey *openpgp.Entity, privateKey *openpgp.Entity, plaintext []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	cyphertextWriter, err := openpgp.Encrypt(buf, []*openpgp.Entity{publicKey}, privateKey, nil, nil)
	if err != nil {
		return []byte{}, err
	}

	if _, err := cyphertextWriter.Write(plaintext); err != nil {
		return []byte{}, err
	}
	defer cyphertextWriter.Close()

	cypherText, err := ioutil.ReadAll(buf)
	if err != nil {
		return []byte{}, err
	}

	return cypherText, err
}
