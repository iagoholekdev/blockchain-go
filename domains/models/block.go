package models

type Block struct {
	Index        int    `json:"index"`
	Timestamp    string `json:"timeStamp"`
	Data         string `json:"data"`
	PreviousHash string `json:"previousHash"`
	Hash         string `json:"hash"`
}
