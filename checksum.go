package checksum

import (
	"crypto/sha1"
	"crypto/sha256"
	"io"
	"os"
)

// Sha256sum calculates SHA256 hash of file
func Sha256sum(fpath string) ([]byte, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return []byte{}, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return []byte{}, err
	}

	return h.Sum(nil), nil
}

// Sha1sum calculates SHA1 hash of file
func Sha1sum(fpath string) ([]byte, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return []byte{}, err
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return []byte{}, err
	}

	return h.Sum(nil), nil
}
