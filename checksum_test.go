package checksum

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"testing"
)

func TestSha256(t *testing.T) {
	expected := "0ca093111f402faa55be1cd71006270644b58619eb0c2408b97b7d24bb70dd09"
	result, err := Sha256sumAsString(filepath.Join("tests", "file"))
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Fatal(fmt.Errorf("invalid hash, expected '%v' but got '%v'", expected, result))
	}
	expectedBytes, err := hex.DecodeString(expected)
	if err != nil {
		t.Fatal(err)
	}
	resultBytes, err := Sha256sum(filepath.Join("tests", "file"))
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Compare(resultBytes, expectedBytes) != 0 {
		t.Fatal(fmt.Errorf("invalid hash, expected '%v' but got '%v'", expectedBytes, result))
	}
}

func TestSha1(t *testing.T) {
	expected := "21d87a17fd56e6a604db5c99ede2511fd9688fc4"
	result, err := Sha1sumAsString(filepath.Join("tests", "file"))
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Fatal(fmt.Errorf("invalid hash, expected '%v' but got '%v'", expected, result))
	}
	expectedBytes, err := hex.DecodeString(expected)
	if err != nil {
		t.Fatal(err)
	}
	resultBytes, err := Sha1sum(filepath.Join("tests", "file"))
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Compare(resultBytes, expectedBytes) != 0 {
		t.Fatal(fmt.Errorf("invalid hash, expected '%v' but got '%v'", expectedBytes, result))
	}
}

func TestSha512(t *testing.T) {
	expected := "c491760d91f3d6b33963a2499546e15ed102cdf17a457e1141043d893af20666cc0a965cf25a70f78574b304635b0a11c9ea6970cd4d5d6746ee65b560facab7"
	result, err := Sha512sumAsString(filepath.Join("tests", "file"))
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Fatal(fmt.Errorf("invalid hash, expected '%v' but got '%v'", expected, result))
	}
	expectedBytes, err := hex.DecodeString(expected)
	if err != nil {
		t.Fatal(err)
	}
	resultBytes, err := Sha512sum(filepath.Join("tests", "file"))
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Compare(resultBytes, expectedBytes) != 0 {
		t.Fatal(fmt.Errorf("invalid hash, expected '%v' but got '%v'", expectedBytes, result))
	}
}

func TestMd5(t *testing.T) {
	expected := "7e512c62b17f6229abe559625f098a33"
	result, err := Md5sumAsString(filepath.Join("tests", "file"))
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Fatal(fmt.Errorf("invalid hash, expected '%v' but got '%v'", expected, result))
	}
	expectedBytes, err := hex.DecodeString(expected)
	if err != nil {
		t.Fatal(err)
	}
	resultBytes, err := Md5sum(filepath.Join("tests", "file"))
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Compare(resultBytes, expectedBytes) != 0 {
		t.Fatal(fmt.Errorf("invalid hash, expected '%v' but got '%v'", expectedBytes, result))
	}
}
