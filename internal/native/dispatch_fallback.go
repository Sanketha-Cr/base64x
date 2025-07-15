//go:build !amd64 || noasm || appengine
// +build !amd64 noasm appengine

package native

import (
	"encoding/base64"
	"unsafe"
)

//go:nosplit
func B64Decode(out *[]byte, src unsafe.Pointer, length int, mod int) (ret int) {
	// Fallback to standard library for non-amd64 architectures
	srcBytes := (*[1 << 30]byte)(src)[:length:length]
	decoded, err := base64.StdEncoding.DecodeString(string(srcBytes))
	if err != nil {
		return -1
	}
	*out = decoded
	return len(decoded)
}

//go:nosplit
func B64Encode(out *[]byte, src *[]byte, mod int) {
	// Fallback to standard library for non-amd64 architectures
	encoded := base64.StdEncoding.EncodeToString(*src)
	*out = []byte(encoded)
}