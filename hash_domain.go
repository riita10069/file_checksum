package main

import (
	"crypto/sha256"
	"encoding/hex"
)

type HashDomain struct {
	message string
}


func NewHashDomain(message string) *HashDomain {
	return &HashDomain{message: message}
}


func (d *HashDomain) getBinaryBySHA256() []byte {
	r := sha256.Sum256([]byte(d.message))
	return r[:]
}

func (d *HashDomain) HexDumpBySHA256() string {
	return hex.EncodeToString(d.getBinaryBySHA256())
}

