package jwt

import (
	"crypto/ed25519"
)

// NewAlgorithmEdDSA returns a new ed25519-based algorithm.
func NewAlgorithmEdDSA(private ed25519.PrivateKey, public ed25519.PublicKey) (Algorithm, error) {
	if private == nil && public == nil {
		return nil, ErrBothKeysAreNil
	}

	a := &edDSAAlg{
		privateKey: private,
		publicKey:  public,
	}
	return a, nil
}

type edDSAAlg struct {
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
}

func (h edDSAAlg) AlgorithmName() AlgorithmName {
	return EdDSA
}

func (h edDSAAlg) SignSize() int {
	return ed25519.SignatureSize
}

func (h edDSAAlg) Sign(payload []byte) ([]byte, error) {
	return ed25519.Sign(h.privateKey, payload), nil
}

func (h edDSAAlg) Verify(payload, signature []byte) error {
	if len(signature) != ed25519.SignatureSize {
		return ErrInvalidSignature
	}
	if !ed25519.Verify(h.publicKey, payload, signature) {
		return ErrInvalidSignature
	}
	return nil
}
