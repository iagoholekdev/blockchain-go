package types

import "fmt"

const ()

type Hash [32]uint8

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		msg := fmt.Sprintf("bytes should be 32 ")
		panic(msg)
	}
	var value [32]uint8
	for i := 0; i < 32; i++ {
		value[i] = b[i]
	}
	return Hash(value)
}
