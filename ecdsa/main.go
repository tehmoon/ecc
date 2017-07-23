package ecdsa

import (
  "crypto/ecdsa"
  "crypto/elliptic"
  "io"
  "math/big"
)

// PublicKey represents an ECDSA public key.
type PublicKey struct {
  elliptic.Curve
  X, Y *big.Int
}

// PrivateKey represents a ECDSA private key.
type PrivateKey struct {
  PublicKey
  D *big.Int
}

type ecdsaSignature struct {
  R, S *big.Int
}

func Sign(rand io.Reader, priv *ecdsa.PrivateKey, hash []byte) (*big.Int, *big.Int, error) {
  r, s, err := ecdsa.Sign(rand, priv, hash)

  return r, s, err
}

func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool {
  params := pub.Curve.Params()

  if (len(hash) * 8) != params.BitSize {
    return false
  }

  if r.Sign() <= 0 || s.Sign() <= 0 {
    return false
  }

  if r.Cmp(params.N) >= 0 || s.Cmp(params.N) >= 0 {
    return false
  }

  e := new(big.Int).SetBytes(hash)

  w := new(big.Int).ModInverse(s, params.N)


  u1 := new(big.Int).Mul(e, w)
  u1 = new(big.Int).Mod(u1, params.N)

  u2 := new(big.Int).Mul(r, w)
  u2 = new(big.Int).Mod(u2, params.N)

  u1x, u1y := pub.Curve.ScalarBaseMult(u1.Bytes())
  u2x, u2y := pub.Curve.ScalarMult(pub.X, pub.Y, u2.Bytes())

  x, y := pub.Curve.Add(u1x, u1y, u2x, u2y)

  if x.Sign() == 0 && y.Sign() == 0 {
    return false
  }

  v := new(big.Int).Mod(x, params.N)

  ok := v.Cmp(r)

  if ok == 0 {
    return true
  }

  return false
}

func GenerateKey(c elliptic.Curve, rand io.Reader) (*ecdsa.PrivateKey, error) {
  priv, err := ecdsa.GenerateKey(c, rand)

  return priv, err
}
