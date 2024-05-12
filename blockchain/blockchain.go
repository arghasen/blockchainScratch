package blockchain

type BlockChain struct {
	Blocks []*Block
}
type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
	Coinbase bool
}

func InitBlockChain() *BlockChain {
	return &BlockChain{Blocks: []*Block{CreateGenesisBlock()}}
}

func (bc *BlockChain) AddBlock(data string, coinbaseRcpt string, transactions []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	coinbaseTransaction := &Transaction{
		Sender:   "Coinbase",
		Receiver: coinbaseRcpt,
		Amount:   10.0,
		Coinbase: true,
	}
	transactions = append([]*Transaction{coinbaseTransaction}, transactions...)
	newBlock := CreateBlock(data, prevBlock.Hash, transactions)
	bc.Blocks = append(bc.Blocks, newBlock)
}
