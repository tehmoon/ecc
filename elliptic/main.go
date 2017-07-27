package elliptic

import (
  "crypto/elliptic"
  "math/big"
)

type Curve interface {
  // Params returns the parameters for the curve.
  Params() *elliptic.CurveParams
  // IsOnCurve reports whether the given (x,y) lies on the curve.
  IsOnCurve(x, y *big.Int) bool
  // Add returns the sum of (x1,y1) and (x2,y2)
  Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int)
  // Double returns 2*(x,y)
  Double(x1, y1 *big.Int) (x, y *big.Int)
  // ScalarMult returns k*(Bx,By) where k is a number in big-endian form.
  ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int)
  // ScalarBaseMult returns k*G, where G is the base point of the group
  // and k is an integer in big-endian form.
  ScalarBaseMult(k []byte) (x, y *big.Int)
  // Solve the equation with x when y^2
  SolveY(y *big.Int) (y1, y2 *big.Int)
}
