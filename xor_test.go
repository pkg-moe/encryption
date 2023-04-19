package encryption

import (
	"encoding/base64"
	"testing"
)

func TestXor(t *testing.T) {
	b, _ := base64.StdEncoding.DecodeString("U1lKCAgI")
	key := string(b)

	t.Run("encrypt", func(t *testing.T) {
		data := "testdemo.data.com/http://s"
		out := Xor([]byte(data), []byte(key))
		want := "'<9|lm>6dli|2w)ge';->x2'|*"
		if string(out) != want {
			t.Errorf("want: %s, got: %s", want, out)
		}
	})

	t.Run("decrypt", func(t *testing.T) {
		data := "'<9|lm>6dli|2w)ge';->x2'|*"
		out := Xor([]byte(data), []byte(key))
		want := "testdemo.data.com/http://s"
		if string(out) != want {
			t.Errorf("want: %s, got: %s", want, out)
		}
	})
}
