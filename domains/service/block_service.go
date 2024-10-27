package service

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/iagoholekdev/blockchain/domains/models"
)

func CalculateHash(b models.Block) string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(oldBlock models.Block, data string) (models.Block, error) {
	var newBlock models.Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = data
	newBlock.PreviousHash = oldBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock, nil
}

func IsBlockValid(newBlock, oldBlock models.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PreviousHash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
