package blockchain

import (
	"fmt"
	"time"

	"github.com/blockhain/domains/models"
	service "github.com/blockhain/domains/service"
)

func main() {
	genesisBlock := models.Block{}
	genesisBlock = models.Block{0, time.Now().String(), "Genesis Block", "", ""}
	genesisBlock.Hash = service.CalculateHash(genesisBlock)
	fmt.Println(genesisBlock)

	secondBlock, _ := service.GenerateBlock(genesisBlock, "Second Block Data")
	fmt.Println(secondBlock)

	fmt.Println(service.isBlockValid(secondBlock, genesisBlock))
}
