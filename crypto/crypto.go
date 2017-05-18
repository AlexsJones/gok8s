package crypto

import (
	"crypto/rand"
	"io"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/nacl/secretbox"
)

//HashPassword ...
func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

//CheckPasswordHash ...
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//EncryptText ....
func EncryptText(text []byte, secretKey [32]byte) []byte {
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}
	return secretbox.Seal(nonce[:], text, &nonce, &secretKey)
}

//DecryptText ...
func DecryptText(encrypted []byte, key [32]byte) ([]byte, bool) {
	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])
	return secretbox.Open([]byte{}, encrypted[24:], &decryptNonce, &key)
}

//ByteSplit ...
func ByteSplit(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}
