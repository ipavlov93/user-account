package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIssuerCode_String_Unknown(t *testing.T) {
	// ✅ Safe: create issuer using constructor returns UnknownIssuer for invalid input
	// returns UnknownIssuer for unknown inputs.
	code := NewIssuerCode("some-unknown-code")
	assert.Equal(t, "Issuer(UNKNOWN)", code.String())

	// ⚠️ Unsafe: create issuer code using literal
	code = issuer("some-unknown-code")
	assert.Equal(t, "Issuer(some-unknown-code)", code.String())
}
