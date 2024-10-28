package models

import "github.com/iagoholekdev/blockchain-go/domains/types"

type Header struct {
	Version   uint32
	PrevBlock types.Hash
	TimeStamp uint64
	Height    uint32
	Nonce     uint64
}

type Block struct {
	Header
	Transactions []Transaction
}
