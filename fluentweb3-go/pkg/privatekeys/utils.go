package privatekeys

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	defaults "github.com/mcuadros/go-defaults"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"
	"github.com/tyler-smith/go-bip39"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"time"
)

var (
	ErrInvalidMnemonic error = errors.New("invalid mnemonic")
)

func NewFromKeyString(keyString string) (*ecdsa.PrivateKey, error) {
	if len(keyString) >= 2 && keyString[0:2] == "0x" {
		keyString = keyString[2:]
	}

	privateKey, err := crypto.HexToECDSA(keyString)

	if err != nil {
		return nil, errors.Wrap(err, "invalid HEX format of private key")
	}

	return privateKey, nil
}

func NewRandom() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(crypto.S256(), rand.Reader)
}

func NewFromMnemonic(mnemonic string, addressIndex int, option *MnemonicOption) (*ecdsa.PrivateKey, error) {

	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, ErrInvalidMnemonic
	}

	if option == nil {
		option = &MnemonicOption{}
	}

	defaults.SetDefaults(option)

	seed := bip39.NewSeed(mnemonic, option.Password)

	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		log.Fatal(err)
	}

	_path, err := hdwallet.ParseDerivationPath(fmt.Sprintf("%v/%v", option.BaseDerivePath, addressIndex))
	if err != nil {
		return nil, err
	}

	account, err := wallet.Derive(_path, false)
	if err != nil {
		log.Fatal(err)
	}

	return wallet.PrivateKey(account)
}

func NewFromKeystore(keyjson []byte, auth string) (*ecdsa.PrivateKey, error) {
	key, err := keystore.DecryptKey(keyjson, auth)
	if err != nil {
		return nil, err
	}
	return key.PrivateKey, nil
}

func NewFromKeystoreFile(filePath string, auth string) (*ecdsa.PrivateKey, error) {
	keyjson, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return NewFromKeystore(keyjson, auth)
}

func SaveAsKeystore(key *ecdsa.PrivateKey, dirPath string, auth string) error {
	addr := crypto.PubkeyToAddress(key.PublicKey)
	// load accounts from dir, return error if exists
	ks := keystore.NewKeyStore(dirPath, keystore.StandardScryptN, keystore.StandardScryptP)
	if ks.HasAddress(addr) {
		return keystore.ErrAccountAlreadyExists
	}

	keyjson, err := ToKeystore(key, auth)
	if err != nil {
		return err
	}

	a := accounts.Account{
		Address: addr,
		URL:     accounts.URL{Scheme: keystore.KeyStoreScheme, Path: filepath.Join(dirPath, keyFileName(addr))},
	}

	tmpName, err := writeTemporaryKeyFile(a.URL.Path, keyjson)
	if err != nil {
		return err
	}

	return os.Rename(tmpName, a.URL.Path)
}

func ToKeystore(key *ecdsa.PrivateKey, auth string) ([]byte, error) {
	id, err := uuid.NewRandom()
	for err != nil {
		id, err = uuid.NewRandom()
	}

	ks := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(key.PublicKey),
		PrivateKey: key,
	}

	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP

	return keystore.EncryptKey(ks, auth, scryptN, scryptP)
}

func writeTemporaryKeyFile(file string, content []byte) (string, error) {
	// Create the keystore directory with appropriate permissions
	// in case it is not present yet.
	const dirPerm = 0700
	if err := os.MkdirAll(filepath.Dir(file), dirPerm); err != nil {
		return "", err
	}
	// Atomic write: create a temporary hidden file first
	// then move it into place. TempFile assigns mode 0600.
	f, err := ioutil.TempFile(filepath.Dir(file), "."+filepath.Base(file)+".tmp")
	if err != nil {
		return "", err
	}
	if _, err := f.Write(content); err != nil {
		f.Close()
		os.Remove(f.Name())
		return "", err
	}
	f.Close()
	return f.Name(), nil
}

// keyFileName implements the naming convention for keyfiles:
// UTC--<created_at UTC ISO8601>-<address hex>
func keyFileName(keyAddr common.Address) string {
	ts := time.Now().UTC()
	return fmt.Sprintf("UTC--%s--%s", toISO8601(ts), hex.EncodeToString(keyAddr[:]))
}

func toISO8601(t time.Time) string {
	var tz string
	name, offset := t.Zone()
	if name == "UTC" {
		tz = "Z"
	} else {
		tz = fmt.Sprintf("%03d00", offset/3600)
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d-%02d-%02d.%09d%s",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), tz)
}
