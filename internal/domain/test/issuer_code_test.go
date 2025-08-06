package test

import (
	"testing"
	"user-account/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestIssuerCode_String_Unknown(t *testing.T) {
	// ✅ Safe: create issuer using constructor returns UnknownIssuer for invalid input
	issuer := domain.NewIssuerCode("some-unknown-issuer")
	assert.Equal(t, "Issuer(UNKNOWN)", issuer.String())

	// ⚠️ Unsafe: to set issuer using direct literal assignment
	issuer = "some-unknown-issuer"
	assert.Equal(t, "Issuer(some-unknown-issuer)", issuer.String())

	// ✅ Safe: to validate and to set issuer (value) using SetValidIssuer()
	issuer = domain.SetValidIssuer("some-unknown-issuer")
	assert.Equal(t, "Issuer(UNKNOWN)", issuer.String())
}
