package secp256k1

import (
  "math/big"
  "crypto/elliptic"
)

type CurveParams struct {
  *elliptic.CurveParams
  A *big.Int
}

func (curve *CurveParams) SolveY(x *big.Int) (*big.Int, *big.Int) {
  // y^2 = x^3 + 0x + 7
  three := new(big.Int).SetInt64(3)
  xcube := new(big.Int).Exp(x, three, curve.P)

  ysquare := new(big.Int).Add(xcube, curve.B)

  y1 := new(big.Int).ModSqrt(ysquare, curve.P)

  y2 := new(big.Int).Sub(curve.P, y1)

  return y1, y2
}

func (curve *CurveParams) Add(Px, Py, Qx, Qy *big.Int) (*big.Int, *big.Int) {
  // S = ( Py - Qy ) / ( Px - Qx )
  // Rx = S**2 - ( Px + Qx )
  // Ry = S * ( Px - Rx ) - Qy

  subPQy := new(big.Int).Sub(Py, Qy)
  subPQx := new(big.Int).Sub(Px, Qx)

  S := new(big.Int).Mul(subPQy, new(big.Int).ModInverse(subPQx, curve.P))

  S2 := new(big.Int).Exp(S, big.NewInt(2), nil)

  Rx := new(big.Int).Sub(S2, new(big.Int).Add(Px, Qx))

  sPxR := new(big.Int).Mul(S, new(big.Int).Sub(Px, Rx))

  Ry := new(big.Int).Sub(sPxR, Py)

  return new(big.Int).Mod(Rx, curve.P), new(big.Int).Mod(Ry, curve.P)
}

func (curve *CurveParams) Double(Px, Py *big.Int) (*big.Int, *big.Int) {
  // S = ( 3 * Px**2 + a ) / 2 * Py )
  // Rx = S**2 - 2Px
  // Ry = S * ( Px - Px ) - Py

  xx := new(big.Int).Exp(Px, big.NewInt(2), nil)
  threeX := new(big.Int).Mul(big.NewInt(3), xx)
  twoY := new(big.Int).Mul(big.NewInt(2), Py)
  S := new(big.Int).Mul(new(big.Int).Add(threeX, curve.A), new(big.Int).ModInverse(twoY, curve.P))

  Rx := new(big.Int).Sub(new(big.Int).Exp(S, big.NewInt(2), nil), new(big.Int).Mul(big.NewInt(2), Px))

  Ry := new(big.Int).Sub(new(big.Int).Mul(S, new(big.Int).Sub(Px, Rx)), Py)

  return new(big.Int).Mod(Rx, curve.P), new(big.Int).Mod(Ry, curve.P)
}

func (curve *CurveParams) Params() (*elliptic.CurveParams) {
  return curve.CurveParams
}

func (curve *CurveParams) ScalarBaseMult(k []byte) (*big.Int, *big.Int) {
  return curve.ScalarMult(curve.Gx, curve.Gy, k)
}

func (curve *CurveParams) ScalarMult(Bx, By *big.Int, k []byte) (*big.Int, *big.Int) {
  Rx := new(big.Int).Set(Bx)
  Ry := new(big.Int).Set(By)

  var AccX *big.Int
  var AccY *big.Int

  // LSB First
  for i := len(k) - 1; i > -1; i-- {
    byte := k[i]
    for bitnumber := 0; bitnumber < 8; bitnumber++ {
      if byte & 1 == 1 {
        if AccX == nil && AccY == nil {
          AccX = new(big.Int).Set(Rx)
          AccY = new(big.Int).Set(Ry)
        } else {
          AccX, AccY = curve.Add(AccX, AccY, Rx, Ry)
        }
      }

      byte = byte >> 1
      Rx, Ry = curve.Double(Rx, Ry)
    }
  }

  return AccX, AccY
}

func (curve *CurveParams) IsOnCurve(x, y *big.Int) bool {
  // y^2 = x^3+0x+7

  yy := new(big.Int).Exp(y, big.NewInt(2), nil)

  xxx := new(big.Int).Exp(x, big.NewInt(3), nil)

  rightSide := new(big.Int).Add(xxx, big.NewInt(7))

  mody := new(big.Int).Mod(yy, curve.P)
  modx := new(big.Int).Mod(rightSide, curve.P)

  ok := false

  if mody.Cmp(modx) == 0 {
    ok = true
  }

  return ok
}
