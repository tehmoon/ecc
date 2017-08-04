package main

import (
  "fmt"
  "../ecdsa"
  "../secp256k1"
  "crypto/rand"
)

func main() {
  priv1, _ := ecdsa.GenerateKey(secp256k1.Curve, rand.Reader)
  priv2, _ := ecdsa.GenerateKey(secp256k1.Curve, rand.Reader)

  j := priv1.D.Bytes()
  k := priv2.D.Bytes()

  Ax, Ay := secp256k1.Curve.ScalarBaseMult(j)
  Bx, By := secp256k1.Curve.ScalarBaseMult(k)

  key1, _ := secp256k1.Curve.ScalarMult(Bx, By, j)
  key2, _ := secp256k1.Curve.ScalarMult(Ax, Ay, k)


  fmt.Printf("0x%x\n", key1.Bytes())
  fmt.Printf("0x%x\n", key2.Bytes())
}
