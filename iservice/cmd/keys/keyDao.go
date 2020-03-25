package keys

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/irisnet/irishub-sdk-go/types"
)

var keysPath = os.ExpandEnv(filepath.Join("$HOME", ".iservice"))

type KeyDAO struct {
	Path string // the Path to store keys
}

func NewKeyDAO() KeyDAO {
	return KeyDAO{
		Path: keysPath,
	}
}

func (dao KeyDAO) Write(name string, store types.Store) error {
	bz, err := json.Marshal(store)
	if err != nil {
		return err
	}
	filePath := os.ExpandEnv(filepath.Join(dao.Path, name+".json"))

	_, err = os.Stat(dao.Path)
	if err != nil && os.IsNotExist(err) {
		if err := os.Mkdir(dao.Path, os.ModePerm); err != nil {
			return err
		}
	}
	fp, err := os.OpenFile(
		filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644,
	)
	defer fp.Close()

	if err != nil {
		return err
	}

	fmt.Fprintf(fp, "%s\n", string(bz))
	return nil
}

func (dao KeyDAO) Read(name string) (types.Store, error) {
	filePath := dao.Path + string(os.PathSeparator) + name + ".json"
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	contentByte, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var keyInfo types.KeyInfo
	err = json.Unmarshal(contentByte, &keyInfo)
	if err != nil {
		return nil, err
	}
	return keyInfo, nil
}

func (dao KeyDAO) Delete(name string) error {
	filePath := dao.Path + string(os.PathSeparator) + name + ".json"
	err := os.Remove(filePath)
	return err
}

func (dao KeyDAO) Encrypt(data string, password string) (string, error) {
	return data, nil
}

func (dao KeyDAO) Decrypt(data string, password string) (string, error) {
	return data, nil
}
