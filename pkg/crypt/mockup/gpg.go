package crypto

import "github.com/ProtonMail/go-crypto/openpgp"

func IsKeyLocked(key []byte) (bool, string, error) // May also be armored

func ReadKey(
	key []byte, // May also be armored
	password string,
) (*openpgp.Identity, error) // Try to unarmor first, if it doesn't work ignore and continue as though it were plain text

type EncryptConfig struct {
	PublicKey       *openpgp.Identity
	ArmorCyphertext bool
}

type SignatureConfig struct {
	PrivateKey      *openpgp.Identity
	ArmorSignature  bool
	DetachSignature bool
}

func EncryptSign(
	encryptConfig *EncryptConfig, // May also be nil
	signatureConfig *SignatureConfig, // May also be nil

	plaintext []byte,
) (cyphertext []byte, signature []byte, err error) // May also be only cyphertext or signature

type DecryptConfig struct {
	PrivateKey *openpgp.Identity
}

type VerifyConfig struct {
	PublicKey         *openpgp.Identity
	DetachedSignature []byte // May also be armored
}

func DecryptVerify(
	decryptConfig *DecryptConfig, // May also be nil
	verfiyConfig *VerifyConfig, // May also be nil

	cyphertext []byte, // May also be armored
) (plaintext []byte, verified bool, err error) // verified is always false if verifyConfig == nil

func ExportKey(
	key []byte,
	onlyPublicKey bool,
	armor bool,
) ([]byte, error) // May also be armored
