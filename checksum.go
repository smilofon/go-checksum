package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

func hashFile(fpath string, cb func() hash.Hash) ([]byte, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return []byte{}, err
	}
	defer f.Close()

	h := cb()
	if _, err := io.Copy(h, f); err != nil {
		return []byte{}, err
	}

	return h.Sum(nil), nil
}

func hashFileToString(fpath string, cb func(string) ([]byte, error)) (string, error) {
	h, err := cb(fpath)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h), nil
}

// Sha1sum calculates SHA1 hash of file
func Sha1sum(fpath string) ([]byte, error) {
	return hashFile(fpath, sha1.New)
}

// Sha1sumAsString returns SHA1 hash of file as hexadecimal string
func Sha1sumAsString(fpath string) (string, error) {
	return hashFileToString(fpath, Sha1sum)
}

// Sha256sum calculates SHA256 hash of file
func Sha256sum(fpath string) ([]byte, error) {
	return hashFile(fpath, sha256.New)
}

// Sha256sumAsString returns SHA256 hash of file as hexadecimal string
func Sha256sumAsString(fpath string) (string, error) {
	return hashFileToString(fpath, Sha256sum)
}

// Sha512sum calculates SHA512 hash of file
func Sha512sum(fpath string) ([]byte, error) {
	return hashFile(fpath, sha512.New)
}

// Sha512sumAsString returns SHA512 hash of file as hexadecimal string
func Sha512sumAsString(fpath string) (string, error) {
	return hashFileToString(fpath, Sha512sum)
}

// Md5sum calculates MD5 hash of file
func Md5sum(fpath string) ([]byte, error) {
	return hashFile(fpath, md5.New)
}

// Md5sumAsString returns MD5 hash of file as hexadecimal string
func Md5sumAsString(fpath string) (string, error) {
	return hashFileToString(fpath, Md5sum)
}
