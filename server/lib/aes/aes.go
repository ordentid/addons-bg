package customAES

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/speps/go-hashids"
	"golang.org/x/crypto/pbkdf2"
)

type CustomAES struct {
	passphrase string
	hash       *hashids.HashID
}

func NewCustomAES(passphrase string) *CustomAES {
	hd := hashids.NewData()
	hd.Salt = passphrase
	h, _ := hashids.NewWithData(hd)
	aes := &CustomAES{
		passphrase: passphrase,
		hash:       h,
	}

	return aes
}

func deriveKey(passphrase string, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 8)
		// http://www.ietf.org/rfc/rfc2898.txt
		// Salt.
		rand.Read(salt)
	}
	return pbkdf2.Key([]byte(passphrase), salt, 1000000000, 64, sha256.New), salt
}

func (t *CustomAES) HashIDEncode(text string) (string, error) {
	id, err := strconv.Atoi(text)
	if err != nil {
		// handle error
		logrus.Errorln(err)
		return "", err
	}
	return t.hash.EncodeInt64([]int64{int64(id)})
}

func (t *CustomAES) HashIDDecode(text string) (string, error) {
	data, err := t.hash.DecodeWithError(text)
	if err != nil {
		return "", err
	}
	if len(data) < 1 {
		return "0", nil
	} else {
		return fmt.Sprintf("%d", data[0]), nil
	}
}

func (t *CustomAES) Encrypt(plaintext string) (text string, err error) {
	// key, salt := deriveKey(t.passphrase, nil)
	iv := make([]byte, 12)
	// http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38d.pdf
	// Section 8.2
	rand.Read(iv)
	b, err := aes.NewCipher([]byte(t.passphrase))
	// b, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		return "", err
	}
	data := aesgcm.Seal(nil, iv, []byte(plaintext), nil)
	return hex.EncodeToString(iv) + "-" + hex.EncodeToString(data), nil
	// return hex.EncodeToString(salt) + "-" + hex.EncodeToString(iv) + "-" + hex.EncodeToString(data), nil
}

func (t *CustomAES) Decrypt(ciphertext string) (text string, err error) {
	if ciphertext == "" {
		return ciphertext, nil
	}
	arr := strings.Split(ciphertext, "-")
	// salt, err := hex.DecodeString(arr[0])
	// if err != nil {
	// 	return "", err
	// }
	iv, err := hex.DecodeString(arr[0])
	if err != nil {
		return "", err
	}
	data, err := hex.DecodeString(arr[1])
	if err != nil {
		return "", err
	}
	// key, _ := deriveKey(t.passphrase, salt)

	b, err := aes.NewCipher([]byte(t.passphrase))
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(b)
	if err != nil {
		return "", err
	}
	data, err = aesgcm.Open(nil, iv, data, nil)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
