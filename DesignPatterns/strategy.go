package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type HashAlgorithm interface {
	Hash(p *passwordProtector)
}

type passwordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

func NewPasswordProtector(user string, passName string, algorithm HashAlgorithm) *passwordProtector {
	return &passwordProtector{
		user:          user,
		passwordName:  passName,
		hashAlgorithm: algorithm,
	}
}

func (p *passwordProtector) setHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *passwordProtector) Protect() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct{}
type SHA256 struct{}
type MD5 struct{}

func (SHA) Hash(p *passwordProtector) {
	h := sha1.New()
	h.Write([]byte(p.passwordName))
	sha1Hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("Hashing using SHA: %s\n", sha1Hash)
}

func (SHA256) Hash(p *passwordProtector) {
	h := sha256.New()
	h.Write([]byte(p.passwordName))
	sha256Hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("Hashing using SHA256: %s\n", sha256Hash)
}

func (MD5) Hash(p *passwordProtector) {
	h := md5.New()
	h.Write([]byte(p.passwordName))
	md5Hash := hex.EncodeToString(h.Sum(nil))
	fmt.Printf("Hashing using MD5: %s\n", md5Hash)
}

func main() {
	sha := SHA{}
	passProtector := NewPasswordProtector("usuario", "mipassword", sha)
	passProtector.Protect()

	// again new hash algorithm
	sha256 := SHA256{}
	passProtector.setHashAlgorithm(sha256)
	passProtector.Protect()

	// set new hash algorithm
	md5 := MD5{}
	passProtector.setHashAlgorithm(md5)
	passProtector.Protect()

}
