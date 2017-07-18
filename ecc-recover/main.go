package main

import (
//  "time"
//  "math/big"
  "crypto/rand"
  "../secp256k1"
//  "../ecdsa"
  "fmt"
  "encoding/hex"
  "crypto/ecdsa"
)

func main() {
//  x1, _ := new(big.Int).SetString("34F9460F0E4F08393D192B3C5133A6BA099AA0AD9FD54EBCCFACDFA239FF49C6", 16)
//  y1, _ := new(big.Int).SetString("0B71EA9BD730FD8923F6D25A7A91E7DD7728A960686CB5A901BB419E0F2CA232", 16)
//
//  x2, _ := new(big.Int).SetString("c6ac6a1a06797c871bce3ddcecea64526d5abebb96cf8212773e2473049dd6f8", 16)
//  y2, _ := new(big.Int).SetString("06648a19bef040a78d38fd0e4f9f21a1de99a241a6a67bbf245e7b60f154841d", 16)
//
//  now := time.Now()
//  fmt.Println(secp256k1.Curve.IsOnCurve(x1, y1))
//  fmt.Println(secp256k1.Curve.IsOnCurve(x2, y2))
//
//  rx, ry := secp256k1.Curve.Add(x1, y1, x2, y2)
//  fmt.Println(secp256k1.Curve.IsOnCurve(rx, ry))
//
//  fmt.Println(x1, y1)
//  doubleX, doubleY := secp256k1.Curve.Double(x1, y1)
//  fmt.Println(doubleX, doubleY)
//  fmt.Println(secp256k1.Curve.Double(doubleX, doubleY))
//  fmt.Println(secp256k1.Curve.IsOnCurve(secp256k1.Curve.ScalarBaseMult([]byte{100})))
//  fmt.Println(time.Since(now))

//  r, _ := new(big.Int).SetString("9D7799380AC4FD4821FCE0CEA4A14C9E9011669D7BD1C4CF26C87288020716F0", 16)
//  s, _ := new(big.Int).SetString("F21A9EBF65CFB11B05F04963C9553C2968E350E245ED480A806D38D2B30939A3", 16)

//  X, _ := new(big.Int).SetString("ec796f7f4b04ef1a58955240c65fb4808687e412b4ba31fc0c03dc104f9782a7", 16)
//  Y, _ := new(big.Int).SetString("e06bb80f67e0e2544d5c58b167810afb39e823185a94ea45e5b28e376befc596", 16)
//  D, _ := new(big.Int).SetString("C3D28139DBFF0241490481617BB3477829FFC97F732D4200B0ADE7F1A760D3CE", 16)
//
//  pub := &ecdsa.PublicKey{
//    secp256k1.Curve,
//    X,
//    Y,
//  }
//
//  priv := &ecdsa.PrivateKey{
//    *pub,
//    D,
//  }

  hash, _ := hex.DecodeString("31f7a65e315586ac198bd798b6629ce4903d0899476d5741a9f32e2e521b6a66")

//  r, s, _ = ecdsa.Sign(rand.Reader, priv, hash)
  priv, _ := ecdsa.GenerateKey(secp256k1.Curve, rand.Reader)
  fmt.Println(priv.X, priv.Y)
  fmt.Println(secp256k1.Curve.IsOnCurve(priv.X, priv.Y))
  r, s, _ := ecdsa.Sign(rand.Reader, priv, hash)

  fmt.Println(ecdsa.Verify(&priv.PublicKey, hash, r, s))

//  fmt.Println(r, s)
//
//  fmt.Println(ecdsa.Verify(pub, hash, r, s))

//  priv, err := ecdsa.GenerateKey(secp256k1.Curve, rand.Reader)
//  if err != nil {
//    panic(err)
//  }
//
//  fmt.Println(priv.D, len(priv.D.Bytes()))
//  fmt.Println(priv.X, len(priv.X.Bytes()), priv.Y, len(priv.Y.Bytes()))
//  fmt.Println(secp256k1.Curve.IsOnCurve(priv.X, priv.Y))
}
