package main

import (
//  "time"
  "math/big"
//  "crypto/rand"
  "../secp256k1"
//  "../ecdsa"
  "fmt"
  "encoding/hex"
  "crypto/ecdsa"
//  "crypto/elliptic"
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

  r, _ := new(big.Int).SetString("456C52550B86BDE8243D34F75B7ECFCD4529836BF14E0215495FB26BD875C84E", 16)
  s, _ := new(big.Int).SetString("65579FB5B5D0BEF50C70D1C7E59870A4272D875FFE851870939DC80A9DAD9CF9", 16)
//
  X, _ := new(big.Int).SetString("96860f36f24bd621ad15cd9462e1c06c42038406a355eda8d64b27c7f1b1ea22", 16)
  Y, _ := new(big.Int).SetString("9f1f54dd01f70e6185bef7c8cf250b8527593ba13ecc73ec9ae2b4c7a9f370d7", 16)
//  D, _ := new(big.Int).SetString("17536788293279665543056757257736289267388803240572974354862148677541593326046", 10)

  pub := &ecdsa.PublicKey{
    secp256k1.Curve,
    X,
    Y,
  }

//  priv := &ecdsa.PrivateKey{
//    *pub,
//    D,
//  }

  hash, _ := hex.DecodeString("31f7a65e315586ac198bd798b6629ce4903d0899476d5741a9f32e2e521b6a66")
//  hash, _ := new(big.Int).SetString("61068175463767656649243450861611941262877802461254223945978289328437287147697", 10)
//  fmt.Println(len(hash.Bytes()))

//  r, s, _ = ecdsa.Sign(rand.Reader, priv, hash)
//  priv, _ := ecdsa.GenerateKey(secp256k1.Curve, rand.Reader)
//  fmt.Println(priv.X, priv.Y)

// RECOVER
////  params := secp256k1.Curve
////
////  x := new(big.Int).Exp(priv.X, new(big.Int).SetInt64(3), params.P)
////  ysquare := new(big.Int).Add(x, params.B)
////
////  y := new(big.Int).ModSqrt(ysquare, params.P)
////  negY := new(big.Int).Sub(params.P, y)
////
////  fmt.Println("y:", y)
////  fmt.Println("-y", negY)

  fmt.Println(X, Y)
  params := secp256k1.Curve

  x := new(big.Int).Exp(r, new(big.Int).SetInt64(3), params.P)
  ysquare := new(big.Int).Add(x, params.B)

  Ry1 := new(big.Int).ModSqrt(ysquare, params.P)
  Ry2 := new(big.Int).Sub(params.P, Ry1)

  R := new(big.Int).ModInverse(r, params.N)

  S := s.Bytes()

  sRx1, sRy1 := secp256k1.Curve.ScalarMult(r, Ry1, S)
  sRx2, sRy2 := secp256k1.Curve.ScalarMult(r, Ry2, S)

  zGx, zGy := secp256k1.Curve.ScalarMult(params.Gx, params.Gy, hash)

  zGy.Sub(params.P, zGy)

  add1x, add1y := secp256k1.Curve.Add(sRx1, sRy1, zGx, zGy)
  add2x, add2y := secp256k1.Curve.Add(sRx2, sRy2, zGx, zGy)

  x1, y1 := secp256k1.Curve.ScalarMult(add1x, add1y, R.Bytes())
  x2, y2 := secp256k1.Curve.ScalarMult(add2x, add2y, R.Bytes())

  fmt.Println(x1, y1)
  fmt.Println(x2, y2)


//  fmt.Println(priv.X, priv.Y)
//  fmt.Println(secp256k1.Curve.IsOnCurve(priv.X, priv.Y))
//  r, s, _ := ecdsa.Sign(rand.Reader, priv, hash.Bytes())
//
//  fmt.Println(r, s)

//  fmt.Println(ecdsa.GenerateKey(secp256k1.Curve, rand.Reader))
  fmt.Println(ecdsa.Verify(pub, hash, r, s))

//  fmt.Println("blih")
//
//  priv, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
//  r, s, _ = ecdsa.Sign(rand.Reader, priv, hash)
//
//  fmt.Println(ecdsa.Verify(&priv.PublicKey, hash,r,s))

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
