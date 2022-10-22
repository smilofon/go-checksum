package checksum

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"testing"
)

func TestSha256(t *testing.T) {
	expectedResult, err := hex.DecodeString("0ca093111f402faa55be1cd71006270644b58619eb0c2408b97b7d24bb70dd09")
	if err != nil {
		t.Fatal(err)
	}
	result, err := Sha256sum(filepath.Join("tests", "file"))
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Compare(result, expectedResult) != 0 {
		t.Fatal(fmt.Errorf("invalid sha256, expected '%v' but got '%v'", expectedResult, result))
	}
}

func TestSha1(t *testing.T) {
	expectedResult, err := hex.DecodeString("21d87a17fd56e6a604db5c99ede2511fd9688fc4")
	if err != nil {
		t.Fatal(err)
	}
	result, err := Sha1sum(filepath.Join("tests", "file"))
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Compare(result, expectedResult) != 0 {
		t.Fatal(fmt.Errorf("invalid sha256, expected '%v' but got '%v'", expectedResult, result))
	}
}
