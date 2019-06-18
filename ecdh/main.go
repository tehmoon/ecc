package main

import (
  "fmt"
//  "../ecdsa"
//  "../secp256k1"
//  "crypto/rand"
  "golang.org/x/crypto/curve25519"
)

func main() {
//  priv1, _ := ecdsa.GenerateKey(secp256k1.Curve, rand.Reader)
//  priv2, _ := ecdsa.GenerateKey(secp256k1.Curve, rand.Reader)
//
//  j := priv1.D.Bytes()
//  k := priv2.D.Bytes()
//
//  Ax, Ay := secp256k1.Curve.ScalarBaseMult(j)
//  Bx, By := secp256k1.Curve.ScalarBaseMult(k)
//
//  key1, _ := secp256k1.Curve.ScalarMult(Bx, By, j)
//  key2, _ := secp256k1.Curve.ScalarMult(Ax, Ay, k)
//
//
//  fmt.Printf("0x%x\n", key1.Bytes())
//  fmt.Printf("0x%x\n", key2.Bytes())

  j := [32]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30, 0x31, 0x32}
  k := [32]byte{0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x60, 0x61, 0x62, 0x63, 0x64}

  j[0] &= 248
  k[0] &= 248

  j[31] &= 127
  k[31] &= 127

  j[31] |= 64
  k[31] |= 64

  A := [32]byte{}
  B := [32]byte{}

  curve25519.ScalarBaseMult(&A, &j)
  curve25519.ScalarBaseMult(&B, &k)

  Ak := [32]byte{}
  Bj := [32]byte{}

  curve25519.ScalarMult(&Bj, &j, &B)
  curve25519.ScalarMult(&Ak, &k, &A)

  fmt.Printf("%x\n", j)
  fmt.Printf("%x\n", k)
  fmt.Printf("%x\n", A)
  fmt.Printf("%x\n", B)
  fmt.Printf("%x\n", Ak)
  fmt.Printf("%x\n", Bj)
}
