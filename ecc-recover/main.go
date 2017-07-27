package main

// TODO:
//   - finish flag parser
//   - hash message
//   - output possible keys

import (
//  "time"
  "math/big"
//  "crypto/rand"
  "../secp256k1"
  "../ecdsa"
  "errors"
  "fmt"
  "../elliptic"
  "strings"
  "hash"
//  "encoding/hex"
//  "crypto/ecdsa"
//  "crypto/elliptic"
  "flag"
  "os"
  "crypto/sha256"
)

func init() {
  SupportedCurves = make(Curves, 0)
  SupportedCurves.Add(secp256k1.Curve)

  SupportedHashFuncs = make(HashFuncs, 0)
  SupportedHashFuncs.Add("sha256", sha256.New())
}

type Config struct {
  Curve elliptic.Curve
  Message string
  HashFunc *HashFunc
  R *big.Int
  S *big.Int
}

type Curves []elliptic.Curve
type HashFuncs []*HashFunc
type HashFunc struct {
  Name string
  Func hash.Hash
}

var SupportedCurves Curves
var SupportedHashFuncs HashFuncs

func (curves *Curves) Add(curve elliptic.Curve) {
  *curves = append(*curves, curve)
}

func (curves Curves) Search(name string) (elliptic.Curve, bool) {
  name = strings.ToLower(name)

  for _, curve := range curves {
    if strings.ToLower(curve.Params().Name) == name {
      return curve, true
    }
  }

  return nil, false
}

func (hashFuncs *HashFuncs) Add(name string, f hash.Hash) {
  hashFunc := &HashFunc{
    Name: strings.ToLower(name),
    Func: f,
  }

  *hashFuncs = append(*hashFuncs, hashFunc)
}

func (hashFuncs HashFuncs) Search(name string) (*HashFunc, bool) {
  name = strings.ToLower(name)

  for _, hashFunc := range hashFuncs {
    if hashFunc.Name == name {
      return hashFunc, true
    }
  }

  return nil, false
}

func parseArgs() (*Config, error) {
  var (
    curveName, message string
    hashFuncName, r, s string
    listCurves, listHashFuncs bool
  )

  flag.StringVar(&curveName, "name", "", "Name of the curve.")
  flag.BoolVar(&listCurves, "list_curves", false, "List all supported curves.")
  flag.BoolVar(&listHashFuncs, "list_hash_functions", false, "List all supported hash functions.")
  flag.StringVar(&hashFuncName,"hash", "sha256", "Select hash function.")
  flag.StringVar(&message, "message", "", "Message that goes with signature r,s.")
  flag.StringVar(&r, "r", "", "R part of the signature in hex/bin/dec.")
  flag.StringVar(&s, "s", "", "S part of the signature in hex/bin/dec.")

  flag.Parse()

  if listCurves {
    for _, curve := range SupportedCurves {
      fmt.Printf("- %s\n", curve.Params().Name)
    }

    return nil, NewFlagError("")
  }

  if listHashFuncs {
    for _, hashFunc := range SupportedHashFuncs {
      fmt.Printf("- %s\n", hashFunc.Name)
    }

    return nil, NewFlagError("")
  }

  if curveName == "" {
    return nil, NewFlagError("-name is mandatory and should be one of the supported curve")
  }

  if r == "" {
    return nil, NewFlagError("-r is mandatory")
  }

  if s == "" {
    return nil, NewFlagError("-s is mandatory")
  }

  if message == "" {
    return nil, NewFlagError("-message is mandatory")
  }

  curve, found := SupportedCurves.Search(curveName)
  if ! found {
    return nil, NewFlagError("-name should be one of -list_curves")
  }

  hashFunc, found := SupportedHashFuncs.Search(hashFuncName)
  if ! found {
    return nil, NewFlagError("-hash should be one of -list_hash_functions")
  }

  R, ok := new(big.Int).SetString(r, 0)
  if ! ok {
    return nil, NewFlagError("-r is neither in hex, dec or bin")
  }

  S, ok := new(big.Int).SetString(s, 0)
  if ! ok {
    return nil, NewFlagError("-s is neither in hex, doc or bin")
  }

  config := &Config{
    R: R,
    S: S,
    Message: message,
    HashFunc: hashFunc,
    Curve: curve,
  }

  return config, nil
}

type FlagError struct {
  Code int
  Message string
}

func NewFlagError(message string) error {
  return &FlagError{
    Message: message,
  }
}

func (err FlagError) Exit() {
  if err.Message != "" {
    fmt.Fprintf(os.Stderr, err.Error())
    flag.Usage()
  }

  os.Exit(0)
}

func (err FlagError) Error() string {
  return fmt.Sprintf("%s.\n", err.Message)
}

func main() {
  config, err := parseArgs()
  if err != nil {
    if err, ok := err.(*FlagError); ok {
      err.Exit()
    }

    message := fmt.Sprintf("Error not recognized.\nErr:%v", err)
    panic(errors.New(message))
  }

  config.HashFunc.Func.Write([]byte(config.Message))
  hash := config.HashFunc.Func.Sum(nil)

  pub1, pub2 := ecdsa.Recover(config.Curve, hash, config.R, config.S)

  fmt.Printf("x1: %x\ty1: %x\n", pub1.X, pub1.Y)
  fmt.Printf("x2: %x\ty2: %x\n", pub2.X, pub2.Y)

//  fmt.Println(ecdsa.Recover(secp256k1.Curve, hash, r, s))
}
