package models

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/iagoholekdev/blockchain-go/domains/types"
	"github.com/stretchr/testify/assert"
)

func TestHeaderEncodeDecode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		TimeStamp: uint64(time.Now().UnixNano()),
		Height:    10,
		Nonce:     989897,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))
	fmt.Println(h)

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)
}

func TestBlockEncodeDecode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			TimeStamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     989897,
		},
		Transactions: nil,
	}
	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))
	assert.Equal(t, b, bDecode)
}

func TestBlockHaash(t *testing.T) {
	slice := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			TimeStamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     989897,
		},
		Transactions: []Transaction{},
	}

	hash := slice.Hash()
	assert.False(t, hash.IsZero())

}
