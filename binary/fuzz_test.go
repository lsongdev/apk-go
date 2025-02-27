package binary_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/lsongdev/apk-go/binary"
)

func FuzzNewXMLFile(f *testing.F) {
	data, err := os.ReadFile("testdata/AndroidManifest.xml")
	if err != nil {
		f.Fatal(err)
	}
	f.Add(data)

	f.Fuzz(func(t *testing.T, data []byte) {
		_, err := binary.NewXMLFile(bytes.NewReader(data))
		if err != nil {
			t.Skip(err)
		}
	})
}
