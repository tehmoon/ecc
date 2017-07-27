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
//  "../ecdsa"
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
  HashFunc hash.Hash
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

func (curves Curves) IsSupported(name string) bool {
  found := false

  name = strings.ToLower(name)

  for _, curve := range curves {
    if strings.ToLower(curve.Params().Name) == name {
      found = true
      break
    }
  }

  return found
}

func (hashFuncs *HashFuncs) Add(name string, f hash.Hash) {
  hashFunc := &HashFunc{
    Name: strings.ToLower(name),
    Func: f,
  }

  *hashFuncs = append(*hashFuncs, hashFunc)
}

func (hashFuncs HashFuncs) IsSupported(name string) bool {
  found := false

  name = strings.ToLower(name)

  for _, hashFunc := range hashFuncs {
    if hashFunc.Name == name {
      found = true
      break
    }
  }

  return found
}

func parseArgs() (*Config, error) {
  var (
    name, message string
    hashFunc, r, s string
    listCurves, listHashFuncs bool
  )

  flag.StringVar(&name, "name", "", "Name of the curve.")
  flag.BoolVar(&listCurves, "list_curves", false, "List all supported curves.")
  flag.BoolVar(&listHashFuncs, "list_hash_functions", false, "List all supported hash functions.")
  flag.StringVar(&hashFunc,"hash", "sha256", "Select hash function.")
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

  if name == "" {
    return nil, NewFlagError("-name is mandatory and should be one of the supported curve")
  }

  if r == "" {
    return nil, NewFlagError("-r is mandatory")
  }

  if r == "" {
    return nil, NewFlagError("-r is mandatory")
  }

  if s == "" {
    return nil, NewFlagError("-s is mandatory")
  }

  if ! SupportedCurves.IsSupported(name) {
    return nil, NewFlagError("-name should be one of -list_curves")
  }

  R, ok := new(big.Int).SetString(r, 0)
  if ! ok {
    return nil, NewFlagError("-r is neither in hex, dec or bin")
  }

  S, ok := new(big.Int).SetString(s, 0)
  if ! ok {
    return nil, NewFlagError("-s is neither in hex, doc or bin")
  }

  fmt.Println(S, R)

  return nil, nil
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

  fmt.Println(config)

//  r, _ := new(big.Int).SetString("456C52550B86BDE8243D34F75B7ECFCD4529836BF14E0215495FB26BD875C84E", 16)
//  s, _ := new(big.Int).SetString("65579FB5B5D0BEF50C70D1C7E59870A4272D875FFE851870939DC80A9DAD9CF9", 16)
//  X, _ := new(big.Int).SetString("96860f36f24bd621ad15cd9462e1c06c42038406a355eda8d64b27c7f1b1ea22", 16)
//  Y, _ := new(big.Int).SetString("9f1f54dd01f70e6185bef7c8cf250b8527593ba13ecc73ec9ae2b4c7a9f370d7", 16)
//  D, _ := new(big.Int).SetString("17536788293279665543056757257736289267388803240572974354862148677541593326046", 10)

//  pub := &ecdsa.PublicKey{
//    secp256k1.Curve,
//    X,
//    Y,
//  }

//  priv := &ecdsa.PrivateKey{
//    *pub,
//    D,
//  }

//  hash, _ := hex.DecodeString("31f7a65e315586ac198bd798b6629ce4903d0899476d5741a9f32e2e521b6a66")
//  hash, _ := new(big.Int).SetString("61068175463767656649243450861611941262877802461254223945978289328437287147697", 10)
//  fmt.Println(len(hash.Bytes()))

//  r, s, _ = ecdsa.Sign(rand.Reader, priv, hash)
//  priv, _ := ecdsa.GenerateKey(secp256k1.Curve, rand.Reader)
//  fmt.Println(priv.X, priv.Y)

//  fmt.Println(pub)
//  fmt.Println(ecdsa.Recover(secp256k1.Curve, hash, r, s))
}
