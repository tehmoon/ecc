package ecdsa

import (
  "crypto/ecdsa"
  "crypto/elliptic"
  "io"
  "math/big"
)

func Sign(rand io.Reader, priv *ecdsa.PrivateKey, hash []byte) (*big.Int, *big.Int, error) {
  r, s, err := ecdsa.Sign(rand, priv, hash)

  return r, s, err
}

func Verify(pub *ecdsa.PublicKey, hash []byte, r, s *big.Int) bool {
  ok := ecdsa.Verify(pub, hash, r, s)

  return ok
}

func GenerateKey(c elliptic.Curve, rand io.Reader) (*ecdsa.PrivateKey, error) {
  priv, err := ecdsa.GenerateKey(c, rand)

  return priv, err
}
