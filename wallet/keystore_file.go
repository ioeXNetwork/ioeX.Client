package wallet

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"

	. "github.com/ioeXNetwork/ioeX.Utility/common"
)

const (
	OldWalletFile       = "wallet.dat"
	DefaultKeystoreFile = "keystore.dat"
	DefaultPath         = "keystore/"
)

type KeystoreFile struct {
	sync.Mutex

	fileName string
	Version  string

	IV                  string
	PasswordHash        string
	MasterKeyEncrypted  string
	PrivateKeyEncrypted string
}

func CreateKeystoreFile(name string) (*KeystoreFile, error) {

	filePath := DefaultPath + name

	// open keystore dir
	err := os.MkdirAll(DefaultPath, 0755)
	if err != nil {
		return nil, err
	}

	if FileExisted(filePath) {
		return nil, errors.New("key store file already exist")
	}

	file := &KeystoreFile{
		fileName: name,
		Version:  KeystoreVersion,
	}

	return file, nil
}

func OpenKeystoreFile(name string) (*KeystoreFile, error) {

	file := &KeystoreFile{
		fileName: name,
	}

	err := file.LoadFromFile()
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (store *KeystoreFile) SetIV(iv []byte) {
	store.IV = BytesToHexString(iv)
}

func (store *KeystoreFile) SetPasswordHash(passwordHash []byte) {
	store.PasswordHash = BytesToHexString(passwordHash)
}

func (store *KeystoreFile) SetMasterKeyEncrypted(masterKeyEncrypted []byte) {
	store.MasterKeyEncrypted = BytesToHexString(masterKeyEncrypted)
}

func (store *KeystoreFile) SetPrivateKeyEncrypted(privateKeyEncrypted []byte) {
	store.PrivateKeyEncrypted = BytesToHexString(privateKeyEncrypted)
}

func (store *KeystoreFile) GetIV() ([]byte, error) {

	iv, err := HexStringToBytes(store.IV)
	if err != nil {
		return nil, err
	}

	return iv, nil
}

func (store *KeystoreFile) GetPasswordHash() ([]byte, error) {

	passwordHash, err := HexStringToBytes(store.PasswordHash)
	if err != nil {
		return nil, err
	}

	return passwordHash, nil
}

func (store *KeystoreFile) GetMasterKeyEncrypted() ([]byte, error) {

	masterKeyEncrypted, err := HexStringToBytes(store.MasterKeyEncrypted)
	if err != nil {
		return nil, err
	}

	return masterKeyEncrypted, nil
}

func (store *KeystoreFile) GetPrivetKeyEncrypted() ([]byte, error) {

	privateKeyEncrypted, err := HexStringToBytes(store.PrivateKeyEncrypted)
	if err != nil {
		return nil, err
	}

	return privateKeyEncrypted, nil
}

func (store *KeystoreFile) LoadFromFile() error {
	store.Lock()
	defer store.Unlock()

	filePath := DefaultPath + store.fileName

	if _, err := os.Stat(filePath); err != nil {
		return errors.New("keystore file not exist")
	}

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, store)
	if err != nil {
		return err
	}

	return nil
}

func (store *KeystoreFile) SaveToFile() error {
	store.Lock()
	defer store.Unlock()

	filePath := DefaultPath + store.fileName

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	data, err := json.Marshal(*store)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
